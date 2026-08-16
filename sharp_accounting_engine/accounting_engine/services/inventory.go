package accounting_service

import (
	"context"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAllInventoriesAccounts(ctx context.Context, orgID, emplID, legalEntity uuid.UUID) ([]stockwaveEngineModels.Inventory, error) {

	inventory, err := s.inventoryEngine.GetInventoriesOfLegalEntity(ctx, orgID, emplID, legalEntity)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}
