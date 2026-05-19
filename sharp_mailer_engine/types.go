package sharp_mailer_engine

import (
	"context"
	mailer "sharpnr-engine/sharp_mailer"
)

// Re-export mailer types so callers only import sharp_mailer_engine
type MailType = mailer.Type // alias, not a new type — no conversion needed

const (
	MailTypeOTP          = mailer.OTP
	MailTypeSupport      = mailer.Support
	MailTypeBilling      = mailer.Billing
	MailTypeNotification = mailer.Notification
)

// TemplateKey is the filename of the Go HTML template
type TemplateKey string

const (
	TmplVerification TemplateKey = "otp.html"
	TmplInvoice      TemplateKey = "billing.html"
	TmplWelcome      TemplateKey = "welcome.html"
	TmplNotification TemplateKey = "notifications.html"
)

// internalEmailJob is what flows through the worker queue channels.
// Kept unexported — callers use the typed constructors (NewOTPJob etc.)
type internalEmailJob struct {
	To       string
	Subject  string
	Template TemplateKey
	Data     any
	From     string
	Ctx      context.Context
}