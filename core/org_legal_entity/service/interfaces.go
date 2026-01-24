package org_legal_entity_service

import (
	"context"
	org_legal_entity_contracts "sharpnr-engine/core/org_legal_entity/contracts"
)

type OrgLegalEntityServiceInterface interface {
	CreateNewLegalEntity(ctx context.Context, input org_legal_entity_contracts.NewLegalEntityInput) error
}
