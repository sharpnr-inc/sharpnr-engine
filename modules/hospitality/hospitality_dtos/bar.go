package hospitalitydtos

import (
	"time"

	"github.com/google/uuid"
)

type BarTablesResponse struct {
	TableName   string    `json:"table_name"`
	TableId     uuid.UUID `json:"table_id"`
	TableStatus string    `json:"table_status"`
}

type BarFloorResponse struct {
	FloorId   uuid.UUID           `json:"floor_id"`
	FloorName string              `json:"floor_name"`
	Tables    []BarTablesResponse `json:"tables"`
}

type GetBarDetails_RESPONSE struct {
	BarId        string              `json:"bar_id"`
	BarName      string              `json:"bar_name"`
	IsMultiFloor bool                `json:"is_multiFloor"`
	Floors       []BarFloorResponse  `json:"floors,omitempty"`
	Tables       []BarTablesResponse `json:"tables,omitempty"`
}

type CreateTableInBar_RESPONSE struct {
	BarId  string
	Floors []BarFloorResponse `json:"floors"`
	Tables []BarTablesResponse
}

type BAR_UpdateBarDetails_Response struct {
	BarID                          uuid.UUID `json:"bar_id"`
	OrgId                          uuid.UUID `json:"org_id"`
	BarDisplayName                 string    `json:"bar_display_name"`
	BarLegalName                   string    `json:"bar_legal_name"`
	OpeningTime                    time.Time `json:"opening_time"`
	ClosingTime                    time.Time `json:"closing_time"`
	IsKitchen                      bool      `json:"is_kitchen"`
	KitchenId                      uuid.UUID `json:"kitchen_id"`
	IsShareLegalDetailsWithMainOrg bool      `json:"is_share_legal_details_with_main_org"`
	DisplayName                    string    `json:"display_name"`
	LegalName                      string    `json:"legal_name"`
	CinNumber                      string    `json:"cin_number"`
	FssaiNumber                    string    `json:"fssai_number"`
	GstNumber                      string    `json:"gst_number"`
	IsMultifloor                   bool      `json:"is_multifloor"`
	AddressLine                    string    `json:"address_line"`
	City                           string    `json:"city"`
	State                          string    `json:"state"`
	Country                        string    `json:"country"`
	Pincode                        string    `json:"pincode"`
}

type BAR_Add_Floor_Req struct {
	BarId     uuid.UUID
	FloorName string
}

type Bar_CreateTablesInBar_Request struct {
	BarId     uuid.UUID   `json:"bar_id"`
	FloorId   uuid.UUID   `json:"floor_id"` // optional now
	TableList []Bar_Table `json:"table_list"`
}

type Bar_Table struct {
	TableName     string `json:"table_name"`
	TableCapacity int    `json:"table_capacity"`
}

type BAR_BarUpdateReq struct {
	// Basic Details
	ID                *uuid.UUID
	Name              *string `json:"name"`
	DisplayName       *string `json:"display_name"`
	LegalName         *string `json:"legal_name"`
	RegisteredAddress *string `json:"registered_address"`

	// Organization (usually not updated, but kept for flexibility)
	OrganizationID *uuid.UUID `json:"organization_id"`

	// Operating Hours
	OpeningTime *time.Time `json:"opening_time"`
	ClosingTime *time.Time `json:"closing_time"`

	// Kitchen Info
	IsKitchen *bool      `json:"is_kitchen"`
	KitchenID *uuid.UUID `json:"kitchen_id"`

	// Legal Details
	IsShareLegalDetailsWithMainOrg *bool   `json:"is_share_legal_details_with_main_org"`
	GSTNumber                      *string `json:"gst_number"`
	CINNumber                      *string `json:"cin_number"`
	FSSAINumber                    *string `json:"fssai_number"`

	// Location Details
	IsMultifloor *bool   `json:"is_multifloor"`
	AddressLine  *string `json:"address_line"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
	Pincode      *string `json:"pincode"`
}
