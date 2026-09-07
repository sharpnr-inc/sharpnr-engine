package orchestrator_service

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

	"github.com/google/uuid"
)

// inventoryOrchestrator implements InventoryOrchestratorServiceInterface.
type inventoryOrchestrator struct {
	txManager       dbtx.TxManager
	inventoryEngine stockwaveEngineServices.StockwaveEngineServicesInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryOrchestrator) AddNewInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) error {

	panic("")
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryOrchestrator) ListAllInventoryAccount(ctx context.Context, orgID, emplID uuid.UUID) error {

	panic("")
}

/* ---------------------------------- <***> --------------------------------- */

func (s *inventoryOrchestrator) GetInventoryAccountDetails(ctx context.Context, orgID, emplID, InventoryAccountID uuid.UUID) error {
	panic("")
}

/* ---------------------------------- <***> --------------------------------- */
