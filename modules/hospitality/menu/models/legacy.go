package hospitalityMenu_models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 0️⃣ MENU GROUP (Context: Bar, Restaurant, Room Service, etc.)
// -------------------------------
type MenuGroup struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	PropertyID     uuid.UUID

	// e.g., "Restaurant Menu", "Bar Menu", "Room Service Menu"
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`

	OwnerModule string
	// RESTAURANT, BAR, HOTEL, HOUSEKEEPING, CONCIERGE

	OwnerID uuid.UUID
	// restaurant_id, hotel_id, property_id

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 1️⃣ MENU CATEGORY
// -------------------------------
type MenuCategory struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuGroupID uuid.UUID `gorm:"type:uuid;not null"`

	// e.g., "Main Course", "Appetizers", "Desserts"
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 1.5️⃣ MENU SUBCATEGORY
// -------------------------------
type MenuSubCategory struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuCategoryID uuid.UUID `gorm:"type:uuid;not null"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`

	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 2️⃣ MENU ITEM
// -------------------------------
type MenuItem struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuSubCategoryID uuid.UUID `gorm:"type:uuid"`
	OrganizationID    uuid.UUID `gorm:"type:uuid;not null"`

	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`
	IsVeg       bool   `gorm:"default:true"`

	BasePrice   float64 `gorm:"not null"`
	IsAvailable bool    `gorm:"default:true"`

	TaxCategoryCode string
	// FOOD, ALCOHOL, ROOM_SERVICE, LAUNDRY, CONCIERGE

	Unit string `gorm:"default:'plate'"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// This is Menu item have many variant group
	// this is Menu item have many Addon group
}

type MenuVariantGroup struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	MenuItemID uuid.UUID `gorm:"not null"`

	Name          string `gorm:"size:50;not null"` // e.g., "Size", "Crust", "Spice Level"
	IsMultiSelect bool   `gorm:"default:false"`    // user can pick multiple? (e.g., toppings)

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// This tables is collection of variants
}

// -------------------------------
// 3️⃣ MENU VARIANT
// -------------------------------
type MenuVariant struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuVariantGroupID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItemID         uuid.UUID `gorm:"type:uuid;not null"` // FIXED (was MenuItemId without gorm tag)

	Name     string  `gorm:"size:50;not null"`
	Price    float64 `gorm:"not null"`
	IsActive bool    `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 5️⃣ ADDONS
// -------------------------------
type MenuItemAddon struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`

	Name  string  `gorm:"size:100;not null"`
	Price float64 `gorm:"not null"`
	IsVeg bool    `gorm:"default:true"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
