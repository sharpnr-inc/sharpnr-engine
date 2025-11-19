package hospitalitydtos

import "github.com/google/uuid"

type Menu_AddItemMenu_REQUEST struct {
}

type Menu_AddItemMenu_RESPONSE struct {
}

type Menu_AddCategoryInMenu_RES struct {
	CategoryId   uuid.UUID
	CategoryName string
}

type Menu_AddCategoryInMenu_REQ struct {
	CategoryName        string
	CategoryDescription string
	MenuGroupId         uuid.UUID
}

type Menu_AddSubCategoryInCategory_REQ struct {
	CategoryID             uuid.UUID
	SubCategoryName        string
	SubCategoryDescription string
	MenuType               string
}
type Menu_AddSubCategoryInCategory_RES struct {
	SubCategoryName string
	SubCategoryID   uuid.UUID
	CategoryName    string
}

type Menu_GetAllMenuGroup_RESPONSE struct {
	Id          uuid.UUID
	Name        string
	Type        string
	Description string
}

type Menu_CreateCategoryINMenuGroup_REQ struct {
	MenuGroupId         string `json:"menu_group_id"`
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
}

type Menu_CreateSubCategoryINCategory_REQ struct {
	MenuGroupId            string
	CategoryId             string
	SubCategoryName        string
	SubCategoryDescription string
}

type Menu_Category_RESPONSE struct {
	Id            uuid.UUID                   `json:"id"`
	Name          string                      `json:"name"`
	Description   string                      `json:"description"`
	SubCategories []Menu_SubCategory_RESPONSE `json:"sub_categories"`
}

type Menu_SubCategory_RESPONSE struct {
	Id          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	MenuItems   []Menu_Item_RESPONSE `json:"menu_items"`
}

type Menu_CreateSubCategory_REQ struct {
	CategoryID      string `json:"category_id"`
	SubCategoryName string `json:"sub_category_name"`
	SubCategoryDesc string `json:"sub_category_description"`
}

type Menu_CreateMenuItem_REQ struct {
	CategoryID    string  `json:"category_id"`
	SubCategoryID string  `json:"sub_category_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	IsVeg         bool    `json:"is_veg"`
	BasePrice     float64 `json:"base_price"`
	ServedBy      string  `json:"served_by"` // optional, default "KITCHEN"
}

type Menu_VariantGroup_RESPONSE struct {
	Id            uuid.UUID               `json:"id"`
	Name          string                  `json:"name"`
	IsMultiSelect bool                    `json:"is_multi_select"`
	Variants      []Menu_Variant_RESPONSE `json:"variants"`
}

//------------------------------------------------------------------------------//

type Menu_Item_RESPONSE struct {
	Id            uuid.UUID                    `json:"id"`
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	Price         float32                      `json:"price"`
	VariantGroups []Menu_VariantGroup_RESPONSE `json:"variant_groups"`
	AddonGroups   []Menu_AddGroupList          `json:"add_groups"`
}

type Menu_Variant_RESPONSE struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

type Menu_ADD_Variant_REQ struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Menu_CreateVariantGroupAndVariant_REQ struct {
	MenuItemID  uuid.UUID              `json:"menu_item_id"`
	VariantName string                 `json:"variant_name"`
	Variant     []Menu_ADD_Variant_REQ `json:"variant"`
}

type Menu_AddGroupList struct {
	ID     uuid.UUID
	Name   string
	AddOns []Menu_AddOns
}

type Menu_AddOns struct {
	ID    uuid.UUID
	Name  string
	Price float64
}
