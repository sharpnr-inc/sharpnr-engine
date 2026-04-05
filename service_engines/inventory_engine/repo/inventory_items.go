package inventory_engine_repo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	inventory_engine_contracts "sharpnr-engine/service_engines/inventory_engine/contracts"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateInventoryItem - Standard record creation
func (r *inventoryEngineRepo) CreateInventoryItem(ctx context.Context, item *inventory_engine_models.InventoryItem) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(item).Error
}

// GetInventoryItemByID - Fetch by primary key (ID)
func (r *inventoryEngineRepo) GetInventoryItemByID(ctx context.Context, itemID uuid.UUID) (*inventory_engine_models.InventoryItem, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result inventory_engine_models.InventoryItem
	err := db.WithContext(ctx).First(&result, "id = ?", itemID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInventoryItemsByInventoryID - List all items in a specific inventory
func (r *inventoryEngineRepo) GetInventoryItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.InventoryItem, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var items []inventory_engine_models.InventoryItem
	err := db.WithContext(ctx).Where("inventory_id = ?", inventoryID).Find(&items).Error
	return items, err
}

// GetInventoryItemBySKU - Fetch by unique SKU
func (r *inventoryEngineRepo) GetInventoryItemBySKU(ctx context.Context, sku string) (*inventory_engine_models.InventoryItem, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result inventory_engine_models.InventoryItem
	err := db.WithContext(ctx).Where("sku = ?", sku).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLowStockItemsByInventoryID - Filter items below threshold
func (r *inventoryEngineRepo) GetLowStockItemsByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]inventory_engine_models.InventoryItem, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var items []inventory_engine_models.InventoryItem
	// Logic: quantity <= min_stock_level
	err := db.WithContext(ctx).
		Where("inventory_id = ? AND quantity <= min_stock_level", inventoryID).
		Find(&items).
		Error
	return items, err
}

// UpdateInventoryItem - Fixed to support partial/zero-value updates via Pointer Struct
func (r *inventoryEngineRepo) UpdateInventoryItem(
	ctx context.Context,
	itemID uuid.UUID,
	params inventory_engine_contracts.UpdateInventoryItemParams,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Model(&inventory_engine_models.InventoryItem{}).
		Where("id = ?", itemID).
		Updates(params)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("inventory item not found: %s", itemID)
	}
	return nil
}

// AdjustInventoryItemQuantity - Atomic update using SQL expression
func (r *inventoryEngineRepo) AdjustInventoryItemQuantity(ctx context.Context, itemID uuid.UUID, delta float64) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// gorm.Expr ensures the calculation happens in the database: UPDATE ... SET quantity = quantity + delta
	result := db.WithContext(ctx).
		Model(&inventory_engine_models.InventoryItem{}).
		Where("id = ?", itemID).
		Update("quantity", gorm.Expr("quantity + ?", delta))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("inventory item not found: %s", itemID)
	}
	return nil
}

// DeleteInventoryItem - Soft Delete implementation
func (r *inventoryEngineRepo) DeleteInventoryItem(ctx context.Context, itemID uuid.UUID) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).Delete(&inventory_engine_models.InventoryItem{}, "id = ?", itemID)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("inventory item not found: %s", itemID)
	}
	return nil
}
