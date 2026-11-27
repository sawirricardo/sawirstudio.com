package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"net/mail"
	"net/netip"
	"net/url"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sawirricardo/sawirstudio.com/internal/db"
	"github.com/sawirricardo/sawirstudio.com/internal/notify"
)

// Subscribers is the blog's newsletter list, in Postgres. A signup gets an email with a
// link to confirm it (double opt-in); every email we send has a link to unsubscribe.
type Subscribers struct {
	pool *pgxpool.Pool
	q    *db.Queries
	mail Notifier // nil: signups answer 503, since nobody could confirm
	// From is the address the newsletter's emails come from, like newsletter@sawirstudio.com.
	From string
	// SiteURL starts the links in the emails (not the request's Host, which a client sets).
	SiteURL string
}

// NewSubscribers keeps the list in pool and sends its emails with n, from the address from.
func NewSubscribers(pool *pgxpool.Pool, n Notifier, from, siteURL string) *Subscribers {
	return &Subscribers{pool: pool, q: db.New(pool), mail: n, From: from, SiteURL: strings.TrimSuffix(siteURL, "/")}
}

// At most signupLimit signups per IP address in rateWindow minutes, and one confirmation
// email per address per confirmCooldown, so the form can't be used to flood someone's inbox.
const (
	signupLimit     = 5
	confirmCooldown = time.Hour
)

// The translations' path prefixes and their language codes (internal/site/langs.go).
var langPrefixes = map[string]string{
	"zh-hans": "zh-Hans", "zh-hant": "zh-Hant", "ja": "ja", "de": "de", "he": "he", "ru": "ru",
	"fr": "fr", "hi": "hi", "nl": "nl", "it": "it", "ar": "ar",
}

// signupSource is the page the form was sent from (the Referer's path, on our own site),
// and that page's language.
func signupSource(r *http.Request) (path, lang string) {
	lang = "en"
	ref, err := url.Parse(r.Referer())
	if err != nil || ref.Path == "" || !strings.HasPrefix(ref.Path, "/") {
		return "", lang
	}
	path = truncate(ref.Path, 300)
	first, _, _ := strings.Cut(strings.TrimPrefix(path, "/"), "/")
	if code, ok := langPrefixes[first]; ok {
		lang = code
	}
	return path, lang
}

func newToken() string {
	b := make([]byte, 24)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func validEmail(s string) bool {
	if s == "" || len(s) > 320 {
		return false
	}
	a, err := mail.ParseAddress(s)
	return err == nil && a.Address == s && strings.Contains(s[strings.LastIndex(s, "@")+1:], ".")
}

func (s *Subscribers) link(path, token string) string {
	return s.SiteURL + path + "?token=" + url.QueryEscape(token)
}

// confirmationEmail asks the address to confirm its signup.
func (s *Subscribers) confirmationEmail(sub db.Subscriber) notify.Message {
	confirm := s.link("/subscribe/confirm", sub.Token)
	text := "Someone, hopefully you, asked to get new posts from the Sawir Studio blog at this address.\n\n" +
		"Confirm your subscription:\n" + confirm + "\n\n" +
		"If it wasn't you, ignore this email and we won't send you anything.\n\n" +
		"-- \nSawir Studio\n" + s.SiteURL + "/\n"
	page := `<p>Someone, hopefully you, asked to get new posts from the Sawir Studio blog at this address.</p>` +
		`<p><a href="` + html.EscapeString(confirm) + `">Confirm your subscription</a></p>` +
		`<p>If it wasn't you, ignore this email and we won't send you anything.</p>` +
		`<p>-- <br>Sawir Studio<br><a href="` + html.EscapeString(s.SiteURL) + `/">` + html.EscapeString(strings.TrimPrefix(s.SiteURL, "https://")) + `</a></p>`
	return notify.Message{
		To: sub.Email, From: s.From, FromName: "Sawir Studio",
		Subject: "Confirm your subscription to the Sawir Studio blog",
		Text:    text, HTML: emailHTML(page),
	}
}

func emailHTML(body string) string {
	return `<!DOCTYPE html><html><body style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif; font-size: 16px; line-height: 1.5; color: #1a1a1a; max-width: 560px;">` +
		body + `</body></html>`
}

// sendConfirmation emails sub its confirmation link and records when it went out.
func (s *Subscribers) sendConfirmation(ctx context.Context, sub db.Subscriber) error {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	if err := s.mail.Send(ctx, s.confirmationEmail(sub)); err != nil {
		return err
	}
	return s.q.MarkConfirmationSent(context.WithoutCancel(ctx), sub.ID)
}

// signup handles one address from the form. It answers the same way whether the address is
// new, waiting for its confirmation or already subscribed, so the form can't be used to find
// out who is on the list.
func (s *Subscribers) signup(ctx context.Context, email, source, lang string, ip *netip.Addr, ua string) error {
	sub, err := s.q.GetSubscriberByEmail(ctx, email)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		sub, err = s.q.CreateSubscriber(ctx, db.CreateSubscriberParams{
			Email: email, Token: newToken(), Lang: lang, Source: source, Ip: ip, UserAgent: ua,
		})
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil // the same address, signing up twice at once
		}
		if err != nil {
			return err
		}
		return s.sendConfirmation(ctx, sub)
	case err != nil:
		return err
	case sub.UnsubscribedAt.Valid:
		sub, err = s.q.ResubscribeSubscriber(ctx, db.ResubscribeSubscriberParams{
			ID: sub.ID, Token: newToken(), Lang: lang, Source: source, Ip: ip, UserAgent: ua,
		})
		if err != nil {
			return err
		}
		return s.sendConfirmation(ctx, sub)
	}
	if err := s.q.TouchSubscriber(ctx, db.TouchSubscriberParams{ID: sub.ID, Ip: ip, UserAgent: ua}); err != nil {
		return err
	}
	if sub.ConfirmedAt.Valid {
		return nil
	}
	if sub.ConfirmationSentAt.Valid && time.Since(sub.ConfirmationSentAt.Time) < confirmCooldown {
		return nil
	}
	return s.sendConfirmation(ctx, sub)
}

