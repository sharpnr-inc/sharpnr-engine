package legal_entity_contracts

import "time"

type UpdateLegalEntityNamesRequestDto struct {
	LegalName         *string
	DisplayName       *string
	IncorporationDate *time.Time
	FYStartMonth      *int
}

/* ---------------------------------- <***> --------------------------------- */

type UpdateLegalEntityAddressRequestDto struct {
	Type         *string
	AddressLine1 *string
	AddressLine2 *string
	Landmark     *string
	City         *string
	State        *string

	GSTStateCode *string

	Pincode *string
	Country *string
}

/* ---------------------------------- <***> --------------------------------- */
