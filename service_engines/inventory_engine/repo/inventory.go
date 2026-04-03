package inventory_engine_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
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

func (r *inventoryEngineRepo) GetInventoryByID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.Inventory, error) {
	var inventoryList []inventory_engine_models.Inventory

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Using Find with an ID filter populates the slice with the matching record
	err := db.WithContext(ctx).
		Where("id = ?", inventoryID).
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
