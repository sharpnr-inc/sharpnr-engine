package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// inventoryOrchestrator implements InventoryOrchestratorServiceInterface.
type inventoryOrchestrator struct {
	txManager       dbtx.TxManager
	inventoryEngine stockwaveEngineServices.InventoryEngineServicesInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryOrchestrator) AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID, request *stockwaveEngineDtos.NewInventoryAccountRequestDto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := s.inventoryEngine.AddNewInventoryAccount(txCtx, orgID, emplID, request); err != nil {
			return err
		}

		return nil

	})
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryOrchestrator) ListAllInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) ([]stockwaveEngineModels.InventoryAccount, error) {
	var results []stockwaveEngineModels.InventoryAccount

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

func (s *inventoryOrchestrator) GetInventoryAccountDetails(ctx context.Context, orgID, emplID, InventoryAccountID uuid.UUID) ([]stockwaveEngineModels.Inventory, error) {
	var inventoryList []stockwaveEngineModels.Inventory

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
