package hospitalityMenu_models

import (
	"time"

	"github.com/google/uuid"
)

type MenuGroupTaxCategory struct {
	ID            uuid.UUID
	MenuGroupID   uuid.UUID
	LegalEntityID uuid.UUID
	TaxCategoryID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (MenuGroupTaxCategory) TableName() string {
	return "hospitality.menu_group_tax_category"
}

type MenuCategoryTaxCategory struct {
	ID             uuid.UUID
	MenuCategoryID uuid.UUID
	LegalEntityID  uuid.UUID
	TaxCategoryID  uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (MenuCategoryTaxCategory) TableName() string {
	return "hospitality.menu_category_tax_category"
}

type MenuSubCategoryTaxCategory struct {
	ID              uuid.UUID
	MenuSubCategory uuid.UUID
	LegalEntityID   uuid.UUID
	TaxCategoryID   uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (MenuSubCategoryTaxCategory) TableName() string {
	return "hospitality.menu_sub_category_tax_category"
}

type MenuItemTaxCategory struct {
	ID               uuid.UUID
	MenuItemCategory uuid.UUID
	LegalEntityID    uuid.UUID
	TaxCategoryID    uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (MenuItemTaxCategory) TableName() string {
	return "hospitality.menu_item_tax_category"
}
