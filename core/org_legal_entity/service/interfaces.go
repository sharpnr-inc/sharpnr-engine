package legal_entity_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"

	"github.com/google/uuid"
)

type OrgLegalEntityServiceInterface interface {
	CreateNewLegalEntity(ctx context.Context, input legal_entity_contracts.NewLegalEntityInput) (legalEntityID uuid.UUID, err error)
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*legal_entity_contracts.LegalEntityReader, error)

	LegalEntityTaxCongfigsServiceInterface
}

type LegalEntityTaxCongfigsServiceInterface interface {
}
