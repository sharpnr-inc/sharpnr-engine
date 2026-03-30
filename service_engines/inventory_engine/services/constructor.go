package inventory_engine_services

import (
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sero"
	inventory_engine_repo "sharpnr-engine/service_engines/inventory_engine/repo"
)

type inventoryEngine struct {
	repo       inventory_engine_repo.InventoryEngineRepositoryInterface
	txManager  dbtx.TxManager
	seroEngine *sero.SEROEngine
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
