package server

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/sawirricardo/sawirstudio.com/internal/notify"
)

func site(t *testing.T) string {
	root := t.TempDir()
	files := map[string]string{
		"index.html":          "<html>home</html>",
		"index.md":            "# Home\n",
		"about-us/index.html": "<html>about</html>",
		"about-us/index.md":   "# About\n",
		"404.html":            "<html>not found</html>",
		"ja/index.html":       "<html>ja</html>",
		"ja/404.html":         "<html>ja not found</html>",
		"assets/a.css":        "body{}",
		"README.md":           "secret",
		"scripts/x.py":        "print()",
		".env":                "DATABASE_URL=x",
		".assetsignore":       "README.md\nscripts\n",
		"_redirects":          "# old\n/projects /projects/ 301\n/about /about-us/ 301\n",
	}
	for name, body := range files {
		p := filepath.Join(root, name)
		os.MkdirAll(filepath.Dir(p), 0o755)
		os.WriteFile(p, []byte(body), 0o644)
	}
	return root
}

func do(h http.Handler, method, target string, hdr map[string]string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	for k, v := range hdr {
		req.Header.Set(k, v)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestStatic(t *testing.T) {
	h, err := New(Config{Site: os.DirFS(site(t)), TrustProxy: true})
	if err != nil {
		t.Fatal(err)
	}
	rec := do(h, "GET", "/", nil, "")
	if rec.Code != 200 || rec.Body.String() != "<html>home</html>" {
		t.Fatalf("home: %d %q", rec.Code, rec.Body)
	}
	links := strings.Join(rec.Header().Values("Link"), ", ")
	for _, want := range []string{"</index.md>", "</llms.txt>", "</sitemap.xml>"} {
		if !strings.Contains(links, want) {
			t.Errorf("home Link headers %q are missing %s", links, want)
		}
	}
	if rec.Header().Get("Vary") != "Accept" {
		t.Error("pages should vary on Accept")
	}
	if rec.Header().Get("Cache-Control") != pageCache {
		t.Errorf("page Cache-Control = %q", rec.Header().Get("Cache-Control"))
	}
	etag := rec.Header().Get("ETag")
	if rec = do(h, "GET", "/", map[string]string{"If-None-Match": etag}, ""); etag == "" || rec.Code != 304 {
		t.Errorf("revalidating with ETag %q: %d, want 304", etag, rec.Code)
	}

	rec = do(h, "GET", "/about-us/", map[string]string{"Accept": "text/markdown, text/html;q=0.9"}, "")
	if rec.Body.String() != "# About\n" || !strings.HasPrefix(rec.Header().Get("Content-Type"), "text/markdown") {
		t.Errorf("markdown: %q %q", rec.Header().Get("Content-Type"), rec.Body)
	}
	if rec.Header().Get("Cache-Control") != "private, no-store" {
		t.Errorf("markdown must not be cached by Cloudflare: %q", rec.Header().Get("Cache-Control"))
	}
	if rec.Header().Get("x-markdown-tokens") != "2" {
		t.Errorf("x-markdown-tokens = %q", rec.Header().Get("x-markdown-tokens"))
	}

	for target, want := range map[string]string{"/about-us": "/about-us/", "/projects": "/projects/", "/about": "/about-us/"} {
		rec = do(h, "GET", target, nil, "")
		if rec.Code != 301 || rec.Header().Get("Location") != want {
			t.Errorf("%s: %d -> %q, want 301 -> %q", target, rec.Code, rec.Header().Get("Location"), want)
		}
	}

	for _, target := range []string{"/README.md", "/scripts/x.py", "/.env", "/.assetsignore", "/_redirects", "/nope/", "/etc/passwd"} {
		rec = do(h, "GET", target, nil, "")
		if rec.Code != 404 || rec.Body.String() != "<html>not found</html>" {
			t.Errorf("%s: %d %q, want the 404 page", target, rec.Code, rec.Body)
		}
	}

	for target, want := range map[string]string{"/ja/nope/": "<html>ja not found</html>", "/ja/a/b/c": "<html>ja not found</html>", "/jax/": "<html>not found</html>"} {
		if rec = do(h, "GET", target, nil, ""); rec.Code != 404 || rec.Body.String() != want {
			t.Errorf("%s: %d %q, want %q", target, rec.Code, rec.Body, want)
		}
	}

	rec = do(h, "GET", "/assets/a.css", nil, "")
	if rec.Code != 200 || rec.Header().Get("Cache-Control") != assetCache {
		t.Errorf("asset: %d %q", rec.Code, rec.Header().Get("Cache-Control"))
	}

	rec = do(h, "POST", "/api/inquiries", nil, "")
	if rec.Header().Get("Cache-Control") != "no-store" {
		t.Errorf("API Cache-Control = %q", rec.Header().Get("Cache-Control"))
	}

	rec = do(h, "GET", "/about-us/?x=1", map[string]string{"X-Forwarded-Proto": "http"}, "")
	if rec.Code != 301 || rec.Header().Get("Location") != "https://example.com/about-us/?x=1" {
		t.Errorf("https redirect: %d %q", rec.Code, rec.Header().Get("Location"))
	}

	req := httptest.NewRequest("GET", "/posts/?a=1", nil)
	req.Host = "www.sawirstudio.com"
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != 301 || rec.Header().Get("Location") != "https://sawirstudio.com/posts/?a=1" {
		t.Errorf("www: %d %q", rec.Code, rec.Header().Get("Location"))
	}

	rec = do(h, "POST", "/api/inquiries", nil, "name=x")
	if rec.Code != 503 {
		t.Errorf("inquiry without a database: %d, want 503", rec.Code)
	}
}

// TestInquiries needs a migrated database: TEST_DATABASE_URL=postgres://... (see README).
func TestInquiries(t *testing.T) {
	url_ := os.Getenv("TEST_DATABASE_URL")
	if url_ == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, url_)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	if _, err := pool.Exec(ctx, "truncate inquiries"); err != nil {
		t.Fatal(err)
	}
	mail := &fakeMailer{}
	inq := NewInquiries(pool, mail)
	h, err := New(Config{Site: os.DirFS(site(t)), Inquiries: inq})
	if err != nil {
		t.Fatal(err)
	}
	post := func(form url.Values, accept string) *httptest.ResponseRecorder {
		return do(h, "POST", "/api/inquiries", map[string]string{
			"Content-Type": "application/x-www-form-urlencoded", "Accept": accept,
		}, form.Encode())
	}
	good := url.Values{"name": {"Ana"}, "email": {"ana@example.com"}, "company": {"Acme"},
		"need": {"Website", "Design"}, "message": {"We need a site."}}

	if rec := post(good, "application/json"); rec.Code != 201 || !strings.Contains(rec.Body.String(), `"ok":true`) {
		t.Fatalf("good inquiry: %d %s", rec.Code, rec.Body)
	}
	var name, company string
	var needs []string
	if err := pool.QueryRow(ctx, "select name, company, needs from inquiries").Scan(&name, &company, &needs); err != nil {
		t.Fatal(err)
	}
	if name != "Ana" || company != "Acme" || strings.Join(needs, ",") != "Website,Design" {
		t.Errorf("stored %q %q %v", name, company, needs)
	}

	inq.Wait()
	if len(mail.sent) != 1 {
		t.Fatalf("%d emails, want 1", len(mail.sent))
	}
	m := mail.sent[0]
	if m.Subject != "New inquiry from Ana (Acme)" || m.ReplyTo != "ana@example.com" ||
		!strings.Contains(m.Text, "We need a site.") || !strings.Contains(m.Text, "Website, Design") {
		t.Errorf("email: %+v", m)
	}
	var notified bool
	pool.QueryRow(ctx, "select notified_at is not null from inquiries").Scan(&notified)
	if !notified {
		t.Error("notified_at should be set after the email went out")
	}

	// Without JavaScript the browser gets a page back.
	if rec := post(good, "text/html"); rec.Code != 201 || !strings.Contains(rec.Body.String(), "Thanks!") {
		t.Errorf("html reply: %d %s", rec.Code, rec.Body)
	}

	bad := url.Values{"name": {"Ana"}, "email": {"not an email"}, "message": {"hi"}}
	if rec := post(bad, "application/json"); rec.Code != 422 {
		t.Errorf("bad email: %d", rec.Code)
	}
	odd := url.Values{"name": {"Ana"}, "email": {"a@b.co"}, "message": {"hi"}, "need": {"<script>"}}
	if rec := post(odd, "application/json"); rec.Code != 422 {
		t.Errorf("unknown need: %d", rec.Code)
	}

	bot := url.Values{"name": {"Bot"}, "email": {"b@b.co"}, "message": {"buy"}, "website": {"http://spam"}}
	if rec := post(bot, "application/json"); rec.Code != 200 {
		t.Errorf("honeypot: %d, want a quiet 200", rec.Code)
	}

	inq.Wait()
	if len(mail.sent) != 2 {
		t.Errorf("%d emails, want 2 (none for bots or invalid inquiries)", len(mail.sent))
	}
	// An inquiry whose email didn't go out is sent at the next start.
	pool.Exec(ctx, "update inquiries set notified_at = null where email = 'ana@example.com' and id = (select min(id) from inquiries)")
	inq.SendPending(ctx, slog.New(slog.DiscardHandler))
	inq.Wait()
	if len(mail.sent) != 3 {
		t.Errorf("%d emails after SendPending, want 3", len(mail.sent))
	}
	mail.sent = mail.sent[:2]

	var n int
	pool.QueryRow(ctx, "select count(*) from inquiries").Scan(&n)
	if n != 2 {
		t.Errorf("%d inquiries stored, want 2 (bot and invalid ones dropped)", n)
	}

	// The same address can send rateLimit inquiries per window.
	for i := n; i < rateLimit; i++ {
		post(good, "application/json")
	}
	if rec := post(good, "application/json"); rec.Code != 429 {
		t.Errorf("over the limit: %d, want 429", rec.Code)
	}
}

func TestNoSite(t *testing.T) {
	if _, err := New(Config{Site: os.DirFS(t.TempDir())}); err == nil {
		t.Error("a site without index.html should be refused")
	}
}

type fakeMailer struct {
	mu   sync.Mutex
	sent []notify.Message
}

func (f *fakeMailer) Send(_ context.Context, m notify.Message) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sent = append(f.sent, m)
	return nil
}
