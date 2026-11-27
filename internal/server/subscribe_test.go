package server

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sawirricardo/sawirstudio.com/internal/notify"
)

func testPool(t *testing.T) *pgxpool.Pool {
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}
	pool, err := pgxpool.New(context.Background(), u)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(), "truncate subscribers cascade"); err != nil {
		t.Fatal(err)
	}
	return pool
}

var linkRE = regexp.MustCompile(`https://sawirstudio\.test(/[a-z/]+)\?token=([A-Za-z0-9_%-]+)`)

// linkIn is the path and query of the first link to path in an email's text.
func linkIn(t *testing.T, m notify.Message, path string) string {
	t.Helper()
	for _, l := range linkRE.FindAllStringSubmatch(m.Text, -1) {
		if l[1] == path {
			return l[1] + "?token=" + l[2]
		}
	}
	t.Fatalf("no %s link in %q", path, m.Text)
	return ""
}

// TestSubscribe needs a migrated database: TEST_DATABASE_URL=postgres://... (see README).
func TestSubscribe(t *testing.T) {
	pool := testPool(t)
	ctx := context.Background()
	mail := &fakeMailer{}
	subs := NewSubscribers(pool, mail, "newsletter@notify.sawirstudio.test", "https://sawirstudio.test/")
	h, err := New(Config{Site: os.DirFS(site(t)), Subscribers: subs, TrustProxy: true})
	if err != nil {
		t.Fatal(err)
	}
	post := func(email, ip, referer, accept string, extra ...string) *httptest.ResponseRecorder {
		form := url.Values{"email": {email}}
		for i := 0; i+1 < len(extra); i += 2 {
			form.Set(extra[i], extra[i+1])
		}
		return do(h, "POST", "/api/subscribe", map[string]string{
			"Content-Type": "application/x-www-form-urlencoded", "Accept": accept,
			"CF-Connecting-IP": ip, "Referer": referer,
		}, form.Encode())
	}
	sent := func() []notify.Message {
		mail.mu.Lock()
		defer mail.mu.Unlock()
		return append([]notify.Message(nil), mail.sent...)
	}

	// A new address gets one confirmation email, from the newsletter's address.
	rec := post("ana@example.com", "192.0.2.1", "https://sawirstudio.com/ja/posts/services-is-next/", "application/json")
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), `"ok":true`) || rec.Header().Get("Cache-Control") != "no-store" {
		t.Fatalf("signup: %d %s", rec.Code, rec.Body)
	}
	if got := sent(); len(got) != 1 || got[0].To != "ana@example.com" || got[0].From != "newsletter@notify.sawirstudio.test" {
		t.Fatalf("confirmation emails: %+v", got)
	}
	var lang, source string
	var confirmSent bool
	pool.QueryRow(ctx, "select lang, source, confirmation_sent_at is not null from subscribers").Scan(&lang, &source, &confirmSent)
	if lang != "ja" || source != "/ja/posts/services-is-next/" || !confirmSent {
		t.Errorf("stored lang %q source %q confirmation sent %v", lang, source, confirmSent)
	}

	// Signing up again soon (any capitalization) answers the same, without another email.
	if rec := post("Ana@Example.com", "192.0.2.2", "", "application/json"); rec.Code != 200 {
		t.Errorf("repeat signup: %d", rec.Code)
	}
	if n := len(sent()); n != 1 {
		t.Errorf("%d emails after a repeat signup, want 1", n)
	}

	// The link in the email confirms; a made-up one doesn't.
	confirm := linkIn(t, sent()[0], "/subscribe/confirm")
	if rec := do(h, "GET", "/subscribe/confirm?token=nope", nil, ""); rec.Code != 404 {
		t.Errorf("bad confirm link: %d", rec.Code)
	}
	if rec := do(h, "GET", confirm, nil, ""); rec.Code != 200 || !strings.Contains(rec.Body.String(), "subscribed") {
		t.Fatalf("confirm: %d %s", rec.Code, rec.Body)
	}
	if rec := do(h, "GET", confirm, nil, ""); rec.Code != 200 {
		t.Errorf("confirming twice: %d", rec.Code)
	}

	// Already subscribed: the same answer, no email (the form doesn't tell who's on the list).
	if rec := post("ana@example.com", "192.0.2.3", "", "application/json"); rec.Code != 200 || !strings.Contains(rec.Body.String(), `"ok":true`) {
		t.Errorf("subscribed address: %d %s", rec.Code, rec.Body)
	}
	if n := len(sent()); n != 1 {
		t.Errorf("%d emails, want 1", n)
	}

	// Without JavaScript the browser gets a page, linking back to the post.
	rec = post("ben@example.com", "192.0.2.4", "https://sawirstudio.com/posts/x/", "text/html")
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "Check your email") || !strings.Contains(rec.Body.String(), `href="/posts/x/"`) {
		t.Errorf("html reply: %d %s", rec.Code, rec.Body)
	}

	// Bad addresses, bots and floods.
	if rec := post("not an email", "192.0.2.5", "", "application/json"); rec.Code != 422 {
		t.Errorf("bad email: %d", rec.Code)
	}
	if rec := post("bot@example.com", "192.0.2.5", "", "application/json", "website", "http://spam"); rec.Code != 200 {
		t.Errorf("honeypot: %d, want a quiet 200", rec.Code)
	}
	var n int
	pool.QueryRow(ctx, "select count(*) from subscribers where email = 'bot@example.com'").Scan(&n)
	if n != 0 {
		t.Error("the bot's signup was stored")
	}
	for i := range signupLimit {
		post("flood"+string(rune('a'+i))+"@example.com", "198.51.100.7", "", "application/json")
	}
	if rec := post("floodz@example.com", "198.51.100.7", "", "application/json"); rec.Code != 429 {
		t.Errorf("over the limit: %d, want 429", rec.Code)
	}

	// A failed confirmation email is an error, not "check your email", and is retried.
	mail.fail = errors.New("cloudflare is down")
	if rec := post("cleo@example.com", "192.0.2.6", "", "application/json"); rec.Code != 500 {
		t.Errorf("failed email: %d, want 500", rec.Code)
	}
	mail.fail = nil
	before := len(sent())
	if rec := post("cleo@example.com", "192.0.2.6", "", "application/json"); rec.Code != 200 || len(sent()) != before+1 {
		t.Errorf("retry after a failed email: %d, %d emails (want %d)", rec.Code, len(sent()), before+1)
	}

	// Sending a post: confirmed subscribers only, in their language, once each.
	root := site(t)
	for p, title := range map[string]string{"posts/hello/index.html": "Hello", "ja/posts/hello/index.html": "こんにちは", "posts/second/index.html": "Second"} {
		os.MkdirAll(filepath.Join(root, filepath.Dir(p)), 0o755)
		os.WriteFile(filepath.Join(root, p), []byte(`<html><head><meta name="description" content="A &amp; B">`+
			`<meta property="og:title" content="`+title+` | Sawir Studio"></head></html>`), 0o644)
	}
	siteFS := os.DirFS(root)
	if n, err := subs.SendIssue(ctx, siteFS, "/posts/hello/", SendOptions{DryRun: true}); err != nil || n != 1 {
		t.Fatalf("dry run: %d %v, want 1 (only ana is confirmed)", n, err)
	}
	before = len(sent())
	if n, err := subs.SendIssue(ctx, siteFS, "/posts/hello/", SendOptions{}); err != nil || n != 1 {
		t.Fatalf("send: %d %v", n, err)
	}
	issue := sent()[before]
	if issue.To != "ana@example.com" || issue.Subject != "こんにちは" || !strings.Contains(issue.Text, "https://sawirstudio.test/ja/posts/hello/") ||
		!strings.Contains(issue.Text, "A & B") || issue.Headers["List-Unsubscribe-Post"] != "List-Unsubscribe=One-Click" ||
		!strings.HasPrefix(issue.Headers["List-Unsubscribe"], "<https://sawirstudio.test/unsubscribe?token=") {
		t.Errorf("issue email: %+v", issue)
	}
	if n, _ := subs.SendIssue(ctx, siteFS, "/posts/hello/", SendOptions{}); n != 0 {
		t.Errorf("sending again sent %d, want 0", n)
	}
	// An address that bounces is skipped, not retried, and doesn't hold up the rest.
	for _, e := range []string{"gil@example.com", "hal@example.com"} {
		if _, err := subs.Import(ctx, strings.NewReader("email\n"+e+"\n")); err != nil {
			t.Fatal(err)
		}
	}
	mail.failTo = map[string]error{"gil@example.com": fmt.Errorf("x: %w", notify.ErrNotDelivered)}
	before = len(sent())
	if n, err := subs.SendIssue(ctx, siteFS, "/posts/hello/", SendOptions{}); err != nil || n != 1 || sent()[before].To != "hal@example.com" {
		t.Errorf("send with a bounce: %d %v, want 1 to hal", n, err)
	}
	if n, err := subs.SendIssue(ctx, siteFS, "/posts/hello/", SendOptions{DryRun: true}); err != nil || n != 0 {
		t.Errorf("after a bounce, a rerun would send %d, want 0", n)
	}
	// Anything else (Cloudflare down, rate limited) stops the send, to carry on later.
	mail.failTo = nil
	mail.fail = errors.New("429")
	if n, err := subs.SendIssue(ctx, siteFS, "/posts/second/", SendOptions{}); err == nil || n != 0 {
		t.Errorf("a failing send: %d %v, want an error", n, err)
	}
	mail.fail = nil

	if _, err := subs.SendIssue(ctx, siteFS, "/about-us/", SendOptions{}); err == nil {
		t.Error("only posts can be sent")
	}

	// Unsubscribing: the link shows a button (mail scanners open links), the POST unsubscribes.
	unsub := linkIn(t, issue, "/unsubscribe")
	if rec := do(h, "GET", unsub, nil, ""); rec.Code != 200 || !strings.Contains(rec.Body.String(), `method="post"`) {
		t.Fatalf("unsubscribe page: %d %s", rec.Code, rec.Body)
	}
	var gone bool
	pool.QueryRow(ctx, "select unsubscribed_at is not null from subscribers where email = 'ana@example.com'").Scan(&gone)
	if gone {
		t.Fatal("opening the unsubscribe link unsubscribed")
	}
	// The mail app's one-click unsubscribe (RFC 8058).
	rec = do(h, "POST", unsub, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, "List-Unsubscribe=One-Click")
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "unsubscribed") {
		t.Fatalf("unsubscribe: %d %s", rec.Code, rec.Body)
	}
	if rec := do(h, "POST", "/unsubscribe?token=nope", nil, ""); rec.Code != 404 {
		t.Errorf("bad unsubscribe link: %d", rec.Code)
	}
	if n, err := subs.SendIssue(ctx, siteFS, "/posts/second/", SendOptions{DryRun: true}); err != nil || n != 2 {
		t.Errorf("after ana unsubscribed, a post would go to %d (%v), want 2 (gil and hal)", n, err)
	}
	// Its old confirmation link doesn't sign it back up.
	if rec := do(h, "GET", confirm, nil, ""); rec.Code != 404 {
		t.Errorf("old confirm link after unsubscribing: %d, want 404", rec.Code)
	}

	// Signing up again after unsubscribing starts over: a new link to confirm.
	before = len(sent())
	if rec := post("ana@example.com", "192.0.2.8", "", "application/json"); rec.Code != 200 || len(sent()) != before+1 {
		t.Fatalf("resubscribe: %d, %d emails", rec.Code, len(sent()))
	}
	if again := linkIn(t, sent()[before], "/subscribe/confirm"); again == confirm {
		t.Error("a new signup should get a new token")
	}

	// ConvertKit's export: confirmed, minus those who left, keeping the ones we have.
	csv := "\ufeffEmail Address,First Name,Created At,Status\n" +
		"dee@example.com,Dee,2023-04-01 10:00:00 UTC,active\n" +
		"ANA@example.com,Ana,2022-01-01 10:00:00 UTC,active\n" +
		"eve@example.com,Eve,2022-01-01 10:00:00 UTC,cancelled\n" +
		"nonsense,,,\n"
	res, err := subs.Import(ctx, strings.NewReader(csv))
	if err != nil || res != (ImportResult{Added: 1, Existing: 1, Skipped: 2}) {
		t.Errorf("import: %+v %v", res, err)
	}
	var year int
	pool.QueryRow(ctx, "select extract(year from confirmed_at)::int from subscribers where email = 'dee@example.com'").Scan(&year)
	if year != 2023 {
		t.Errorf("imported confirmed_at year %d, want 2023", year)
	}
}

func TestSubscribeOff(t *testing.T) {
	h, err := New(Config{Site: os.DirFS(site(t))})
	if err != nil {
		t.Fatal(err)
	}
	if rec := do(h, "POST", "/api/subscribe", nil, "email=a@b.co"); rec.Code != 503 {
		t.Errorf("signup without a database: %d, want 503", rec.Code)
	}
	if rec := do(h, "GET", "/unsubscribe?token=x", nil, ""); rec.Code != 503 {
		t.Errorf("unsubscribe without a database: %d, want 503", rec.Code)
	}
}
