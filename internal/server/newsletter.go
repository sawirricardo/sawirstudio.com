package server

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"html"
	"io"
	"io/fs"
	"log/slog"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/sawirricardo/sawirstudio.com/internal/db"
	"github.com/sawirricardo/sawirstudio.com/internal/notify"
)

// post is what a newsletter email says about a blog post, read from its built page.
type post struct {
	URL, Title, Description string
}

var (
	metaTag  = regexp.MustCompile(`<meta\s[^>]*>`)
	metaAttr = regexp.MustCompile(`([a-z:-]+)="([^"]*)"`)
)

// readPost reads a post's title and description from its built page in site, like
// posts/services-is-next/index.html.
func readPost(site fs.FS, siteURL, urlPath string) (post, error) {
	b, err := fs.ReadFile(site, path.Join(strings.Trim(urlPath, "/"), "index.html"))
	if err != nil {
		return post{}, err
	}
	p := post{URL: siteURL + urlPath}
	for _, tag := range metaTag.FindAllString(string(b), -1) {
		attrs := map[string]string{}
		for _, m := range metaAttr.FindAllStringSubmatch(tag, -1) {
			attrs[m[1]] = html.UnescapeString(m[2])
		}
		switch {
		case attrs["property"] == "og:title":
			p.Title = strings.TrimSuffix(attrs["content"], " | Sawir Studio")
		case attrs["name"] == "description":
			p.Description = attrs["content"]
		}
	}
	if p.Title == "" {
		return post{}, fmt.Errorf("%s has no og:title", urlPath)
	}
	return p, nil
}

// issueEmail is the email about a new post, for one subscriber.
func (s *Subscribers) issueEmail(p post, sub db.Subscriber) notify.Message {
	unsub := s.link("/unsubscribe", sub.Token)
	text := p.Title + "\n\n" + p.Description + "\n\nRead it: " + p.URL + "\n\n" +
		"-- \nYou get this because you subscribed to the Sawir Studio blog.\nUnsubscribe: " + unsub + "\n"
	body := `<h1 style="font-size: 24px; line-height: 1.25;"><a href="` + html.EscapeString(p.URL) + `" style="color: #1a1a1a;">` + html.EscapeString(p.Title) + `</a></h1>` +
		`<p>` + html.EscapeString(p.Description) + `</p>` +
		`<p><a href="` + html.EscapeString(p.URL) + `">Read the post</a></p>` +
		`<p style="margin-top: 32px; font-size: 13px; color: #666;">You get this because you subscribed to the Sawir Studio blog. ` +
		`<a href="` + html.EscapeString(unsub) + `" style="color: #666;">Unsubscribe</a></p>`
	return notify.Message{
		To: sub.Email, From: s.From, FromName: "Sawir Studio",
		Subject: p.Title, Text: text, HTML: emailHTML(body),
		Headers: map[string]string{
			"List-Unsubscribe":      "<" + unsub + ">",
			"List-Unsubscribe-Post": "List-Unsubscribe=One-Click",
		},
	}
}

// SendOptions change what SendIssue does.
type SendOptions struct {
	DryRun bool          // only count who would get it
	Test   string        // send only to this address, and record nothing
	Delay  time.Duration // between emails
	Log    *slog.Logger
}

