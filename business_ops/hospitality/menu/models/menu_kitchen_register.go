package hospitalityMenu_models

import (
	hospitality_utilities "sharpnr-engine/business_ops/hospitality/utilities_api"
	"time"

	"github.com/google/uuid"
)

type MenuGroupKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuGroupID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	KitchenID   uuid.UUID `gorm:"type:uuid;not null"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuGroupKitchenRegister) TableName() string {
	return "hospitality.menu_group_kitchen_register"
}

type MenuCategoryKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuCategoryID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	KitchenID      uuid.UUID `gorm:"type:uuid;not null"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuCategoryKitchenRegister) TableName() string {
	return "hospitality.menu_category_kitchen_register"
}

type MenuSubCategoryKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuSubCategoryID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	KitchenID         uuid.UUID `gorm:"type:uuid;not null"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuSubCategoryKitchenRegister) TableName() string {
	return "hospitality.menu_subcategory_kitchen_register"
}

type MenuItemKitchenRegister struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	MenuItemID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	KitchenID  uuid.UUID `gorm:"type:uuid;not null"`

	OwnersID uuid.UUID `gorm:"type:uuid;not null;index"`

	OwnerType hospitality_utilities.HospitalityModulesType `gorm:"size:50;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (MenuItemKitchenRegister) TableName() string {
	return "hospitality.menu_item_kitchen_register"
}