// replier answers a form post with JSON (fetch from main.js) or a page (no JavaScript).
func replier(w http.ResponseWriter, r *http.Request, page func(ok bool, msg string) result) func(int, string) {
	wantsJSON := strings.Contains(r.Header.Get("Accept"), "application/json")
	return func(status int, msg string) {
		if wantsJSON {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(map[string]any{"ok": status < 300, "message": msg})
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(status)
		resultPage(page(status < 300, msg)).Render(r.Context(), w)
	}
}

// subscribeHandler is POST /api/subscribe, from the form at the end of each blog post.
func subscribeHandler(cfg Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		source, lang := signupSource(r)
		back := source
		if back == "" {
			back = "/posts/"
		}
		reply := replier(w, r, func(ok bool, msg string) result {
			if ok {
				return result{Title: "Check your email", Heading: "Check your email", Message: msg, LinkHref: back, LinkText: "Back to the post"}
			}
			return result{Title: "Not subscribed", Heading: "You're not subscribed yet", Message: msg, LinkHref: back, LinkText: "Back to the post"}
		})
		const sent = "We've sent you an email. Click the link in it to confirm your subscription."
		s := cfg.Subscribers
		if s == nil || s.mail == nil {
			reply(http.StatusServiceUnavailable, "The newsletter signup is not available right now.")
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 16<<10)
		if err := r.ParseForm(); err != nil {
			reply(http.StatusBadRequest, "We couldn't read the form.")
			return
		}
		if r.PostForm.Get("website") != "" { // the hidden field only bots fill in
			reply(http.StatusOK, sent)
			return
		}
		email := strings.TrimSpace(r.PostForm.Get("email"))
		if !validEmail(email) {
			reply(http.StatusUnprocessableEntity, "That email address doesn't look right.")
			return
		}
		var ipp *netip.Addr
		if ip, err := netip.ParseAddr(clientIP(cfg, r)); err == nil {
			ip = ip.Unmap()
			ipp = &ip
			n, err := s.q.CountRecentSignupsFromIP(r.Context(), db.CountRecentSignupsFromIPParams{Ip: ipp, Minutes: rateWindow})
			if err != nil {
				cfg.Log.Error("signup rate check", "err", err)
				reply(http.StatusInternalServerError, "Something went wrong. Please try again in a moment.")
				return
			}
			if n >= signupLimit {
				reply(http.StatusTooManyRequests, "That's a lot of signups. Please try again in a few minutes.")
				return
			}
		}
		if err := s.signup(r.Context(), email, source, lang, ipp, truncate(r.UserAgent(), 500)); err != nil {
			cfg.Log.Error("signup", "err", err)
			reply(http.StatusInternalServerError, "We couldn't send the confirmation email. Please try again in a moment.")
			return
		}
		cfg.Log.Info("signup", "source", source)
		reply(http.StatusOK, sent)
	})
}

