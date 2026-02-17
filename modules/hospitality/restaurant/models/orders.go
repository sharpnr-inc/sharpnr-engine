package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantOrders struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	RestaurantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	LegalEntityID  uuid.UUID `gorm:"type:uuid;not null;index"`

	RestaurantDiningSessionID *uuid.UUID `gorm:"type:uuid;index"`
	RestaurantPickUpOrdersID  *uuid.UUID `gorm:"type:uuid;index"`

	OrderNumber string `gorm:"type:varchar(50);not null;index"`

	Status string `gorm:"type:varchar(30);not null;index"`
	// PENDING, CONFIRMED, PREPARING, READY, SERVED, CANCELLED

	TotalAmount float64 `gorm:"type:numeric(12,2);default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantOrders) TableName() string {
	return "hospitality.restaurant_orders"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantOrderItems struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	RestaurantOrderID uuid.UUID `gorm:"type:uuid;not null;index"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null;index"`

	Quantity  float64 `gorm:"type:numeric(10,3);not null"`
	UnitPrice float64 `gorm:"type:numeric(12,2);not null"`
	Total     float64 `gorm:"type:numeric(12,2);not null"`

	Status string `gorm:"type:varchar(30);not null;index"`
	// PENDING, PREPARING, READY, SERVED, CANCELLED

	Notes string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RestaurantOrderItems) TableName() string {
	return "hospitality.restaurant_order_items"
}
