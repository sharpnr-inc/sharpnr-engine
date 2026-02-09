package hospitality_menu_contracts

import (
	hospitality_utilities "sharpnr-engine/modules/hospitality/utilities_api"

	"github.com/google/uuid"
)

type CreateMenuInput struct {
	MenuName        string
	MenuDescription *string
	OwnerID         uuid.UUID
	OrganizationID  uuid.UUID
	PropertyID      uuid.UUID
	OwnerModuleType hospitality_utilities.HospitalityModulesType
	Categories      *[]AddMenuCategoryInput
}

type AddMenuCategoryInput struct {
	MenuID              uuid.UUID
	CategoryName        string
	CategoryDescription string
	SubCategories       *[]AddSubCategoryInput
}

type AddSubCategoryInput struct {
	MenuCategoryID    uuid.UUID
	SubCatName        string
	SubCatDescription string
	MenuItem          *[]AddMenuItemInput
}

type AddMenuItemInput struct {
	MenuID            uuid.UUID
	MenuSubCategoryID uuid.UUID
	Name              string
	Description       string
	Unit              string
	BasePrice         float64
	VariantGroup      *[]AddVariantGroupInput
	AddOns            *[]AddAddOnInput
}

type AddVariantGroupInput struct {
	MenuItemID       uuid.UUID
	VariantGroupName string
	Variant          *[]AddVariantInput
}

type AddVariantInput struct {
	VariantGroupID uuid.UUID
	MenuItemID     uuid.UUID
	Name           string
	Price          string
}

type AddAddOnInput struct {
	MenuItemID uuid.UUID
	Name       string
	Price      float64
}
