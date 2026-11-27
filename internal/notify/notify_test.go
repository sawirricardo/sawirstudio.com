package notify

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSend(t *testing.T) {
	var got map[string]any
	var auth, path string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, path = r.Header.Get("Authorization"), r.URL.Path
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &got)
		io.WriteString(w, `{"success":true,"errors":[],"result":{"delivered":["me@example.com"]}}`)
	}))
	defer srv.Close()
	m := &Mailer{AccountID: "acc", Token: "tok", From: "inquiries@notify.example.com", FromName: "site",
		To: "me@example.com", BaseURL: srv.URL}
	err := m.Send(context.Background(), Message{Subject: "New\ninquiry", Text: "hi", ReplyTo: "ana@example.com", ReplyToName: "Ana"})
	if err != nil {
		t.Fatal(err)
	}
	if auth != "Bearer tok" || path != "/accounts/acc/email/sending/send" {
		t.Errorf("auth %q path %q", auth, path)
	}
	if got["subject"] != "New inquiry" || got["to"] != "me@example.com" {
		t.Errorf("subject %q to %q", got["subject"], got["to"])
	}
	if rt, _ := got["reply_to"].(map[string]any); rt["address"] != "ana@example.com" {
		t.Errorf("reply_to %v", got["reply_to"])
	}
}

func TestSendToSeveral(t *testing.T) {
	var to []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var got map[string]any
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &got)
		to = append(to, got["to"].(string))
		io.WriteString(w, `{"success":true,"errors":[],"result":{"delivered":["x"]}}`)
	}))
	defer srv.Close()
	m := &Mailer{AccountID: "acc", Token: "tok", From: "f@x.com", To: "me@example.com, you@example.com", BaseURL: srv.URL}
	if err := m.Send(context.Background(), Message{Subject: "s", Text: "t"}); err != nil {
		t.Fatal(err)
	}
	if strings.Join(to, " ") != "me@example.com you@example.com" {
		t.Errorf("sent to %v, want one email to each", to)
	}
}

func TestSendTo(t *testing.T) {
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &got)
		io.WriteString(w, `{"success":true,"errors":[],"result":{"queued":["ana@example.com"]}}`)
	}))
	defer srv.Close()
	m := &Mailer{AccountID: "acc", Token: "tok", From: "inquiries@notify.example.com", FromName: "site",
		To: "me@example.com", BaseURL: srv.URL}
	err := m.Send(context.Background(), Message{Subject: "Confirm", Text: "hi", HTML: "<p>hi</p>",
		To: "ana@example.com", From: "newsletter@notify.example.com", FromName: "Sawir Studio",
		Headers: map[string]string{"List-Unsubscribe": "<https://example.com/u>"}})
	if err != nil {
		t.Fatal(err)
	}
	from, _ := got["from"].(map[string]any)
	headers, _ := got["headers"].(map[string]any)
	if got["to"] != "ana@example.com" || from["address"] != "newsletter@notify.example.com" || from["name"] != "Sawir Studio" ||
		got["html"] != "<p>hi</p>" || headers["List-Unsubscribe"] != "<https://example.com/u>" {
		t.Errorf("request %v", got)
	}
}

func TestSendErrors(t *testing.T) {
	for _, body := range []string{
		`{"success":false,"errors":[{"code":10000,"message":"Authentication error"}]}`,
		`{"success":true,"errors":[],"result":{"delivered":[],"permanent_bounces":["me@example.com"]}}`,
		`not json`,
	} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, body) }))
		m := &Mailer{AccountID: "a", Token: "t", From: "f@x.com", To: "me@example.com", BaseURL: srv.URL}
		err := m.Send(context.Background(), Message{Subject: "s", Text: "t"})
		if err == nil || !strings.Contains(err.Error(), "cloudflare email") {
			t.Errorf("%s: err = %v", body, err)
		}
		if bounced := strings.Contains(body, "permanent_bounces"); errors.Is(err, ErrNotDelivered) != bounced {
			t.Errorf("%s: errors.Is(err, ErrNotDelivered) should be %v", body, bounced)
		}
		srv.Close()
	}
}
