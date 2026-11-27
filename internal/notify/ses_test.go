package notify

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type fakeSES struct {
	in  *sesv2.SendEmailInput
	err error
}

func (f *fakeSES) SendEmail(_ context.Context, in *sesv2.SendEmailInput, _ ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error) {
	f.in = in
	return &sesv2.SendEmailOutput{MessageId: aws.String("id")}, f.err
}

func TestSESSend(t *testing.T) {
	c := &fakeSES{}
	m := &SESMailer{Client: c, From: "inquiries@notify.example.com", FromName: "site", To: "me@example.com", ConfigurationSet: "cs"}

	// An inquiry: to us, with Reply-To the inquirer.
	if err := m.Send(context.Background(), Message{Subject: "New\ninquiry", Text: "hi", ReplyTo: "ana@example.com", ReplyToName: "Ana Ö"}); err != nil {
		t.Fatal(err)
	}
	in := c.in
	if got := in.Destination.ToAddresses; len(got) != 1 || got[0] != "me@example.com" {
		t.Errorf("to %v", got)
	}
	if *in.FromEmailAddress != `"site" <inquiries@notify.example.com>` || *in.Content.Simple.Subject.Data != "New inquiry" ||
		*in.ConfigurationSetName != "cs" || in.Content.Simple.Body.Html != nil {
		t.Errorf("from %q subject %q", *in.FromEmailAddress, *in.Content.Simple.Subject.Data)
	}
	if rt := in.ReplyToAddresses; len(rt) != 1 || !strings.Contains(rt[0], "<ana@example.com>") || !strings.HasPrefix(rt[0], "=?utf-8?") {
		t.Errorf("reply-to %v (a non-ASCII name must be encoded)", rt)
	}

	// A newsletter email: its own recipient, sender, HTML and headers.
	err := m.Send(context.Background(), Message{Subject: "Post", Text: "t", HTML: "<p>t</p>", To: "ana@example.com",
		From: "newsletter@notify.example.com", FromName: "Sawir Studio",
		Headers: map[string]string{"List-Unsubscribe-Post": "List-Unsubscribe=One-Click"}})
	if err != nil {
		t.Fatal(err)
	}
	in = c.in
	if in.Destination.ToAddresses[0] != "ana@example.com" || *in.FromEmailAddress != `"Sawir Studio" <newsletter@notify.example.com>` ||
		*in.Content.Simple.Body.Html.Data != "<p>t</p>" || len(in.Content.Simple.Headers) != 1 ||
		*in.Content.Simple.Headers[0].Name != "List-Unsubscribe-Post" {
		t.Errorf("newsletter request: %+v", in)
	}

	c.err = errors.New("Throttling")
	if err := m.Send(context.Background(), Message{Subject: "s", Text: "t"}); err == nil || errors.Is(err, ErrNotDelivered) {
		t.Errorf("a failed call should be an error, and not a per-recipient one: %v", err)
	}
}
