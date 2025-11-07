package hospitalitymodels

import (
	"time"

	"sharpnr-engine/core/models"

	"github.com/google/uuid"
)

// -------------------------------------------
// 1️⃣ HospitalityOrder
// -------------------------------------------

type HospitalityOrder struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	OrderSource string  `gorm:"size:50;not null"`          // HOTEL_ROOM, RESTAURANT, BAR
	Status      string  `gorm:"size:20;default:'PENDING'"` // PENDING, PREPARING, SERVED, COMPLETED, CANCELLED
	TotalAmount float64 `gorm:"not null;default:0"`
	Remarks     string  `gorm:"size:255"`

	// Optional foreign keys depending on source
	BookingID         *uuid.UUID `gorm:"type:uuid" json:"booking_id,omitempty"`          // for HOTEL_ROOM orders
	RestaurantTableID *uuid.UUID `gorm:"type:uuid" json:"restaurant_table_id,omitempty"` // for RESTAURANT orders
	BarTableID        *uuid.UUID `gorm:"type:uuid" json:"bar_table_id,omitempty"`        // for BAR orders
	FloorID           *uuid.UUID `gorm:"type:uuid" json:"floor_id,omitempty"`            // optional floor reference
	RoomID            *uuid.UUID `gorm:"type:uuid" json:"room_id,omitempty"`             // for hotel room service

	OrderItems []HospitalityOrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`

	PlacedAt    time.Time `gorm:"autoCreateTime"`
	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// -------------------------------------------
// 2️⃣ HospitalityOrderItem
// -------------------------------------------

type HospitalityOrderItem struct {
	ID      uuid.UUID        `gorm:"type:uuid;primaryKey"`
	OrderID uuid.UUID        `gorm:"type:uuid;not null"`
	Order   HospitalityOrder `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItem   MenuItem  `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	VariantID *uuid.UUID   `gorm:"type:uuid" json:"variant_id,omitempty"`
	Variant   *MenuVariant `gorm:"foreignKey:VariantID;constraint:OnDelete:SET NULL;"`

	Quantity   int     `gorm:"not null;default:1"`
	UnitPrice  float64 `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`

	Addons []HospitalityOrderAddon `gorm:"foreignKey:OrderItemID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------------------
// 3️⃣ HospitalityOrderAddon
// -------------------------------------------

type HospitalityOrderAddon struct {
	ID          uuid.UUID            `gorm:"type:uuid;primaryKey"`
	OrderItemID uuid.UUID            `gorm:"type:uuid;not null"`
	OrderItem   HospitalityOrderItem `gorm:"foreignKey:OrderItemID;constraint:OnDelete:CASCADE;"`

	AddonID uuid.UUID `gorm:"type:uuid;not null"`
	Addon   Addon     `gorm:"foreignKey:AddonID;constraint:OnDelete:CASCADE;"`

	Price float64 `gorm:"not null"`
	IsVeg bool    `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
