package hospitalityMenu_models

import (
	hospitality_utilities "sharpnr-engine/modules/hospitality/utilities_api"
	"time"

	"github.com/google/uuid"
)

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
