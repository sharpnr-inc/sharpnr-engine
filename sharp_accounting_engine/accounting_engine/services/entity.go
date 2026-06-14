package accounting_service

import (
	"context"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"

	"github.com/google/uuid"
)

func (s *accountingServices) GetLegalEntities(ctx context.Context, orgID, emplID uuid.UUID) ([]legal_entity_contracts.LegalEntityResponseDto, error) {
	return s.legalEntity.GetLegalEntityByOrgID(ctx, orgID, emplID)
}
