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
