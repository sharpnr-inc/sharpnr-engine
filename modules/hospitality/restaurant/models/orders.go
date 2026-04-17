package restaurant_models

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantOrders struct {
	ID                        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrganizationID            uuid.UUID  `gorm:"type:uuid;not null;index"`
	RestaurantID              uuid.UUID  `gorm:"type:uuid;not null;index"`
	LegalEntityID             uuid.UUID  `gorm:"type:uuid;not null;index"`
	RestaurantDiningSessionID *uuid.UUID `gorm:"type:uuid;index"`
	RestaurantPickUpOrdersID  *uuid.UUID `gorm:"type:uuid;index"`
	OrderType                 string     `gorm:"default;null"`
	OrderNumber               string     `gorm:"type:varchar(50);not null;index"`
	Status                    string     `gorm:"type:varchar(30);not null;index"` // PENDING, CONFIRMED, PREPARING, READY, SERVED, CANCELLED
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func (RestaurantOrders) TableName() string {
	return "hospitality.restaurant_orders"
}

/* ----------------------------------- <*> ---------------------------------- */

type RestaurantOrderItems struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantOrderID uuid.UUID `gorm:"type:uuid;not null;index"`
	MenuItemID        uuid.UUID `gorm:"type:uuid;not null;index"`
	MenuItemName      string
	Quantity          int     `gorm:"not null"`
	UnitPrice         float64 `gorm:"type:numeric(12,2);not null"`
	Status            string  `gorm:"type:varchar(30);not null;index"` // PENDING, PREPARING, READY, SERVED, CANCELLED
	Notes             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (RestaurantOrderItems) TableName() string {
	return "hospitality.restaurant_order_items"
}

/* ----------------------------------- Item Variants --------------------------------- */

type RestaurantOrderItemsVariants struct {
	ID                    uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantOrderItemID uuid.UUID `gorm:"type:uuid;not null;index"`
	VariantID             uuid.UUID `gorm:"type:uuid;not null"`
	VariantGroupID        uuid.UUID `gorm:"type:uuid;not null"`
	Name                  string    `gorm:"type:varchar(100);not null"`
	PriceAtOrder          float64   `gorm:"type:numeric(12,2);not null"`
	CreatedAt             time.Time
}

func (RestaurantOrderItemsVariants) TableName() string {
	return "hospitality.restaurant_order_items_variants"
}

/* ----------------------------------- Item Add-Ons ---------------------------------- */

type RestaurantOrderItemsAddOns struct {
	ID                    uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantOrderItemID uuid.UUID `gorm:"type:uuid;not null;index"`
	AddOnID               uuid.UUID `gorm:"type:uuid;not null"`
	Name                  string    `gorm:"type:varchar(100);not null"`
	Quantity              int       `gorm:"type:integer;not null;default:1"`
	PriceAtOrder          float64   `gorm:"type:numeric(12,2);not null"`
	CreatedAt             time.Time
}

func (RestaurantOrderItemsAddOns) TableName() string {
	return "hospitality.restaurant_order_items_addons"
}
