package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// Full audit trail of every outbound webhook attempt
// Enables retry, debugging, and delivery guarantees
type WebhookDeliveryLog struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	EndpointID uuid.UUID `gorm:"type:uuid;not null;index"`

	// What triggered this delivery
	EventType string `gorm:"type:varchar(50);not null;index"`  // PAYMENT_SUCCESS, REFUND_INITIATED...
	EventID   string `gorm:"type:varchar(100);not null;index"` // PaymentID, RefundID etc.

	// Full signed payload that was sent
	Payload string `gorm:"type:text;not null"`

	// PENDING, SUCCESS, FAILED
	Status     string `gorm:"type:varchar(20);not null;default:'PENDING';index"`
	AttemptNum int    `gorm:"default:1"`

	// HTTP response from the caller's endpoint
	ResponseCode    *int    `gorm:"type:int"`
	ResponseBody    *string `gorm:"type:text"`
	ResponseLatency *int    `gorm:"type:int"` // milliseconds

	ErrorMessage *string    `gorm:"type:text"`
	AttemptedAt  time.Time  `gorm:"not null"`
	NextRetryAt  *time.Time `gorm:"type:timestamp;index"` // null = no retry scheduled
}

func (WebhookDeliveryLog) TableName() string { return "payments.webhook_delivery_logs" }
