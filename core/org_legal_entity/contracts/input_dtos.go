package legal_entity_contracts

import (
	"time"

	"github.com/google/uuid"
)

// Request Body for the Creating the Tax Registration For the Legal Entity
// LegalEntityTaxRegistration
type CreateLegalEntityTaxRegistrationRequestDto struct {
	LegalEntityID      uuid.UUID `json:"legal_entity_id"`
	TaxType            string    `json:"tax_type"`
	RegistrationNumber string    `json:"registration_number"`

	Country string `json:"country"`
	State   string `json:"state"`

	EffectiveFrom time.Time  `json:"effective_from"`
	EffectiveTo   *time.Time `json:"effective_to,omitempty"`
}

/* ---------------------------------- <***> --------------------------------- */

type CreateLegalEntityTaxCategory struct {
	LegalEntityID      uuid.UUID `json:"legal_entity_id"`
	CategoryName       string    `json:"category_name"`
	Code               string    `json:"code"`
	Description        string    `json:"description,omitempty"`
	PreferredAccountID uuid.UUID
}

/* ---------------------------------- <***> --------------------------------- */
