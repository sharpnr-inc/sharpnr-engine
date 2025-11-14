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
}
