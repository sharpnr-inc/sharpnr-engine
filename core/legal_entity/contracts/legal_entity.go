package legal_entity_contracts

import (
	legal_entity_constants "sharpnr-engine/core/legal_entity/constants"
	"time"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

type NewLegalEntityRequestDto struct {
	OrgID       uuid.UUID
	LegalName   string
	DisplayName string
	Type        legal_entity_constants.LegalEntityType
	Addresses   []CreateLegalEntityAddressRequest
}

type CreateLegalEntityAddressRequest struct {
	AddressLine1 string
	AddressLine2 string

	AddressType legal_entity_constants.LegalEntityAddressType

	City         string
	State        string
	GSTStateCode string
	Country      string
	Pincode      string
	IsPrimary    bool
}

/* ---------------------------------- <***> --------------------------------- */

type LegalEntityResponseDto struct {
	ID          uuid.UUID
	DisplayName string
	
	EntityType  legal_entity_constants.LegalEntityType
	Country     string
	AddedAt     time.Time
}
