package hospitality_order_engine_models

import (
	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities"
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 1️⃣ HOSPITALITY ORDER (The Header)
// -------------------------------
type HospitalityOrder struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyID     uuid.UUID `gorm:"type:uuid;not null;index"`
	// OwnnersID is the ID of the Restaurant/Hotel/Bar that owns this order. This allows us to easily query all orders for a specific owner.
	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	Module hospitality_utilities.HospitalityModulesType `gorm:"type:varchar(50);index"` // HOTEL, RESTAURANT, BAR

	// Status: CREATED, CONFIRMED, IN_PROGRESS, COMPLETED, CANCELLED
	Status string `gorm:"type:varchar(20);index;default:'CREATED'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (HospitalityOrder) TableName() string { return "hospitality.hospitality_order" }

// -------------------------------
// 2️⃣ ORDER ITEM
// -------------------------------
type HospitalityOrderItem struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID uuid.UUID `gorm:"type:uuid;not null;index"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	BasePrice  float64   `gorm:"type:decimal(12,2);not null"` // Captured snapshot

	Quantity int    `gorm:"not null;default:1"`
	Notes    string `gorm:"type:text"`

	// Status: PENDING, PREPARING, SERVED, CANCELLED
	Status string `gorm:"type:varchar(20);index;default:'PENDING'"`

	CreatedAt time.Time
}

func (HospitalityOrderItem) TableName() string { return "hospitality.hospitality_order_item" }

// -------------------------------
// 3️⃣ ITEM VARIANTS (e.g. Medium/Large snapshot)
// -------------------------------
type HospitalityOrderItemVariant struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderItemID uuid.UUID `gorm:"type:uuid;index;not null"`

	MenuVariantID uuid.UUID `gorm:"type:uuid;not null"`
	Price         float64   `gorm:"type:decimal(12,2);not null"` // Captured snapshot

	CreatedAt time.Time
}

func (HospitalityOrderItemVariant) TableName() string {
	return "hospitality.hospitality_order_item_variant"
}

// -------------------------------
// 4️⃣ ITEM ADDONS (e.g. Extra Cheese snapshot)
// -------------------------------
type HospitalityOrderItemAddon struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderItemID uuid.UUID `gorm:"type:uuid;index;not null"`

	AddonID uuid.UUID `gorm:"type:uuid;not null"`

	Price    float64 `gorm:"type:decimal(12,2);not null"` // Captured snapshot
	Quantity int     `gorm:"not null;default:1"`

	CreatedAt time.Time
}

func (HospitalityOrderItemAddon) TableName() string {
	return "hospitality.hospitality_order_item_addon"
}
