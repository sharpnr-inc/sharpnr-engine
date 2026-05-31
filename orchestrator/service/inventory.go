package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	inventory_dtos "sharpnr-engine/sharp_inventory_engine/dtos"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *inventory_dtos.NewInventoryAccountRequestDto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.inventoryEngine.AddNewInventoryAccount(txCtx, orgID, emplID, request); err != nil {
			return err
		}

		return nil

	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) ListAllInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error) {
	var results []inventory_engine_models.InventoryAccount

	// Wrapping in a transaction ensures read consistency across the request
	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		// Inject the current transaction into the context
		txCtx := dbtx.InjectTx(ctx, tx)

		// Delegate retrieval to the inventory engine
		accounts, err := s.inventoryEngine.ListInventoryAccountOfOrg(txCtx, orgID, emplID)
		if err != nil {
			return err
		}

		results = accounts
		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) GetInventoryAccountDetails(ctx context.Context, orgID, emplID, InventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error) {
	var inventoryList []inventory_engine_models.Inventory

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		invt, err := s.inventoryEngine.ListAllInventoryOfInventoryAccount(txCtx, orgID, emplID, InventoryAccountID)
		if err != nil {
			return err
		}

		inventoryList = invt
		return nil
	})

	if err != nil {
		return inventoryList, err
	}

	return inventoryList, nil
}

/* ---------------------------------- <***> --------------------------------- */
