package notify

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

// SESClient is the part of the SES v2 client SESMailer uses (a fake in tests).
type SESClient interface {
	SendEmail(ctx context.Context, in *sesv2.SendEmailInput, opts ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error)
}

// SESMailer sends email through Amazon SES (v2 SendEmail). The From address must be on a
// verified SES identity (sawirstudio.com is, in ap-southeast-1, and covers its subdomains).
// Addresses that bounced or complained are on the account's suppression list: SES accepts
// the email and doesn't deliver it, so Send doesn't report them.
type SESMailer struct {
	Client           SESClient
	From             string
	FromName         string
	To               string // where a Message without its own To goes; addresses separated by commas
	ConfigurationSet string // optional, for SES event publishing
}

func (m *SESMailer) Send(ctx context.Context, msg Message) error {
	to := Addresses(m.To)
	if msg.To != "" {
		to = []string{msg.To}
	}
	from := mail.Address{Name: m.FromName, Address: m.From}
	if msg.From != "" {
		from = mail.Address{Name: oneLine(msg.FromName), Address: msg.From}
	}
	utf8 := aws.String("UTF-8")
	body := &types.Body{Text: &types.Content{Data: aws.String(msg.Text), Charset: utf8}}
	if msg.HTML != "" {
		body.Html = &types.Content{Data: aws.String(msg.HTML), Charset: utf8}
	}
	content := &types.Message{
		Subject: &types.Content{Data: aws.String(oneLine(msg.Subject)), Charset: utf8},
		Body:    body,
	}
	for name, value := range msg.Headers {
		content.Headers = append(content.Headers, types.MessageHeader{Name: aws.String(name), Value: aws.String(value)})
	}
	in := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from.String()),
		Destination:      &types.Destination{ToAddresses: to},
		Content:          &types.EmailContent{Simple: content},
	}
	if msg.ReplyTo != "" {
		in.ReplyToAddresses = []string{(&mail.Address{Name: oneLine(msg.ReplyToName), Address: msg.ReplyTo}).String()}
	}
	if m.ConfigurationSet != "" {
		in.ConfigurationSetName = aws.String(m.ConfigurationSet)
	}
	if _, err := m.Client.SendEmail(ctx, in); err != nil {
		return fmt.Errorf("ses: %w", err)
	}
	return nil
}
