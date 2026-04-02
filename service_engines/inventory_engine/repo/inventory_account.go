package inventory_engine_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
)

func (r *inventoryEngineRepo) CreateInventoryAccount(ctx context.Context, inventoryAccount *inventory_engine_models.InventoryAccount) error {
	// 1. Extract transaction from context (if any), otherwise use repository's default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Perform the Create operation within the provided context
	// This ensures that timeouts or cancellations are respected.
	return db.WithContext(ctx).Create(inventoryAccount).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) GetInventoryAccountsByOrgID(ctx context.Context, orgID uuid.UUID) ([]inventory_engine_models.InventoryAccount, error) {
	var accounts []inventory_engine_models.InventoryAccount

	// Extract transaction from context if it exists; otherwise, use the default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Query for all accounts belonging to the specific Organization ID
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&accounts).
		Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
func (r *inventoryEngineRepo) GetSharedInventoryAccountsByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
) ([]inventory_engine_models.InventoryAccount, error) {

	var accounts []inventory_engine_models.InventoryAccount

	// Extract transaction from context if it exists; otherwise use default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where("organization_id = ? AND is_shared = ?", orgID, true).
		Find(&accounts).
		Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
