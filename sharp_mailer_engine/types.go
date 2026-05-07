package sharp_mailer_engine

import "context"

type TemplateData interface {
	OTPData | BillingData | NotificationData | WelcomeData
}

type OTPData struct {
	Name          string
	Code          string
	ExpiryMinutes int
}

type BillingData struct {
	Name          string
	Amount        string
	Date          string
	Plan          string
	BillingPeriod string // optional
	Link          string // optional
}

type NotificationData struct {
	Title      string
	Subtitle   string // optional
	Message    string
	ActionURL  string // optional
	ActionText string // optional
	FooterNote string // optional
}

type WelcomeData struct {
	Name             string
	Email            string
	OrganizationName string // optional
	ActionURL        string
	ActionText       string
}

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
	TmplNotification TemplateKey = "notifications.html"
)

type EmailJob struct {
	To       string
	Subject  string
	Template TemplateKey
	Data     interface{}
	From     string
	Ctx      context.Context
}
