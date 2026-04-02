package restaurantservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *restaurantService) GetAllInventoryAccounts(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error) {
	var inventoryAcconts []inventory_engine_models.InventoryAccount

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		accounts, err := s.inventoryEngine.GetAllSharedInventoryAccount(txCtx, orgID, emplID)
		if err != nil {
			return err
		}

		inventoryAcconts = accounts

		return nil
	})

	if err != nil {
		return inventoryAcconts, err
	}

	return inventoryAcconts, nil
}
