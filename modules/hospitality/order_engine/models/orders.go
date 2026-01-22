package order_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
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

type OrderItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrderID uuid.UUID

	MenuItemID uuid.UUID
	VariantID  *uuid.UUID

	Quantity int

	Notes string

	Status string
	// PENDING, PREPARING, SERVED, CANCELLED

	CreatedAt time.Time
}
