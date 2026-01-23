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
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	PropertyID     uuid.UUID `gorm:"type:uuid;index"`

	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255)"`

	// Owner: RESTAURANT, BAR, HOTEL, HOUSEKEEPING
	OwnerModule string    `gorm:"type:varchar(50)"`
	OwnerID     uuid.UUID `gorm:"type:uuid;index"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuGroup) TableName() string { return "hospitality.menu_group" }

// -------------------------------
// 1️⃣ MENU CATEGORY (e.g., Starters, Main Course)
// -------------------------------
type MenuCategory struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuGroupID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255)"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuCategory) TableName() string { return "hospitality.menu_category" }

// -------------------------------
// 1.5️⃣ MENU SUBCATEGORY (e.g., Soup, Salad)
// -------------------------------
type MenuSubCategory struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuCategoryID uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255)"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuSubCategory) TableName() string { return "hospitality.menu_subcategory" }

// -------------------------------
// 2️⃣ MENU ITEM
// -------------------------------
type MenuItem struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuSubCategoryID uuid.UUID `gorm:"type:uuid;index"`
	OrganizationID    uuid.UUID `gorm:"type:uuid;not null;index"`

	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255)"`
	IsVeg       bool   `gorm:"default:true"`

	BasePrice   float64 `gorm:"type:decimal(12,2);not null"`
	IsAvailable bool    `gorm:"default:true"`

	// Tax: FOOD, ALCOHOL, ROOM_SERVICE
	TaxCategoryCode string `gorm:"type:varchar(50)"`
	Unit            string `gorm:"type:varchar(20);default:'plate'"`

	IsActive       bool      `gorm:"default:true"`
	ServicePointID uuid.UUID `gorm:"type:uuid;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuItem) TableName() string { return "hospitality.menu_item" }

// -------------------------------
// 2.5️⃣ MENU VARIANT GROUP (Size, Crust)
// -------------------------------
type MenuVariantGroup struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuItemID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name          string `gorm:"type:varchar(50);not null"`
	IsMultiSelect bool   `gorm:"default:false"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuVariantGroup) TableName() string { return "hospitality.menu_variant_group" }

// -------------------------------
// 3️⃣ MENU VARIANT (Small, Large)
// -------------------------------
type MenuVariant struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuVariantGroupID uuid.UUID `gorm:"type:uuid;not null;index"`
	MenuItemID         uuid.UUID `gorm:"type:uuid;not null;index"`

	Name     string  `gorm:"type:varchar(50);not null"`
	Price    float64 `gorm:"type:decimal(12,2);not null"`
	IsActive bool    `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuVariant) TableName() string { return "hospitality.menu_variant" }

// -------------------------------
// 5️⃣ ADDONS (Extra Cheese, Toppings)
// -------------------------------
type MenuItemAddon struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	MenuItemID uuid.UUID `gorm:"type:uuid;not null;index"`

	Name  string  `gorm:"type:varchar(100);not null"`
	Price float64 `gorm:"type:decimal(12,2);not null"`
	IsVeg bool    `gorm:"default:true"`

	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuItemAddon) TableName() string { return "hospitality.menu_item_addon" }
