package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/mail"
	"net/netip"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sawirricardo/sawirstudio.com/internal/db"
	"github.com/sawirricardo/sawirstudio.com/internal/notify"
)

// Inquiries stores contact-form inquiries in Postgres, and emails each new one.
type Inquiries struct {
	pool   *pgxpool.Pool
	q      *db.Queries
	notify Notifier
	wg     sync.WaitGroup
}

// Notifier sends the email about a new inquiry (notify.Mailer).
type Notifier interface {
	Send(ctx context.Context, msg notify.Message) error
}

// NewInquiries stores inquiries in pool; with a Notifier it also emails them.
func NewInquiries(pool *pgxpool.Pool, n Notifier) *Inquiries {
	return &Inquiries{pool: pool, q: db.New(pool), notify: n}
}

// Wait waits for notifications still being sent (at shutdown, and in tests).
func (in *Inquiries) Wait() { in.wg.Wait() }

// sendNotification emails a stored inquiry in the background, so a slow mail API doesn't
// hold up the visitor, and records when it went out.
func (in *Inquiries) sendNotification(log *slog.Logger, id int64, inq inquiryInput, created time.Time) {
	if in.notify == nil {
		return
	}
	in.wg.Add(1)
	go func() {
		defer in.wg.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := in.notify.Send(ctx, inquiryEmail(id, inq, created)); err != nil {
			log.Error("inquiry email", "id", id, "err", err)
			return
		}
		if err := in.q.MarkInquiryNotified(ctx, id); err != nil {
			log.Error("inquiry notified_at", "id", id, "err", err)
		}
	}()
}

// SendPending emails the inquiries of the last week whose email didn't go out (no token
// yet, or Cloudflare failed). The server runs it at start.
func (in *Inquiries) SendPending(ctx context.Context, log *slog.Logger) {
	if in.notify == nil {
		return
	}
	rows, err := in.q.ListUnnotifiedInquiries(ctx)
	if err != nil {
		log.Error("pending inquiry emails", "err", err)
		return
	}
	for _, r := range rows {
		log.Info("sending a pending inquiry email", "id", r.ID)
		in.sendNotification(log, r.ID, inquiryInput{Name: r.Name, Email: r.Email, Company: r.Company, Message: r.Message, Needs: r.Needs}, r.CreatedAt.Time)
		in.wg.Wait() // one at a time
	}
}

