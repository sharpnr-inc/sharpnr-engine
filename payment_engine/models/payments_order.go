package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// The intent — created before money moves
// One order can have multiple payment attempts (retries, partial pays)
type PaymentOrder struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID uuid.UUID `gorm:"type:uuid;not null;index"`

	OrderNumber string `gorm:"type:varchar(50);not null;uniqueIndex:idx_order_number_tenant"`

	// Payer info — snapshot, no FK outside engine
	PayerType string  `gorm:"type:varchar(30);not null"` // CUSTOMER, GUEST, WALK_IN, VENDOR
	PayerRef  *string `gorm:"type:varchar(100)"`         // opaque — caller's customer ID
	PayerName string  `gorm:"type:varchar(200);not null"`

	CurrencyCode string  `gorm:"type:char(3);not null;default:'INR'"`
	AmountPaise  int64   `gorm:"not null"`  // always smallest unit
	PaidPaise    int64   `gorm:"default:0"` // cached — updated on each confirmed payment
	Amount       float64 `gorm:"-"`         // derived: AmountPaise / 100
	PaidAmount   float64 `gorm:"-"`         // derived: PaidPaise / 100

	// PENDING, PARTIALLY_PAID, PAID, EXPIRED, CANCELLED
	Status string `gorm:"type:varchar(20);not null;default:'PENDING';index"`

	// Opaque reference — engine doesn't know what these mean
	// Caller sets: Source="INVOICE", SourceRefID="invoice-uuid-here"
	Source      string  `gorm:"type:varchar(50)"`
	SourceRefID *string `gorm:"type:varchar(100)"`

	// Optional JSON blob — engine stores and echoes back in webhooks, never reads it
	// Use for any caller metadata: { "invoice_number": "INV-001", "property": "Grand Hyatt" }
	SourceMeta *string `gorm:"type:text"`

	ExpiresAt *time.Time `gorm:"type:timestamp"`

	CreatedBy string `gorm:"type:varchar(100);not null"` // opaque — caller's user/employee ref
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PaymentOrder) TableName() string { return "payments.payment_orders" }
