package org_legal_entity_contracts

import (
	"time"

	"github.com/google/uuid"
)

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

type LegalEntityReader struct {
	LegalEntityID uuid.UUID
	Name          string
	EntityType    string
	GST           *string
	CIN           *string
	PAN           *string
	Country       string
	AddedAt       time.Time
}
