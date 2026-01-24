package org_legal_entity_contracts

import "github.com/google/uuid"

type NewLegalEntityInput struct {
	OrgID       uuid.UUID
	LegalName   string
	Type        string
	AddressLine string
	City        string
	State       string
	Country     string
	Pincode     string
}
