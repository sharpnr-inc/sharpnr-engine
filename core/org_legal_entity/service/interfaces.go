package org_legal_entity_service

import (
	"context"
	org_legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"

	"github.com/google/uuid"
)

type OrgLegalEntityServiceInterface interface {
	CreateNewLegalEntity(ctx context.Context, input org_legal_entity_contracts.NewLegalEntityInput) (legalEntityID uuid.UUID, err error)
	GetLegalEntityByOrgID(ctx context.Context, orgID uuid.UUID) ([]*org_legal_entity_contracts.LegalEntityReader, error)
}
