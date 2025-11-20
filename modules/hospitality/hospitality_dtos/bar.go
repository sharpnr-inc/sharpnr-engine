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
	BarId   string             `json:"bar_id"`
	BarName string             `json:"bar_name"`
	Floors  []BarFloorResponse `json:"floors"`
	Tables  []BarTablesResponse
}

type CreateTableInBar_RESPONSE struct {
	BarId  string
	Floors []BarFloorResponse `json:"floors"`
	Tables []BarTablesResponse
}

type BAR_UpdateBarDetails_Request struct {
	BarDisplayName string
	BarLegalName   string
	OpeningTime    time.Time
	ClosingTime    time.Time
}

type BAR_UpdateBarDetails_Response struct {
	BarID                          uuid.UUID
	OrgId                          uuid.UUID
	BarDisplayName                 string
	BarLegalName                   string
	OpeningTime                    time.Time
	ClosingTime                    time.Time
	IsKitchen                      bool
	KitchenId                      uuid.UUID
	IsShareLegalDetailsWithMainOrg bool
	DisplayName                    string
	LegalName                      string
	GstNumber                      string
	IsMultifloor                   bool
	AddressLine                    string
	City                           string
	State                          string
	Country                        string
	Pincode                        string
	Floors                         []BarFloorResponse  `json:"floors"`
	Tables                         []BarTablesResponse `json:"tables"`
}

type BAR_Add_Floor_Req struct {
	BarId     uuid.UUID
	FloorName string
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
