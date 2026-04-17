package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// Replaces tight internal coupling
// When standalone: caller registers their URL here
// When monolith: your own service URL is registered here during bootstrap
type WebhookEndpoint struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID uuid.UUID `gorm:"type:uuid;not null;index"`

	URL         string `gorm:"type:text;not null"`
	Description string `gorm:"type:varchar(200)"`

	// HMAC-SHA256 signing secret — engine signs payload, caller verifies
	SigningSecret string `gorm:"type:text;not null"`

	// JSON array of subscribed events
	// e.g. ["PAYMENT_SUCCESS", "PAYMENT_FAILED", "REFUND_SUCCESS"]
	// Empty = subscribed to all events
	SubscribedEvents string `gorm:"type:text;default:'[]'"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (WebhookEndpoint) TableName() string { return "payments.webhook_endpoints" }
