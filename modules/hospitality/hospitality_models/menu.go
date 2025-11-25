package hospitalitymodels

import (
	"sharpnr-engine/core/models"
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 0️⃣ MENU GROUP (Context: Bar, Restaurant, Room Service, etc.)
// -------------------------------
type MenuGroup struct {
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID           `gorm:"type:uuid;not null"`
	Organization   models.Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;"`

	// e.g., "Restaurant Menu", "Bar Menu", "Room Service Menu"
	Name        string `gorm:"size:100;not null;uniqueIndex:idx_org_group"`
	Description string `gorm:"size:255"`

	// Optional — menu context type
	Type string `gorm:"size:50"` // e.g., "BAR", "HOTEL", "RESTAURANT"

	// Optional — if menu belongs to a specific outlet or branch
	// OutletID *uuid.UUID `gorm:"type:uuid"`
	// Outlet   *Outlet     `gorm:"foreignKey:OutletID;constraint:OnDelete:SET NULL;"`

	Categories []MenuCategory `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 1️⃣ MENU CATEGORY
// -------------------------------
type MenuCategory struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	GroupID        uuid.UUID `gorm:"type:uuid;not null"`
	Group          MenuGroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null"`
	Organization   models.Organization

	// e.g., "Main Course", "Appetizers", "Desserts"
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"size:255"`

	SubCategories []MenuSubCategory `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 1.5️⃣ MENU SUBCATEGORY
// -------------------------------
type MenuSubCategory struct {
	ID             uuid.UUID    `gorm:"type:uuid;primaryKey"`
	CategoryID     uuid.UUID    `gorm:"type:uuid;not null"`
	Category       MenuCategory `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE;"`
	OrganizationID uuid.UUID    `gorm:"type:uuid;not null"`
	Organization   models.Organization
	Name           string     `gorm:"size:100;not null"`
	Description    string     `gorm:"size:255"`
	MenuItems      []MenuItem `gorm:"foreignKey:SubCategoryID;constraint:OnDelete:CASCADE;"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// -------------------------------
// 2️⃣ MENU ITEM
// -------------------------------
type MenuItem struct {
	ID             uuid.UUID        `gorm:"type:uuid;primaryKey"`
	SubCategoryID  *uuid.UUID       `gorm:"type:uuid"`
	SubCategory    *MenuSubCategory `gorm:"foreignKey:SubCategoryID;constraint:OnDelete:SET NULL;"`
	OrganizationID uuid.UUID        `gorm:"type:uuid;not null"`
	Organization   models.Organization

	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:255"`
	IsVeg       bool    `gorm:"default:true"`
	BasePrice   float64 `gorm:"not null"`
	IsAvailable bool    `gorm:"default:true"`

	VariantGroups []MenuVariantGroup `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`
	Addons        []Addon            `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`

	// Optional Kitchen link
	KitchenID *uuid.UUID `gorm:"type:uuid"`
	Kitchen   *Kitchen   `gorm:"foreignKey:KitchenID;constraint:OnDelete:SET NULL;"`

	ServedBy string `gorm:"size:50;default:'KITCHEN'" json:"served_by"`
	Unit     string `gorm:"default:'plate'"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MenuVariantGroup struct {
	ID            uuid.UUID     `gorm:"primaryKey"`
	MenuItemID    uuid.UUID     `gorm:"not null"`
	MenuItem      MenuItem      `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`
	Name          string        `gorm:"size:50;not null"` // e.g., "Size", "Crust", "Spice Level"
	IsMultiSelect bool          `gorm:"default:false"`    // user can pick multiple? (e.g., toppings)
	Variants      []MenuVariant `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
	IsActive      bool          `gorm:"default:true"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// -------------------------------
// 3️⃣ MENU VARIANT
// -------------------------------
type MenuVariant struct {
	ID         uuid.UUID        `gorm:"type:uuid;primaryKey"`
	GroupID    uuid.UUID        `gorm:"type:uuid;not null"`
	Group      MenuVariantGroup `gorm:"constraint:OnDelete:CASCADE;"`
	MenuItemID uuid.UUID        `gorm:"type:uuid;not null"` // FIXED (was MenuItemId without gorm tag)
	Name       string           `gorm:"size:50;not null"`
	Price      float64          `gorm:"not null"`
	IsActive   bool             `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// -------------------------------
// 5️⃣ ADDONS
// -------------------------------
type Addon struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuItemID uuid.UUID `gorm:"type:uuid;not null"`
	MenuItem   MenuItem  `gorm:"foreignKey:MenuItemID;constraint:OnDelete:CASCADE;"`
	Name       string    `gorm:"size:100;not null"`
	Price      float64   `gorm:"not null"`
	IsVeg      bool      `gorm:"default:true"`
	IsActive   bool      `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
