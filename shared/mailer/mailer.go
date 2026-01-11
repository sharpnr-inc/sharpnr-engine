package mailer

import (
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mailer struct {
	client     *mail.SMTPClient
	identities map[Type]Identity
}

func New(
	client *mail.SMTPClient,
	identities map[Type]Identity,
) *Mailer {
	return &Mailer{
		client:     client,
		identities: identities,
	}
}

func (m *Mailer) Send(
	t Type,
	to string,
	subject string,
	body string,
) error {

	identity, ok := m.identities[t]
	if !ok {
		return ErrUnknownEmailType
	}

	email := mail.NewMSG()
	email.SetFrom(identity.FromAddress)
	email.AddTo(to)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	return email.Send(m.client)
}
