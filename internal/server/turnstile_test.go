package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

// fakeSiteverify answers like Cloudflare: token "good" passes, "down" is a 503, others fail.
func fakeSiteverify(t *testing.T) *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.PostForm.Get("secret") != "s3cret" {
			fmt.Fprint(w, `{"success":false,"error-codes":["invalid-input-secret"]}`)
			return
		}
		switch r.PostForm.Get("response") {
		case "good":
			fmt.Fprint(w, `{"success":true,"error-codes":[]}`)
		case "down":
			w.WriteHeader(http.StatusServiceUnavailable)
		default:
			fmt.Fprint(w, `{"success":false,"error-codes":["invalid-input-response"]}`)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestTurnstileVerify(t *testing.T) {
	srv := fakeSiteverify(t)
	ts := &Turnstile{Secret: "s3cret", URL: srv.URL}
	for _, c := range []struct {
		token    string
		ok, down bool
	}{
		{"", false, false},
		{"good", true, false},
		{"forged", false, false},
		{"down", false, true},
	} {
		ok, err := ts.Verify(context.Background(), c.token, "203.0.113.7")
		if ok != c.ok || errors.Is(err, errTurnstileDown) != c.down {
			t.Errorf("Verify(%q) = %v, %v; want %v, down %v", c.token, ok, err, c.ok, c.down)
		}
	}
	// A wrong secret is our mistake, not the visitor's.
	if _, err := (&Turnstile{Secret: "wrong", URL: srv.URL}).Verify(context.Background(), "good", ""); !errors.Is(err, errTurnstileDown) {
		t.Errorf("wrong secret: %v, want errTurnstileDown", err)
	}
	// Nothing listening: Cloudflare can't be reached.
	if _, err := (&Turnstile{Secret: "s3cret", URL: "http://127.0.0.1:1"}).Verify(context.Background(), "good", ""); !errors.Is(err, errTurnstileDown) {
		t.Errorf("unreachable: %v, want errTurnstileDown", err)
	}
}

// TestInquiryTurnstile needs a migrated database, like TestInquiries.
func TestInquiryTurnstile(t *testing.T) {
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("TEST_DATABASE_URL is not set")
	}
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	if _, err := pool.Exec(ctx, "truncate inquiries"); err != nil {
		t.Fatal(err)
	}
	mail := &fakeMailer{}
	inq := NewInquiries(pool, mail)
	h, err := New(Config{Site: os.DirFS(site(t)), Inquiries: inq, Turnstile: &Turnstile{Secret: "s3cret", URL: fakeSiteverify(t).URL}})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		token string
		code  int
	}{
		{"", 422},       // no widget token: a script posting straight to the API
		{"forged", 422}, // Cloudflare says no
		{"good", 201},
		{"down", 201}, // Cloudflare can't answer: keep the lead
	} {
		form := url.Values{"name": {"Ana"}, "email": {"ana@example.com"}, "need": {"Website"},
			"message": {"We need a site."}, "cf-turnstile-response": {c.token}}
		rec := do(h, "POST", "/api/inquiries", map[string]string{
			"Content-Type": "application/x-www-form-urlencoded", "Accept": "application/json",
		}, form.Encode())
		if rec.Code != c.code {
			t.Errorf("token %q: %d %s, want %d", c.token, rec.Code, rec.Body, c.code)
		}
	}
	inq.Wait()
	var n int
	pool.QueryRow(ctx, "select count(*) from inquiries").Scan(&n)
	if n != 2 || len(mail.sent) != 2 {
		t.Errorf("%d stored and %d emailed, want 2 and 2 (good and down)", n, len(mail.sent))
	}
}
