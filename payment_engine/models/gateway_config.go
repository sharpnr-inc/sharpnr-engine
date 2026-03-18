package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// One row per gateway per tenant
// Provider-specific keys (SaltKey, KeySecret etc.) live in ConfigEncrypted as JSON
// Never add provider-specific columns here
type PaymentGatewayConfig struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID uuid.UUID `gorm:"type:uuid;not null;index"`

	ProviderCode string `gorm:"type:varchar(30);not null"` // PHONEPE, RAZORPAY, CASHFREE, STRIPE
	DisplayName  string `gorm:"type:varchar(100);not null"`

	// Encrypted JSON — shape is provider-specific
	// PhonePe:  { "merchant_id": "", "salt_key": "", "salt_index": "" }
	// Razorpay: { "key_id": "", "key_secret": "" }
	// Stripe:   { "publishable_key": "", "secret_key": "", "webhook_secret": "" }
	ConfigEncrypted string `gorm:"type:text;not null"`
	WebhookURL      string `gorm:"type:text;not null"` // ← add

	// Sandbox vs Production
	IsTestMode bool `gorm:"default:false"`
	IsActive   bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PaymentGatewayConfig) TableName() string { return "payments.gateway_configs" }
