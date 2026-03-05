package hospitality_menu_contracts

import (
	menu_constants "sharpnr-engine/modules/hospitality/menu/constants"
	hospitality_utilities "sharpnr-engine/modules/hospitality/utilities_api"

	"github.com/google/uuid"
)

type CreateMenuInput struct {
	MenuName        string                                       `json:"menu_name"`
	MenuDescription *string                                      `json:"menu_description,omitempty"`
	OwnerID         uuid.UUID                                    `json:"owner_id"`
	OrganizationID  *uuid.UUID                                   `json:"organization_id omitempty"`
	PropertyID      *uuid.UUID                                   `json:"property_id,omitempty"`
	OwnerModuleType hospitality_utilities.HospitalityModulesType `json:"owner_module_type"`
	Categories      *[]AddMenuCategoryInput                      `json:"categories,omitempty"`
}

type AddMenuCategoryInput struct {
	MenuID              uuid.UUID              `json:"menu_id"`
	CategoryName        string                 `json:"category_name"`
	CategoryDescription string                 `json:"category_description"`
	SubCategories       *[]AddSubCategoryInput `json:"sub_categories,omitempty"`
}

type AddSubCategoryInput struct {
	MenuCategoryID    uuid.UUID           `json:"menu_category_id"`
	SubCatName        string              `json:"sub_cat_name"`
	SubCatDescription string              `json:"sub_cat_description"`
	MenuItem          *[]AddMenuItemInput `json:"menu_items,omitempty"`
}

type AddMenuItemInput struct {
	MenuID            uuid.UUID                          `json:"menu_id"`
	MenuSubCategoryID uuid.UUID                          `json:"menu_sub_category_id"`
	Name              string                             `json:"name"`
	Description       string                             `json:"description"`
	Unit              string                             `json:"unit"`
	DietaryType       menu_constants.MenuItemDietaryType `json:"dietary_type"`
	BasePrice         float64                            `json:"base_price"`
	VariantGroup      *[]AddVariantGroupInput            `json:"variant_groups,omitempty"`
	AddOns            *[]AddAddOnInput                   `json:"add_ons,omitempty"`
}

type AddVariantGroupInput struct {
	MenuItemID       uuid.UUID          `json:"menu_item_id"`
	VariantGroupName string             `json:"variant_group_name"`
	Variant          *[]AddVariantInput `json:"variants,omitempty"`
}

type AddVariantInput struct {
	VariantGroupID uuid.UUID `json:"variant_group_id"`
	MenuItemID     uuid.UUID `json:"menu_item_id"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
}

type AddAddOnInput struct {
	MenuItemID uuid.UUID `json:"menu_item_id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
}

/* -------------------------- Configuration Related ------------------------- */

type SetKitchenToMenuGroupInput struct {
	MenuID    uuid.UUID `json:"menu_group_id"`
	KitchenID uuid.UUID `json:"kitchen_id"`
}

type SetKitchenToMenuCategoryInput struct {
	MenuCategoryID uuid.UUID `json:"menu_category_id"`
	KitchenID      uuid.UUID `json:"kitchen_id"`
}
type SetKitchenToMenuSubCategoryInput struct {
	MenuSubCategoryID uuid.UUID `json:"menu_sub_category_id"`
	KitchenID         uuid.UUID `json:"kitchen_id"`
}

type SetKitchenToMenuItemInput struct {
	MenuItemID uuid.UUID `json:"menu_item_id"`
	KitchenID  uuid.UUID `json:"kitchen_id"`
}
