package stockwaveEngineRepo

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	stockwaveEngineConstants "sharpnr-engine/stockwave_engine/constants"
	stockwaveEngineModels "sharpnr-engine/stockwave_engine/models"

	"time"

	"github.com/google/uuid"
)

// CreateInventoryItemPrice - Standard record creation for a new price point
func (r *inventoryEngineRepo) CreateInventoryItemPrice(ctx context.Context, price *stockwaveEngineModels.InventoryItemPrice) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(price).Error
}

// GetInventoryItemPriceByID - Fetch a specific price record by its unique ID
func (r *inventoryEngineRepo) GetInventoryItemPriceByID(ctx context.Context, priceID uuid.UUID) (*stockwaveEngineModels.InventoryItemPrice, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result stockwaveEngineModels.InventoryItemPrice
	err := db.WithContext(ctx).First(&result, "id = ?", priceID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInventoryItemPricesByItemID - Retrieves the full audit trail/history of prices for an item
func (r *inventoryEngineRepo) GetInventoryItemPricesByItemID(ctx context.Context, itemID uuid.UUID) ([]stockwaveEngineModels.InventoryItemPrice, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var prices []stockwaveEngineModels.InventoryItemPrice
	err := db.WithContext(ctx).
		Where("inventory_item_id = ?", itemID).
		Order("effective_from DESC").
		Find(&prices).
		Error
	return prices, err
}

// GetActiveInventoryItemPrice - Finds the currently applicable price where effective_to is NULL
func (r *inventoryEngineRepo) GetActiveInventoryItemPrice(
	ctx context.Context,
	itemID uuid.UUID,
	priceType stockwaveEngineConstants.InventoryItemPriceType,
) (*stockwaveEngineModels.InventoryItemPrice, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result stockwaveEngineModels.InventoryItemPrice
	err := db.WithContext(ctx).
		Where("inventory_item_id = ? AND price_type = ? AND effective_to IS NULL", itemID, priceType).
		First(&result).
		Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetActiveInventoryItemPriceBySupplier - Finds the current supplier-specific price
func (r *inventoryEngineRepo) GetActiveInventoryItemPriceBySupplier(
	ctx context.Context,
	itemID uuid.UUID,
	supplierID uuid.UUID,
	priceType stockwaveEngineConstants.InventoryItemPriceType,
) (*stockwaveEngineModels.InventoryItemPrice, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var result stockwaveEngineModels.InventoryItemPrice
	err := db.WithContext(ctx).
		Where("inventory_item_id = ? AND supplier_id = ? AND price_type = ? AND effective_to IS NULL", itemID, supplierID, priceType).
		First(&result).
		Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeactivateInventoryItemPrice - Sets an end date to a price record (Never Hard Deletes)
func (r *inventoryEngineRepo) DeactivateInventoryItemPrice(ctx context.Context, priceID uuid.UUID, effectiveTo time.Time) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	result := db.WithContext(ctx).
		Model(&stockwaveEngineModels.InventoryItemPrice{}).
		Where("id = ?", priceID).
		Update("effective_to", effectiveTo)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("price record not found: %s", priceID)
	}

	return nil
}
