package hospitality_order_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type HospitalityOrder struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID
	PropertyID     uuid.UUID

	ServicePointID uuid.UUID
	// table / room / counter / online

	Domain string
	// HOSPITALITY

	Module string
	// HOTEL / RESTAURANT / BAR / CLOUD_KITCHEN

	Status string
	// CREATED, CONFIRMED, IN_PROGRESS, COMPLETED, CANCELLED

	CreatedAt time.Time
	UpdatedAt time.Time
}

type HospitalityOrderItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrderID uuid.UUID

	MenuItemID uuid.UUID
	BasePrice  float64 // snapshot at order time

	Quantity int

	Notes string

	Status string
	// PENDING, PREPARING, SERVED, CANCELLED

	CreatedAt time.Time
}

type HospitalityOrderItemVariant struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrderItemID uuid.UUID `gorm:"type:uuid;index;not null"`

	MenuVariantID uuid.UUID `gorm:"type:uuid;not null"`

	Price float64 // snapshot

	CreatedAt time.Time
}

type HospitalityOrderItemAddon struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrderItemID uuid.UUID `gorm:"type:uuid;index;not null"`

	AddonID uuid.UUID `gorm:"type:uuid;not null"`

	Price    float64 // snapshot per unit
	Quantity int     // VERY IMPORTANT

	CreatedAt time.Time
}
