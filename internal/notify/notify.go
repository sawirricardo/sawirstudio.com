// Package notify sends email through Cloudflare's Email Service REST API
// (POST /accounts/{account_id}/email/sending/send). The sender is an address on a domain
// with Cloudflare Email Routing (notify.sawirstudio.com), which may send to the account's
// verified destination addresses.
package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Mailer sends email from one address to one address.
type Mailer struct {
	AccountID string
	Token     string // an API token with Email Sending permission
	From      string // like inquiries@notify.sawirstudio.com
	FromName  string
	To        string // a verified destination address
	BaseURL   string // the Cloudflare API, overridable in tests
	Client    *http.Client
}

// FromEnv returns a Mailer from CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_EMAIL_TOKEN, NOTIFY_FROM and
// NOTIFY_TO, or nil and the names of the missing ones.
func FromEnv() (*Mailer, []string) {
	m := &Mailer{
		AccountID: os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		Token:     os.Getenv("CLOUDFLARE_EMAIL_TOKEN"),
		From:      os.Getenv("NOTIFY_FROM"),
		FromName:  "sawirstudio.com",
		To:        os.Getenv("NOTIFY_TO"),
	}
	var missing []string
	for name, v := range map[string]string{"CLOUDFLARE_ACCOUNT_ID": m.AccountID, "CLOUDFLARE_EMAIL_TOKEN": m.Token, "NOTIFY_FROM": m.From, "NOTIFY_TO": m.To} {
		if v == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return nil, missing
	}
	return m, nil
}

// Message is one email. ReplyTo, when set, is where a reply goes.
type Message struct {
	Subject, Text string
	ReplyTo       string
	ReplyToName   string
}

type sendRequest struct {
	To      string   `json:"to"`
	From    address  `json:"from"`
	ReplyTo *address `json:"reply_to,omitempty"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
}

type address struct {
	Address string `json:"address"`
	Name    string `json:"name,omitempty"`
}

type sendResponse struct {
	Success bool `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Result struct {
		Delivered []string `json:"delivered"`
		Queued    []string `json:"queued"`
		Bounces   []string `json:"permanent_bounces"`
		Dropped   []string `json:"suppressed_recipients"`
	} `json:"result"`
}

// Send sends msg, and reports an error unless Cloudflare delivered or queued it.
func (m *Mailer) Send(ctx context.Context, msg Message) error {
	req := sendRequest{
		To:      m.To,
		From:    address{m.From, m.FromName},
		Subject: oneLine(msg.Subject),
		Text:    msg.Text,
	}
	if msg.ReplyTo != "" {
		req.ReplyTo = &address{msg.ReplyTo, oneLine(msg.ReplyToName)}
	}
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	base := m.BaseURL
	if base == "" {
		base = "https://api.cloudflare.com/client/v4"
	}
	hr, err := http.NewRequestWithContext(ctx, http.MethodPost, base+"/accounts/"+m.AccountID+"/email/sending/send", bytes.NewReader(body))
	if err != nil {
		return err
	}
	hr.Header.Set("Authorization", "Bearer "+m.Token)
	hr.Header.Set("Content-Type", "application/json")
	client := m.Client
	if client == nil {
		client = &http.Client{Timeout: 20 * time.Second}
	}
	res, err := client.Do(hr)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(res.Body, 64<<10))
	var out sendResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return fmt.Errorf("cloudflare email: HTTP %d: %.200s", res.StatusCode, raw)
	}
	if !out.Success {
		var msgs []string
		for _, e := range out.Errors {
			msgs = append(msgs, fmt.Sprintf("%d %s", e.Code, e.Message))
		}
		return fmt.Errorf("cloudflare email: HTTP %d: %s", res.StatusCode, strings.Join(msgs, "; "))
	}
	if len(out.Result.Delivered)+len(out.Result.Queued) == 0 {
		return errors.New("cloudflare email: not delivered (bounced or suppressed)")
	}
	return nil
}

func oneLine(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
