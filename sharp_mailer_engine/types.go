package sharp_mailer_engine

import (
	mailerEngineTypes "sharpnr-engine/sharp_mailer_engine/types"
)

// Re-export mailer types so callers only import sharp_mailer_engine
type MailType = mailerEngineTypes.Type // alias, not a new type — no conversion needed

const (
	MailTypeOTP          = mailerEngineTypes.OTP
	MailTypeSupport      = mailerEngineTypes.Support
	MailTypeBilling      = mailerEngineTypes.Billing
	MailTypeNotification = mailerEngineTypes.Notification
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
// Kept unexported — callers use the typed constructors (NewOTPJob etc.),
// which bind the mail type so a job can never land on the wrong queue.
type internalEmailJob struct {
	To       string
	Subject  string
	Template TemplateKey
	Data     any
	mailType MailType
}
