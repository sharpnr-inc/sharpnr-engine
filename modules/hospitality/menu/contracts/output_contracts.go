package hospitality_menu_contracts

import (
	menu_constants "sharpnr-engine/modules/hospitality/menu/constants"

	"github.com/google/uuid"
)

type MenuOutput struct {
	ID              uuid.UUID        `json:"id"`
	MenuName        string           `json:"menu_name"`
	MenuDescription *string          `json:"menu_description"`
	Categories      []CategoryOutput `json:"categories"`
}

type CategoryOutput struct {
	ID                  uuid.UUID           `json:"id"`
	MenuID              uuid.UUID           `json:"menu_id"`
	CategoryName        string              `json:"category_name"`
	CategoryDescription string              `json:"category_description"`
	SubCategories       []SubCategoryOutput `json:"sub_categories"`
}

type SubCategoryOutput struct {
	ID                uuid.UUID        `json:"id"`
	MenuCategoryID    uuid.UUID        `json:"menu_category_id"`
	SubCatName        string           `json:"sub_cat_name"`
	SubCatDescription string           `json:"sub_cat_description"`
	MenuItems         []MenuItemOutput `json:"menu_items"`
}

type MenuItemOutput struct {
	ID                uuid.UUID                          `json:"id"`
	MenuID            uuid.UUID                          `json:"menu_id"`
	MenuSubCategoryID uuid.UUID                          `json:"menu_sub_category_id"`
	Name              string                             `json:"name"`
	Description       string                             `json:"description"`
	Unit              string                             `json:"unit"`
	BasePrice         float64                            `json:"base_price"`
	DietaryType       menu_constants.MenuItemDietaryType `json:"dietary_type"`
	VariantGroups     []VariantGroupOutput               `json:"variant_groups"`
	AddOns            []AddOnOutput                      `json:"add_ons"`
	ServedBy          uuid.UUID                          `json:"served_by"` // KitchenID from MenuItemKitchenRegister
}

type VariantGroupOutput struct {
	ID               uuid.UUID       `json:"id"`
	MenuItemID       uuid.UUID       `json:"menu_item_id"`
	VariantGroupName string          `json:"variant_group_name"`
	Variants         []VariantOutput `json:"variants"`
}

type VariantOutput struct {
	ID             uuid.UUID `json:"id"`
	VariantGroupID uuid.UUID `json:"variant_group_id"`
	MenuItemID     uuid.UUID `json:"menu_item_id"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
}

type AddOnOutput struct {
	ID         uuid.UUID `json:"id"`
	MenuItemID uuid.UUID `json:"menu_item_id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
}
