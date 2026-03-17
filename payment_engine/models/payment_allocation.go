package payment_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// Decouples receiving money from applying money to a document
// Allows: advance payments, reallocate across invoices, credit notes
// Engine stores TargetType/TargetID as opaque refs — never foreign keys
type PaymentAllocation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TenantID  uuid.UUID `gorm:"type:uuid;not null;index"`
	PaymentID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Opaque — caller defines what these mean
	// e.g. TargetType="INVOICE", TargetID="invoice-uuid"
	//      TargetType="ADVANCE", TargetID="advance-account-uuid"
	//      TargetType="CREDIT_NOTE", TargetID="credit-note-uuid"
	TargetType string `gorm:"type:varchar(30);not null"`
	TargetID   string `gorm:"type:varchar(100);not null"`

	AllocatedPaise  int64   `gorm:"not null"`
	AllocatedAmount float64 `gorm:"-"` // derived

	// ACTIVE, REVERSED
	Status string `gorm:"type:varchar(20);not null;default:'ACTIVE';index"`

	CreatedBy    string `gorm:"type:varchar(100);not null"`
	CreatedAt    time.Time
	ReversedAt   *time.Time `gorm:"type:timestamp"`
	ReversedBy   *string    `gorm:"type:varchar(100)"`
	ReversalNote *string    `gorm:"type:text"`
}

func (PaymentAllocation) TableName() string { return "payments.payment_allocations" }
