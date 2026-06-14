package accounting_service

import (
	"context"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

	"github.com/google/uuid"
)

func (s *accountingServices) GetAllInventoriesAccounts(ctx context.Context, orgID, emplID, legalEntity uuid.UUID) ([]inventory_engine_models.Inventory, error) {

	inventory, err := s.inventoryEngine.GetInventoriesOfLegalEntity(ctx, orgID, emplID, legalEntity)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}
