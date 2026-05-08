package sharp_mailer_engine

// Each template has its own struct.
// Add fields here when the HTML template needs new variables.

type OTPData struct {
	Name          string
	Email         string
	Code          string
	ExpiryMinutes int
}

type BillingData struct {
	Name          string
	Amount        string
	Date          string
	Plan          string
	BillingPeriod string // optional — renders only if non-empty
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
