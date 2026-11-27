// Package notify sends the site's email: inquiry notifications and the newsletter. It goes
// through Amazon SES (SESMailer, when SES_REGION is set) or Cloudflare's Email Service REST
// API (Mailer, POST /accounts/{account_id}/email/sending/send). Cloudflare's sender is an
// address on a domain with Email Routing (notify.sawirstudio.com); on the Workers Free plan
// it may send only to the account's verified destination addresses, so it can't send the
// newsletter's confirmation emails to visitors. SES (production access) can.
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

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

// Mailer sends email from one address, to one address unless a Message says otherwise.
type Mailer struct {
	AccountID string
	Token     string // an API token with Email Sending permission
	From      string // like inquiries@notify.sawirstudio.com
	FromName  string
	To        string // verified destination addresses, separated by commas
	BaseURL   string // the Cloudflare API, overridable in tests
	Client    *http.Client
}

// Sender sends one email (Mailer, SESMailer).
type Sender interface {
	Send(ctx context.Context, msg Message) error
}

// FromEnv returns the sender the environment sets up, or nil and the names of the missing
// settings. Both need NOTIFY_FROM (the sender) and NOTIFY_TO (where inquiry emails go: one
// address, or several separated by commas). With
// SES_REGION set it's Amazon SES, with the AWS SDK's usual credentials (AWS_ACCESS_KEY_ID and
// AWS_SECRET_ACCESS_KEY) and optionally SES_CONFIGURATION_SET; otherwise Cloudflare, with
// CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_EMAIL_TOKEN.
func FromEnv(ctx context.Context) (Sender, []string) {
	from, to := os.Getenv("NOTIFY_FROM"), os.Getenv("NOTIFY_TO")
	if region := os.Getenv("SES_REGION"); region != "" {
		if missing := missingVars("NOTIFY_FROM", "NOTIFY_TO"); len(missing) > 0 {
			return nil, missing
		}
		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
		if err != nil {
			return nil, []string{"AWS credentials (" + err.Error() + ")"}
		}
		return &SESMailer{Client: sesv2.NewFromConfig(cfg), From: from, FromName: "sawirstudio.com", To: to,
			ConfigurationSet: os.Getenv("SES_CONFIGURATION_SET")}, nil
	}
	if missing := missingVars("CLOUDFLARE_ACCOUNT_ID", "CLOUDFLARE_EMAIL_TOKEN", "NOTIFY_FROM", "NOTIFY_TO"); len(missing) > 0 {
		return nil, missing
	}
	return &Mailer{AccountID: os.Getenv("CLOUDFLARE_ACCOUNT_ID"), Token: os.Getenv("CLOUDFLARE_EMAIL_TOKEN"),
		From: from, FromName: "sawirstudio.com", To: to}, nil
}

// Addresses splits a comma-separated list of addresses (NOTIFY_TO).
func Addresses(list string) []string {
	var out []string
	for _, a := range strings.Split(list, ",") {
		if a = strings.TrimSpace(a); a != "" {
			out = append(out, a)
		}
	}
	return out
}

func missingVars(names ...string) []string {
	var missing []string
	for _, n := range names {
		if os.Getenv(n) == "" {
			missing = append(missing, n)
		}
	}
	return missing
}

// ErrNotDelivered is the error Send wraps when Cloudflare took the email but didn't deliver
// it: the address bounced, or it's on the suppression list. Retrying won't help.
var ErrNotDelivered = errors.New("cloudflare email: not delivered (bounced or suppressed)")

// Message is one email. ReplyTo, when set, is where a reply goes. To and From, when set,
// replace the Mailer's; HTML is an HTML version of Text; Headers are extra headers, from
// Cloudflare's allowed list (List-Unsubscribe and the like).
type Message struct {
	Subject, Text string
	HTML          string
	ReplyTo       string
	ReplyToName   string
	To            string
	From          string
	FromName      string
	Headers       map[string]string
}

type sendRequest struct {
	To      string            `json:"to"`
	From    address           `json:"from"`
	ReplyTo *address          `json:"reply_to,omitempty"`
	Subject string            `json:"subject"`
	Text    string            `json:"text"`
	HTML    string            `json:"html,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
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

// Send sends msg, and reports an error unless Cloudflare delivered or queued it. Without its
// own To it goes to each of the Mailer's addresses, one email each.
func (m *Mailer) Send(ctx context.Context, msg Message) error {
	if msg.To == "" {
		var errs []error
		for _, to := range Addresses(m.To) {
			msg.To = to
			errs = append(errs, m.Send(ctx, msg))
		}
		return errors.Join(errs...)
	}
	req := sendRequest{
		To:      m.To,
		From:    address{m.From, m.FromName},
		Subject: oneLine(msg.Subject),
		Text:    msg.Text,
		HTML:    msg.HTML,
		Headers: msg.Headers,
	}
	if msg.To != "" {
		req.To = msg.To
	}
	if msg.From != "" {
		req.From = address{msg.From, oneLine(msg.FromName)}
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
		return ErrNotDelivered
	}
	return nil
}

func oneLine(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
