package inventory_engine_services

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"
	inventory_engine_repo "sharpnr-engine/service_engines/inventory_engine/repo"

	"github.com/google/uuid"
)

type inventoryEngine struct {
	repo       inventory_engine_repo.InventoryEngineRepositoryInterface
	txManager  dbtx.TxManager
	seroEngine *sero.SEROEngine
}

// GetInventoryAccountID implements [InventoryEngineServicesInterface].
func (s *inventoryEngine) GetInventoryAccountID(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, inventoryAccountID uuid.UUID) (*inventory_engine_models.InventoryAccount, error) {
	panic("unimplemented")
}

// GetInventoryByID implements [InventoryEngineServicesInterface].
func (s *inventoryEngine) GetInventoryByID(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, inventoryID uuid.UUID) (*inventory_engine_models.Inventory, error) {
	panic("unimplemented")
}

func NewInventoryEngineService(
	repo inventory_engine_repo.InventoryEngineRepositoryInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *inventoryEngine {
	return &inventoryEngine{
		repo:       repo,
		txManager:  txManager,
		seroEngine: seroEngine,
	}
}
