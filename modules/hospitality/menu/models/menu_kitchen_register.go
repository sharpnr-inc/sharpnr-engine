package hospitalityMenu_models

import (
	hospitality_utilities "sharpnr-engine/modules/hospitality/utilities_api"
	"time"

	"github.com/google/uuid"
)

type MenuGroupKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	KitchenID   uuid.UUID `gorm:"type:uuid;not null;index"`
	MenuGroupID uuid.UUID `gorm:"type:uuid;not null;index"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuGroupKitchenRegister) TableName() string {
	return "hospitality.menu_group_kitchen_register"
}

type MenuCategoryKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuCategoryID uuid.UUID `gorm:"type:uuid;not null;index"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null;index"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuCategoryKitchenRegister) TableName() string {
	return "hospitality.menu_category_kitchen_register"
}

type MenuSubCategoryKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuSubCategoryID uuid.UUID `gorm:"type:uuid;not null;index"`
	KitchenID         uuid.UUID `gorm:"type:uuid;not null;index"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuSubCategoryKitchenRegister) TableName() string {
	return "hospitality.menu_subcategory_kitchen_register"
}

type MenuItemKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// index ensures fast lookup when the Order Engine checks the menu item
	MenuItemID uuid.UUID `gorm:"type:uuid;index;unique;not null"`

	KitchenID uuid.UUID `gorm:"type:uuid;not null"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null"`

	// This tells us if it's a BAR, KITCHEN, or BAKERY
	OwnerType hospitality_utilities.HospitalityModulesType `gorm:"size:50"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuItemKitchenRegister) TableName() string {
	return "hospitality.menu_item_kitchen_register"
}