func page(w http.ResponseWriter, r *http.Request, status int, res result) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	resultPage(res).Render(r.Context(), w)
}

var badLink = result{
	Title: "Link not valid", Heading: "That link didn't work",
	Message:  "It may be from an older email, or it was cut short. Sign up again at the end of any blog post, or write to us through the contact form.",
	LinkHref: "/posts/", LinkText: "Go to the blog",
}

// confirmHandler is GET /subscribe/confirm?token=, the link in the confirmation email.
func confirmHandler(cfg Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.Subscribers == nil {
			http.Error(w, "The newsletter is not available right now.", http.StatusServiceUnavailable)
			return
		}
		token := r.URL.Query().Get("token")
		sub, err := cfg.Subscribers.q.ConfirmSubscriber(r.Context(), token)
		if errors.Is(err, pgx.ErrNoRows) || token == "" {
			page(w, r, http.StatusNotFound, badLink)
			return
		}
		if err != nil {
			cfg.Log.Error("confirm", "err", err)
			http.Error(w, "Something went wrong. Please try again in a moment.", http.StatusInternalServerError)
			return
		}
		cfg.Log.Info("subscriber confirmed", "id", sub.ID)
		page(w, r, http.StatusOK, result{
			Title: "Subscribed", Heading: "You're subscribed 🎉",
			Message:  "We'll email you when a new post goes up. Every email has a link to unsubscribe.",
			LinkHref: "/posts/", LinkText: "Read the blog",
		})
	})
}

// unsubscribeHandler is /unsubscribe?token=, the link in every newsletter email. A GET shows
// a button, since mail scanners open links; the POST (the button, or a mail app's one-click
// unsubscribe, RFC 8058) unsubscribes.
func unsubscribeHandler(cfg Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.Subscribers == nil {
			http.Error(w, "The newsletter is not available right now.", http.StatusServiceUnavailable)
			return
		}
		q := cfg.Subscribers.q
		token := r.URL.Query().Get("token")
		done := result{
			Title: "Unsubscribed", Heading: "You're unsubscribed",
			Message:  "We won't send you any more emails. Changed your mind? Sign up again at the end of any blog post.",
			LinkHref: "/posts/", LinkText: "Go to the blog",
		}
		if r.Method == http.MethodPost {
			sub, err := q.UnsubscribeSubscriber(r.Context(), token)
			if errors.Is(err, pgx.ErrNoRows) || token == "" {
				page(w, r, http.StatusNotFound, badLink)
				return
			}
			if err != nil {
				cfg.Log.Error("unsubscribe", "err", err)
				http.Error(w, "Something went wrong. Please try again in a moment.", http.StatusInternalServerError)
				return
			}
			cfg.Log.Info("unsubscribed", "id", sub.ID)
			page(w, r, http.StatusOK, done)
			return
		}
		sub, err := q.GetSubscriberByToken(r.Context(), token)
		if errors.Is(err, pgx.ErrNoRows) || token == "" {
			page(w, r, http.StatusNotFound, badLink)
			return
		}
		if err != nil {
			cfg.Log.Error("unsubscribe", "err", err)
			http.Error(w, "Something went wrong. Please try again in a moment.", http.StatusInternalServerError)
			return
		}
		if sub.UnsubscribedAt.Valid {
			page(w, r, http.StatusOK, done)
			return
		}
		page(w, r, http.StatusOK, result{
			Title: "Unsubscribe", Heading: "Unsubscribe from the blog?",
			Message: fmt.Sprintf("%s won't get emails about new posts any more.", sub.Email),
			Form:    "/unsubscribe?token=" + url.QueryEscape(token), FormButton: "Unsubscribe",
		})
	})
}
