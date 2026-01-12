package sharp_mailer_engine

import "context"

type MailType string

const (
	OTP          MailType = "otp"
	Billing      MailType = "billing"
	Notification MailType = "notification"
	Support      MailType = "support"
)

type TemplateKey string

const (
	TmplVerification TemplateKey = "otp.html"
	TmplInvoice      TemplateKey = "billing.html"
	TmplWelcome      TemplateKey = "welcome.html"
)

type EmailJob struct {
	To       string
	Subject  string
	Template TemplateKey
	Data     interface{}
	From     string
	Ctx      context.Context
}
