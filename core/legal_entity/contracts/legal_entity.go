package legal_entity_contracts

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

type NewLegalEntityRequestDto struct {
	OrgID       uuid.UUID                              `json:"org_id"`
	LegalName   string                                 `json:"legal_name"`
	DisplayName string                                 `json:"display_name"`
	Type        legal_entity_constants.LegalEntityType `json:"type"`
	Addresses   []CreateLegalEntityAddressRequest      `json:"addresses,omitempty"`
}

type CreateLegalEntityAddressRequest struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2,omitempty"`

	AddressType legal_entity_constants.LegalEntityAddressType `json:"address_type"`

	City         string `json:"city"`
	State        string `json:"state"`
	GSTStateCode string `json:"gst_state_code"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode"`
	IsPrimary    bool   `json:"is_primary"`
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityResponseDto struct {
	ID          uuid.UUID `json:"id"`
	DisplayName string    `json:"display_name"`
	LegalName   string
	Status      bool
	EntityType  legal_entity_constants.LegalEntityType `json:"entity_type"`
	
}
