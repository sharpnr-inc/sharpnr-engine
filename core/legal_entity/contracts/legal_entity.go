package legal_entity_contracts

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

type NewLegalEntityRequestDto struct {
	LegalName         string                                 `json:"legal_name"`
	DisplayName       string                                 `json:"display_name"`
	Type              legal_entity_constants.LegalEntityType `json:"type"`
	IncorporationDate string                                 `json:"incorporation_date"`
	FiscalYearMonth   int                                    `json:"fiscal_year_month"`
	// Address is Optional
	Addresses []CreateLegalEntityAddressRequest `json:"addresses,omitempty"`
}

type CreateLegalEntityAddressRequest struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	Landmark     string `json:"landmark"`

	AddressType legal_entity_constants.LegalEntityAddressType `json:"address_type"`

	City string `json:"city"`

	State        string `json:"state"`
	GSTStateCode string `json:"gst_state_code"`

	Country   string `json:"country"`
	Pincode   string `json:"pincode"`
	IsPrimary bool   `json:"is_primary"`
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityResponseDto struct {
	ID                uuid.UUID                              `json:"id"`
	DisplayName       string                                 `json:"display_name"`
	LegalName         string                                 `json:"legal_name"`
	Status            bool                                   `json:"status"`
	EntityType        legal_entity_constants.LegalEntityType `json:"entity_type"`
	IncorporationDate time.Time                              `json:"inc_date"`
}
