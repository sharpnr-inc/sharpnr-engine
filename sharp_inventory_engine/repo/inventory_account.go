package inventory_engine_repo

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/pkg/dbTx"
	inventory_engine_contracts "sharpnr-engine/sharp_inventory_engine/contracts"
	inventory_engine_models "sharpnr-engine/sharp_inventory_engine/models"

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

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) GetInventoryAccountByID(
	ctx context.Context,
	inventoryAccountID uuid.UUID,
) (*inventory_engine_models.InventoryAccount, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var inventoryAccount inventory_engine_models.InventoryAccount

	// Using First to retrieve the record by primary key
	// .WithContext(ctx) ensures the query respects context cancellation/timeouts
	err := db.WithContext(ctx).
		Where("id = ?", inventoryAccountID).
		First(&inventoryAccount).
		Error

	if err != nil {
		return nil, err
	}

	return &inventoryAccount, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) UpdateInventoryAccount(
	ctx context.Context,
	id uuid.UUID,
	params inventory_engine_contracts.UpdateInventoryAccountParams,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Model(&inventory_engine_models.InventoryAccount{}).
		Where("id = ?", id).
		Updates(params)

	if result.Error != nil {
		return result.Error
	}

	// This is the "Check": If no rows were changed, the ID was likely invalid
	if result.RowsAffected == 0 {
		return errors.New("inventory account not found")
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) DeleteInventoryAccount(
	ctx context.Context,
	inventoryAccountID uuid.UUID,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// GORM will automatically perform a Soft Delete because of the DeletedAt field
	result := db.WithContext(ctx).
		Delete(&inventory_engine_models.InventoryAccount{}, inventoryAccountID)

	if result.Error != nil {
		return result.Error
	}

	// Check if the record actually existed to be deleted
	if result.RowsAffected == 0 {
		return errors.New("inventory account not found")
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