func inquiryEmail(id int64, in inquiryInput, created time.Time) notify.Message {
	subject := "New inquiry from " + in.Name
	if in.Company != "" {
		subject += " (" + in.Company + ")"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Name:    %s\nEmail:   %s\n", in.Name, in.Email)
	if in.Company != "" {
		fmt.Fprintf(&b, "Company: %s\n", in.Company)
	}
	if len(in.Needs) > 0 {
		fmt.Fprintf(&b, "Needs:   %s\n", strings.Join(in.Needs, ", "))
	}
	fmt.Fprintf(&b, "\n%s\n\n-- \nInquiry #%d, sent %s from the sawirstudio.com contact form.\nReply to this email to answer %s.\n",
		in.Message, id, created.UTC().Format("2 Jan 2006 15:04 MST"), in.Name)
	return notify.Message{Subject: subject, Text: b.String(), ReplyTo: in.Email, ReplyToName: in.Name}
}

func (in *Inquiries) Ping(ctx context.Context) error { return in.pool.Ping(ctx) }

// At most this many inquiries per IP address in rateWindow minutes.
const (
	rateLimit  = 5
	rateWindow = 10
)

var needs = map[string]bool{"Application": true, "Website": true, "Design": true, "Not sure yet": true}

type inquiryInput struct {
	Name, Email, Company, Message string
	Needs                         []string
}

func (in inquiryInput) validate() error {
	switch {
	case in.Name == "" || in.Email == "" || in.Message == "":
		return errors.New("Please fill in your name, email and message.")
	case utf8.RuneCountInString(in.Name) > 200, utf8.RuneCountInString(in.Company) > 200,
		len(in.Email) > 320, utf8.RuneCountInString(in.Message) > 10000:
		return errors.New("One of the fields is too long.")
	}
	if a, err := mail.ParseAddress(in.Email); err != nil || a.Address != in.Email {
		return errors.New("That email address doesn't look right.")
	}
	for _, n := range in.Needs {
		if !needs[n] {
			return errors.New("Please pick from the listed choices.")
		}
	}
	if in.looksLikeSpam() {
		return errSpam
	}
	return nil
}

// errSpam is the reply to an inquiry that looks like spam, worded so a person can still send it.
var errSpam = errors.New("Please tick what you need help with, so we know this isn't spam.")

// looksLikeSpam is true for a message with a link when no need is ticked: the spam the form
// gets is a pitch with a link and the checkboxes left alone.
func (in inquiryInput) looksLikeSpam() bool {
	if len(in.Needs) > 0 {
		return false
	}
	m := strings.ToLower(in.Message)
	return strings.Contains(m, "http://") || strings.Contains(m, "https://") || strings.Contains(m, "www.")
}

// inquiryHandler is POST /api/inquiries, from the contact form. With JavaScript the form sends
// it with fetch and asks for JSON; without, the browser posts it and gets a thank-you page.
func inquiryHandler(cfg Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wantsJSON := strings.Contains(r.Header.Get("Accept"), "application/json")
		reply := func(status int, msg string) {
			if wantsJSON {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				json.NewEncoder(w).Encode(map[string]any{"ok": status < 300, "message": msg})
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(status)
			resultPage(inquiryResult(status < 300, msg)).Render(r.Context(), w)
		}
		if cfg.Inquiries == nil {
			reply(http.StatusServiceUnavailable, "The contact form is not available right now.")
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 256<<10)
		if err := r.ParseMultipartForm(256 << 10); err != nil && !errors.Is(err, http.ErrNotMultipart) {
			reply(http.StatusBadRequest, "We couldn't read the form.")
			return
		}
		// The hidden "website" field is left empty by people; pretend it worked for bots.
		if r.PostForm.Get("website") != "" {
			reply(http.StatusOK, "Your inquiry has been sent.")
			return
		}
		in := inquiryInput{
			Name:    strings.TrimSpace(r.PostForm.Get("name")),
			Email:   strings.TrimSpace(r.PostForm.Get("email")),
			Company: strings.TrimSpace(r.PostForm.Get("company")),
			Message: strings.TrimSpace(r.PostForm.Get("message")),
			Needs:   r.PostForm["need"],
		}
		if err := in.validate(); err != nil {
			if errors.Is(err, errSpam) {
				cfg.Log.Info("inquiry looks like spam", "email", in.Email)
			}
			reply(http.StatusUnprocessableEntity, err.Error())
			return
		}
		ip, ipErr := netip.ParseAddr(clientIP(cfg, r))
		var ipp *netip.Addr
		if ipErr == nil {
			ip = ip.Unmap()
			ipp = &ip
			n, err := cfg.Inquiries.q.CountRecentInquiriesFromIP(r.Context(), db.CountRecentInquiriesFromIPParams{Ip: ipp, Minutes: rateWindow})
			if err != nil {
				cfg.Log.Error("inquiry rate check", "err", err)
				reply(http.StatusInternalServerError, "Something went wrong. Please try again in a moment.")
				return
			}
			if n >= rateLimit {
				reply(http.StatusTooManyRequests, "That's a lot of inquiries. Please try again in a few minutes.")
				return
			}
		}
		if in.Needs == nil {
			in.Needs = []string{}
		}
		row, err := cfg.Inquiries.q.CreateInquiry(r.Context(), db.CreateInquiryParams{
			Name: in.Name, Email: in.Email, Company: in.Company, Needs: in.Needs, Message: in.Message,
			Ip: ipp, UserAgent: truncate(r.UserAgent(), 500),
		})
		if err != nil {
			cfg.Log.Error("inquiry insert", "err", err)
			reply(http.StatusInternalServerError, "Something went wrong. Please try again in a moment.")
			return
		}
		cfg.Log.Info("inquiry", "id", row.ID, "email", in.Email)
		cfg.Inquiries.sendNotification(cfg.Log, row.ID, in, row.CreatedAt.Time)
		reply(http.StatusCreated, "Your inquiry has been sent.")
	})
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	for n > 0 && !utf8.RuneStart(s[n]) {
		n--
	}
	return s[:n]
}

// clientIP is the request's address; behind Cloudflare, the visitor's (CF-Connecting-IP, or
// the first in X-Forwarded-For).
func clientIP(cfg Config, r *http.Request) string {
	if cfg.TrustProxy {
		if ip := r.Header.Get("CF-Connecting-IP"); ip != "" {
			return strings.TrimSpace(ip)
		}
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			first, _, _ := strings.Cut(xff, ",")
			return strings.TrimSpace(first)
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
