package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type Refund struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID  uuid.UUID `gorm:"type:uuid;not null;index"`
	PaymentID uuid.UUID `gorm:"type:uuid;not null;index"`

	RefundNumber string `gorm:"type:varchar(50);not null;uniqueIndex:idx_refund_number_tenant"`

	CurrencyCode string  `gorm:"type:char(3);not null;default:'INR'"`
	AmountPaise  int64   `gorm:"not null"`
	Amount       float64 `gorm:"-"` // derived

	// INITIATED, PROCESSING, SUCCESS, FAILED
	Status string `gorm:"type:varchar(20);not null;default:'INITIATED';index"`

	Reason string `gorm:"type:text;not null"`

	// Gateway refund tracking — null for offline
	GatewayRefundRef *string    `gorm:"type:varchar(200)"`
	GatewayResponse  *string    `gorm:"type:text"` // raw snapshot
	RefundedAt       *time.Time `gorm:"type:timestamp"`

	// Who/what triggered this refund — opaque
	InitiatedBy string  `gorm:"type:varchar(100);not null"`
	SourceRef   *string `gorm:"type:varchar(100)"` // e.g. cancellation ID from caller

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Refund) TableName() string { return "payments.refunds" }
