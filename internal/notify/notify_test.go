package notify

import (
	"context"
	"encoding/json"
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

func TestSendErrors(t *testing.T) {
	for _, body := range []string{
		`{"success":false,"errors":[{"code":10000,"message":"Authentication error"}]}`,
		`{"success":true,"errors":[],"result":{"delivered":[],"permanent_bounces":["me@example.com"]}}`,
		`not json`,
	} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, body) }))
		m := &Mailer{AccountID: "a", Token: "t", From: "f@x.com", To: "me@example.com", BaseURL: srv.URL}
		if err := m.Send(context.Background(), Message{Subject: "s", Text: "t"}); err == nil || !strings.Contains(err.Error(), "cloudflare email") {
			t.Errorf("%s: err = %v", body, err)
		}
		srv.Close()
	}
}
