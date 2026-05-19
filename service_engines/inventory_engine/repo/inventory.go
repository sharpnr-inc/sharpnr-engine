package inventory_engine_repo

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	inventory_engine_contracts "sharpnr-engine/service_engines/inventory_engine/contracts"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *inventoryEngineRepo) CreateInventory(ctx context.Context, inventory *inventory_engine_models.Inventory) error {
	// Extract transaction from context if it exists; otherwise, use the default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Persist the inventory record using GORM
	// .WithContext(ctx) ensures the operation respects timeouts and cancellations
	return db.WithContext(ctx).Create(inventory).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) GetInventoryByID(
	ctx context.Context,
	inventoryID uuid.UUID,
) (*inventory_engine_models.Inventory, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var inventory inventory_engine_models.Inventory

	// .First() automatically adds "LIMIT 1" and "WHERE deleted_at IS NULL"
	err := db.WithContext(ctx).
		Where("id = ?", inventoryID).
		First(&inventory).
		Error

	if err != nil {
		// You can check if the error is specifically 'Record Not Found'
		// to return a cleaner error message to your service layer.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("inventory with ID %s not found", inventoryID)
		}
		return nil, err
	}

	return &inventory, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) GetInventoryByBusinessUnitID(ctx context.Context, businessUnitID uuid.UUID, inventoryAccountID *uuid.UUID) ([]inventory_engine_models.Inventory, error) {
	var inventoryList []inventory_engine_models.Inventory

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Initialize the query with the mandatory business_unit_id filter
	query := db.WithContext(ctx).Where("business_unit_id = ?", businessUnitID)

	// 2. Check if inventoryAccountID is provided and apply filter
	if inventoryAccountID != nil {
		query = query.Where("inventory_account_id = ?", *inventoryAccountID)
	}

	// Execute the query
	err := query.Find(&inventoryList).Error

	if err != nil {
		return nil, err
	}

	return inventoryList, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) GetInventoryByLegalEntityID(ctx context.Context, legalEntityID uuid.UUID) ([]inventory_engine_models.Inventory, error) {
	var inventoryList []inventory_engine_models.Inventory

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Find(&inventoryList).
		Error

	if err != nil {
		return nil, err
	}

	return inventoryList, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) GetInventoryByInventoryAccountID(ctx context.Context, inventoryAccountID uuid.UUID) ([]inventory_engine_models.Inventory, error) {
	var inventoryList []inventory_engine_models.Inventory

	// 1. Extract transaction from context if it exists; otherwise, use the default DB
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 2. Query for all inventory records linked to the specific Inventory Account
	// This allows you to fetch all stock associated with a specific Bar or Kitchen account
	err := db.WithContext(ctx).
		Where("inventory_account_id = ?", inventoryAccountID).
		Find(&inventoryList).
		Error

	if err != nil {
		return nil, err
	}

	return inventoryList, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) UpdateInventory(
	ctx context.Context,
	inventoryID uuid.UUID,
	params inventory_engine_contracts.UpdateInventoryParams,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// GORM will only update fields in 'params' that are not nil.
	result := db.WithContext(ctx).
		Model(&inventory_engine_models.Inventory{}).
		Where("id = ?", inventoryID).
		Updates(params)

	if result.Error != nil {
		return result.Error
	}

	// Ensure the record actually existed
	if result.RowsAffected == 0 {
		return errors.New("inventory not found")
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *inventoryEngineRepo) DeleteInventory(
	ctx context.Context,
	inventoryID uuid.UUID,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Because the Inventory struct has a gorm.DeletedAt field,
	// this automatically executes a SOFT DELETE.
	result := db.WithContext(ctx).
		Delete(&inventory_engine_models.Inventory{}, inventoryID)

	if result.Error != nil {
		return result.Error
	}

	// Ensure we didn't "delete" a non-existent record
	if result.RowsAffected == 0 {
		return errors.New("inventory not found")
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
