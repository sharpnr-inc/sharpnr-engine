package supplier_contracts

import "github.com/google/uuid"

type AddLegalEntityLinkInSuppliersRequestDto struct {
	LegalEntityID uuid.UUID `json:"legal_entity_id"`
	SuppliersID   uuid.UUID `json:"suppliers_id"`
}
