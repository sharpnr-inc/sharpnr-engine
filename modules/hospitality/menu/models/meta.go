package hospitalityMenu_models

import (
	"time"

	"github.com/google/uuid"
)

type MenuItemHSNCode struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// index ensures fast lookup when the Order Engine checks the menu item
	MenuItemID uuid.UUID `gorm:"type:uuid;index;unique;not null"`

	HSNCode string `gorm:"size:20;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (MenuItemHSNCode) TableName() string {
	return "hospitality.menu_item_hsn_code"
}
