package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Turnstile checks the token Cloudflare Turnstile's widget puts in the contact form
// (cf-turnstile-response) with Cloudflare's siteverify API.
type Turnstile struct {
	Secret string
	URL    string       // default siteverifyURL; tests point it at a fake
	Client *http.Client // default: a 5-second timeout
}

const siteverifyURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

// errTurnstileDown means Cloudflare didn't give an answer (timeout, 5xx, a reply we can't read).
var errTurnstileDown = errors.New("turnstile: no answer from siteverify")

// Verify reports whether Cloudflare accepts token, sent from ip ("" when unknown). An error
// wrapping errTurnstileDown means Cloudflare couldn't be asked, not that the token is bad.
func (t *Turnstile) Verify(ctx context.Context, token, ip string) (bool, error) {
	if token == "" {
		return false, nil
	}
	form := url.Values{"secret": {t.Secret}, "response": {token}}
	if ip != "" {
		form.Set("remoteip", ip)
	}
	endpoint, client := t.URL, t.Client
	if endpoint == "" {
		endpoint = siteverifyURL
	}
	if client == nil {
		client = &http.Client{Timeout: 5 * time.Second}
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("%w: %v", errTurnstileDown, err)
	}
	defer res.Body.Close()
	var out struct {
		Success    bool     `json:"success"`
		ErrorCodes []string `json:"error-codes"`
	}
	if res.StatusCode >= 500 {
		return false, fmt.Errorf("%w: status %d", errTurnstileDown, res.StatusCode)
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return false, fmt.Errorf("%w: %v", errTurnstileDown, err)
	}
	for _, c := range out.ErrorCodes {
		// Our own setup is wrong (bad secret, Cloudflare's own error): don't blame the visitor.
		if c == "missing-input-secret" || c == "invalid-input-secret" || c == "internal-error" {
			return false, fmt.Errorf("%w: %s", errTurnstileDown, c)
		}
	}
	return out.Success, nil
}
