package hospitalitymodels

import (
	"sharpnr-engine/core/models"
	"time"

	"github.com/google/uuid"
)

type MenuGroup struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	Name        string `gorm:"size:100;not null"` // e.g., "Breakfast Menu", "Bar Menu"
	Description string `gorm:"size:255"`

	Categories []MenuCategory `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 1️⃣ MENU CATEGORY
// -------------------------------

type MenuCategory struct {
	ID             uuid.UUID  `gorm:"primaryKey"`
	RestaurantID   string     `gorm:"type:uuid;not null"`
	Restaurant     Restaurant `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE;"`
	OrganizationId uuid.UUID
	Organization   models.Organization
	Name           string     `gorm:"size:100;not null"`
	Description    string     `gorm:"size:255"`
	MenuItems      []MenuItem `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 2️⃣ MENU ITEM (DISH)
// -------------------------------

type MenuItem struct {
	ID             string       `gorm:"type:uuid;primaryKey"`
	CategoryID     string       `gorm:"type:uuid;not null"`
	Category       MenuCategory `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;"`
	OrganizationId uuid.UUID
	Organization   models.Organization

	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:255"`
	IsVeg       bool    `gorm:"default:true"` // true = Veg, false = Non-Veg
	BasePrice   float64 `gorm:"not null"`
	IsAvailable bool    `gorm:"default:true"`

	Variants    []MenuVariant `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`
	AddonGroups []AddonGroup  `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// 👇 NEW: Kitchen Reference
	KitchenID *uuid.UUID `gorm:"type:uuid"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;"`
}

// -------------------------------
// 3️⃣ VARIANTS (Half / Full / Large / Small)
// -------------------------------

type MenuVariant struct {
	ID         string   `gorm:"type:uuid;primaryKey"`
	MenuItemID string   `gorm:"type:uuid;not null"`
	MenuItem   MenuItem `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	Name  string  `gorm:"size:50;not null"` // e.g., "Half", "Full", "Large"
	Price float64 `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 4️⃣ ADDON GROUPS (e.g., Toppings, Drinks)
// -------------------------------

type AddonGroup struct {
	ID         string   `gorm:"type:uuid;primaryKey"`
	MenuItemID string   `gorm:"type:uuid;not null"`
	MenuItem   MenuItem `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	Name          string `gorm:"size:50;not null"` // e.g., "Toppings", "Drinks"
	IsMultiSelect bool   `gorm:"default:false"`    // Can choose multiple addons?

	Addons []Addon `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 5️⃣ ADDONS (e.g., "Extra Cheese", "Coke")
// -------------------------------

type Addon struct {
	ID      string     `gorm:"type:uuid;primaryKey"`
	GroupID string     `gorm:"type:uuid;not null"`
	Group   AddonGroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`

	Name  string  `gorm:"size:100;not null"`
	Price float64 `gorm:"not null"`
	IsVeg bool    `gorm:"default:true"` // in case addon is food item

	CreatedAt time.Time
	UpdatedAt time.Time
}
