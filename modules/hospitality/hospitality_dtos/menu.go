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

type Menu_Item_RESPONSE struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
}
