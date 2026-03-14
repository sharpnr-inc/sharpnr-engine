package hospitality_menu_contracts

import "github.com/google/uuid"

type MenuGroupTaxCategoryRequestDto struct {
	MenuGroupID   uuid.UUID `json:"menuGroupId"`
	TaxCategoryID uuid.UUID `json:"taxCategoryId"`
	LegalEntityID uuid.UUID `json:"legalEntityId"`
}

type MenuCategoryTaxCategoryRequestDto struct {
	MenuCategoryID uuid.UUID `json:"menuCategoryId"`
	TaxCategoryID  uuid.UUID `json:"taxCategoryId"`
	LegalEntityID  uuid.UUID `json:"legalEntityId"`
}

type MenuSubCategoryGroupTaxCategoryRequestDto struct {
	MenuSubCategoryID uuid.UUID `json:"menuSubCategoryId"`
	TaxCategoryID     uuid.UUID `json:"taxCategoryId"`
	LegalEntityID     uuid.UUID `json:"legalEntityId"`
}

type MenuItemTaxCategoryRequestDto struct {
	MenuItemID    uuid.UUID `json:"menuItemId"`
	TaxCategoryID uuid.UUID `json:"taxCategoryId"`
	LegalEntityID uuid.UUID `json:"legalEntityId"`
}