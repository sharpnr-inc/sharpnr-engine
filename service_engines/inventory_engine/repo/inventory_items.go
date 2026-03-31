package inventory_engine_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
)

func (r *inventoryEngineRepo) CreateInventoryItems(ctx context.Context, item *inventory_engine_models.InventoryItem) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(item).Error
}

func (r *inventoryEngineRepo) GetInventoryItemsByID(ctx context.Context, item *inventory_engine_models.InventoryItem) (*inventory_engine_models.InventoryItem, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result inventory_engine_models.InventoryItem
	// Querying using the ID from the passed struct
	err := db.WithContext(ctx).
		Where("id = ?", item.ID).
		First(&result).
		Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *inventoryEngineRepo) GetInventoryItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.InventoryItem, error) {
	var items []inventory_engine_models.InventoryItem

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where("inventory_id = ?", inventoryID).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	return items, nil
}
