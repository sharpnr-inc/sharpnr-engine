package inventory_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type PurchaseOrder struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	InventoryID   uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`
	SupplierID    uuid.UUID `gorm:"type:uuid;not null;index"`

	Reference string // PO-2026-001
	Status    string `gorm:"size:20;default:'DRAFT'"`
	// DRAFT | SENT | PARTIAL_RECEIVED | RECEIVED | CANCELLED

	ExpectedDate *time.Time
	ReceivedAt   *time.Time

	// totals
	TotalAmount float64
	// links back to accounting
	AccountingPurchaseID *uuid.UUID `gorm:"type:uuid"` // set after PO received

	CreatedBy uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ── Purchase Order Line ──
type PurchaseOrderLine struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	PurchaseOrderID uuid.UUID `gorm:"type:uuid;not null;index"`
	InventoryItemID uuid.UUID `gorm:"type:uuid;not null"`

	QuantityOrdered  float64
	QuantityReceived float64 `gorm:"default:0"`
	UnitPrice        float64
	TaxCategoryID    *uuid.UUID `gorm:"type:uuid"` // your existing tax system
	TotalAmount      float64

	CreatedAt time.Time
}
