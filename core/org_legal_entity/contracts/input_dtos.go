package org_legal_entity_contracts

import "github.com/google/uuid"

type CreateTaxRule struct {
	LegalEntityID uuid.UUID `json:"legal_entity_id"`
}

type AddTaxRegistrationInLegalEntityRequestDto struct {
	LegalEntityID uuid.UUID
	
}