// SendIssue emails a blog post (its URL path, like /posts/services-is-next/) to every
// confirmed subscriber who hasn't had it, in their language when the post has a
// translation. Each email sent is recorded, so after an error, running it again carries on
// where it stopped. It returns how many it sent (or would send).
func (s *Subscribers) SendIssue(ctx context.Context, site fs.FS, issue string, o SendOptions) (int, error) {
	if !strings.HasPrefix(issue, "/posts/") || !strings.HasSuffix(issue, "/") || issue == "/posts/" {
		return 0, fmt.Errorf("%q is not a post's path, like /posts/services-is-next/", issue)
	}
	if s.mail == nil && !o.DryRun {
		return 0, errors.New("email is not set up (SES_REGION and AWS credentials, NOTIFY_FROM, NOTIFY_TO; see notify.FromEnv)")
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	english, err := readPost(site, s.SiteURL, issue)
	if err != nil {
		return 0, err
	}
	postFor := func(lang string) post {
		if lang == "" || lang == "en" {
			return english
		}
		if p, err := readPost(site, s.SiteURL, "/"+strings.ToLower(lang)+issue); err == nil {
			return p
		}
		return english
	}

	if o.Test != "" {
		if o.DryRun {
			return 1, nil
		}
		msg := s.issueEmail(english, db.Subscriber{Email: o.Test, Token: "test"})
		msg.Subject = "[Test] " + msg.Subject
		return 1, s.mail.Send(ctx, msg)
	}

	subs, err := s.q.ListUnsentSubscribers(ctx, issue)
	if err != nil {
		return 0, err
	}
	if o.DryRun {
		return len(subs), nil
	}
	sent := 0
	for i, sub := range subs {
		if i > 0 && o.Delay > 0 {
			select {
			case <-ctx.Done():
				return sent, ctx.Err()
			case <-time.After(o.Delay):
			}
		}
		sendCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		err := s.mail.Send(sendCtx, s.issueEmail(postFor(sub.Lang), sub))
		cancel()
		if errors.Is(err, notify.ErrNotDelivered) {
			// The address bounced or is suppressed: skip it (recorded, so a rerun doesn't
			// stop on it again) and carry on with the rest.
			o.Log.Warn("newsletter not delivered", "issue", issue, "subscriber", sub.ID, "err", err)
		} else if err != nil {
			return sent, fmt.Errorf("subscriber %d: %w (%d sent; run it again to carry on)", sub.ID, err, sent)
		}
		if err := s.q.RecordNewsletterSend(context.WithoutCancel(ctx), db.RecordNewsletterSendParams{Issue: issue, SubscriberID: sub.ID}); err != nil {
			return sent, fmt.Errorf("recording the email to subscriber %d: %w", sub.ID, err)
		}
		if err == nil {
			sent++
			o.Log.Info("newsletter sent", "issue", issue, "subscriber", sub.ID, "n", sent, "of", len(subs))
		}
	}
	return sent, nil
}

// Stats counts the list.
func (s *Subscribers) Stats(ctx context.Context) (db.CountSubscribersRow, error) {
	return s.q.CountSubscribers(ctx)
}

// ImportResult is what Import did.
type ImportResult struct {
	Added, Existing, Skipped int
}

var importTimeLayouts = []string{time.RFC3339, "2006-01-02 15:04:05 MST", "2006-01-02 15:04:05 -0700", "2006-01-02 15:04:05", "2006-01-02"}

// Import adds confirmed subscribers from a CSV export (ConvertKit's subscriber export): the
// column named like "email" is the address, a "created_at" or "subscribed" column the date
// they signed up. Rows whose "status" column says they left are skipped, and addresses
// already on the list are left as they are.
func (s *Subscribers) Import(ctx context.Context, r io.Reader) (ImportResult, error) {
	var res ImportResult
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	header, err := cr.Read()
	if err != nil {
		return res, fmt.Errorf("reading the CSV header: %w", err)
	}
	emailCol, dateCol, statusCol := -1, -1, -1
	for i, h := range header {
		h = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(h, "\ufeff")))
		switch {
		case emailCol < 0 && strings.Contains(h, "email"):
			emailCol = i
		case dateCol < 0 && (strings.Contains(h, "created") || strings.Contains(h, "subscribed")):
			dateCol = i
		case statusCol < 0 && (h == "status" || h == "state"):
			statusCol = i
		}
	}
	if emailCol < 0 {
		return res, fmt.Errorf("no email column in the CSV header %v", header)
	}
	field := func(row []string, i int) string {
		if i < 0 || i >= len(row) {
			return ""
		}
		return strings.TrimSpace(row[i])
	}
	for {
		row, err := cr.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return res, err
		}
		email := field(row, emailCol)
		switch strings.ToLower(field(row, statusCol)) {
		case "cancelled", "canceled", "unsubscribed", "bounced", "complained", "inactive", "cold":
			res.Skipped++
			continue
		}
		if !validEmail(email) {
			res.Skipped++
			continue
		}
		when := time.Now()
		for _, layout := range importTimeLayouts {
			if t, err := time.Parse(layout, field(row, dateCol)); err == nil {
				when = t
				break
			}
		}
		n, err := s.q.ImportSubscriber(ctx, db.ImportSubscriberParams{
			Email: email, Token: newToken(), ConfirmedAt: pgtype.Timestamptz{Time: when, Valid: true},
		})
		if err != nil {
			return res, fmt.Errorf("%s: %w", email, err)
		}
		if n == 0 {
			res.Existing++
		} else {
			res.Added++
		}
	}
	return res, nil
}
