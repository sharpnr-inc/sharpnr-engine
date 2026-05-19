package inventory_engine_repo

import (
	"context"
	dbtx "sharpnr-engine/pkg/dbTx"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"
	"time"

	"github.com/google/uuid"
)

// CreateInventoryItemHistory - Records a specific stock movement (Audit Log)
func (r *inventoryEngineRepo) CreateInventoryItemHistory(ctx context.Context, history *inventory_engine_models.InventoryItemHistory) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(history).Error
}

// GetInventoryItemHistoryByItemID - Full ledger for a specific item, sorted by newest first
func (r *inventoryEngineRepo) GetInventoryItemHistoryByItemID(ctx context.Context, itemID uuid.UUID) ([]inventory_engine_models.InventoryItemHistory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var history []inventory_engine_models.InventoryItemHistory
	err := db.WithContext(ctx).
		Where("inventory_item_id = ?", itemID).
		Order("created_at DESC").
		Find(&history).
		Error

	return history, err
}

// GetInventoryItemHistoryByItemIDAndType - Filters history by movement type (e.g., WASTAGE or PURCHASE_RECEIVED)
func (r *inventoryEngineRepo) GetInventoryItemHistoryByItemIDAndType(ctx context.Context, itemID uuid.UUID, changeType string) ([]inventory_engine_models.InventoryItemHistory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var history []inventory_engine_models.InventoryItemHistory
	err := db.WithContext(ctx).
		Where("inventory_item_id = ? AND change_type = ?", itemID, changeType).
		Order("created_at DESC").
		Find(&history).
		Error

	return history, err
}

// GetInventoryItemHistoryByReferenceID - Traces all stock movements linked to a specific source (like a Purchase Order ID)
func (r *inventoryEngineRepo) GetInventoryItemHistoryByReferenceID(ctx context.Context, referenceID uuid.UUID) ([]inventory_engine_models.InventoryItemHistory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var history []inventory_engine_models.InventoryItemHistory
	err := db.WithContext(ctx).
		Where("reference_id = ?", referenceID).
		Order("created_at DESC").
		Find(&history).
		Error

	return history, err
}

// GetInventoryItemHistoryByDateRange - Used for generating periodic stock reports or COGS analysis
func (r *inventoryEngineRepo) GetInventoryItemHistoryByDateRange(ctx context.Context, itemID uuid.UUID, from, to time.Time) ([]inventory_engine_models.InventoryItemHistory, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var history []inventory_engine_models.InventoryItemHistory
	err := db.WithContext(ctx).
		Where("inventory_item_id = ? AND created_at BETWEEN ? AND ?", itemID, from, to).
		Order("created_at DESC").
		Find(&history).
		Error

	return history, err
}
