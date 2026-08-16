package stockwaveEngineDtos

import "github.com/google/uuid"

type NewInventoryRequestDto struct {
	BusinessUnitID *uuid.UUID `json:"business_unit_id,omitempty"`
	LegalEntityID  *uuid.UUID `json:"legal_entity_id,omitempty"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Status      string `json:"status"`
}
