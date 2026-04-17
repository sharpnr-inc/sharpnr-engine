package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// Short-lived session data from the gateway
// Separate from Payment because this data expires and has no accounting meaning
// Cashfree: payment_session_id + QR string
// Razorpay: order_id + short_url
// PhonePe:  redirect_url + transaction_id
// Pinelabs: transaction_ref + terminal_response
type PaymentSession struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID       uuid.UUID `gorm:"type:uuid;not null;index"`
	PaymentOrderID uuid.UUID `gorm:"type:uuid;not null;index"`
	ProviderCode   string    `gorm:"type:varchar(30);not null"`

	// Gateway's session/order identifier
	// Cashfree → payment_session_id
	// Razorpay → razorpay_order_id
	// PhonePe  → merchantTransactionId
	ProviderSessionID string `gorm:"type:varchar(300);not null"`

	// QR or redirect — what gets shown to customer
	// For QR flows: "upi://pay?pa=cashfree@axis&pn=SharpNR&am=850&..."
	// For redirect flows: "https://api.razorpay.com/v1/checkout/..."
	// Null for terminal/POS flows (Pinelabs handles display)
	CheckoutPayload *string `gorm:"type:text"`

	// ACTIVE, PAID, EXPIRED, CANCELLED
	Status    string    `gorm:"type:varchar(20);not null;default:'ACTIVE'"`
	ExpiresAt time.Time `gorm:"not null"`

	// Raw response snapshot from gateway — never mutate after set
	RawResponse *string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PaymentSession) TableName() string { return "payments.payment_sessions" }
