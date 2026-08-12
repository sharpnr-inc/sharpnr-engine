package stockwaveEngineServices

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"
	stockwaveEngineRepo "sharpnr-engine/stockwave_engine/repo"

	"github.com/google/uuid"
)

type inventoryEngine struct {
	repo       stockwaveEngineRepo.InventoryEngineRepositoryInterface
	txManager  dbtx.TxManager
	seroEngine *sero.SEROEngine
}

// GetInventoryAccountID implements [InventoryEngineServicesInterface].
func (s *inventoryEngine) GetInventoryAccountID(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, inventoryAccountID uuid.UUID) (*stockwaveEngineModels.InventoryAccount, error) {
	panic("unimplemented")
}

// GetInventoryByID implements [InventoryEngineServicesInterface].
func (s *inventoryEngine) GetInventoryByID(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, inventoryID uuid.UUID) (*stockwaveEngineModels.Inventory, error) {
	panic("unimplemented")
}

func NewInventoryEngineService(
	repo stockwaveEngineRepo.InventoryEngineRepositoryInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *inventoryEngine {
	return &inventoryEngine{
		repo:       repo,
		txManager:  txManager,
		seroEngine: seroEngine,
	}
}
