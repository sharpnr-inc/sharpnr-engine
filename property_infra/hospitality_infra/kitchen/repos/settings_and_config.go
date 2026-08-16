package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/property_infra/hospitality_infra/kitchen/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                      Kitchen Settings Creates Menthods                     */
/* -------------------------------------------------------------------------- */

func (r *kitchenRepository) CreateKitchenAutoAcceptOrders(ctx context.Context, config *kitchen_models.KitchenAutoAcceptOrders) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(config).Error
}

// CreateKitchenPrinterConfig persists a new printer configuration.
func (r *kitchenRepository) CreateKitchenPrinterConfig(ctx context.Context, config *kitchen_models.KitchenPrinterConfig) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(config).Error
}

// CreateKitchenThrottlingConfig persists the kitchen's capacity/throttling settings.
func (r *kitchenRepository) CreateKitchenThrottlingConfig(ctx context.Context, config *kitchen_models.KitchenThrottlingConfig) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(config).Error
}

// CreateKitchenOperatingHours persists the opening and closing times for a kitchen.
func (r *kitchenRepository) CreateKitchenOperatingHours(ctx context.Context, config *kitchen_models.KitchenOperatingHours) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(config).Error
}

/* -------------------------------------------------------------------------- */
/*                        Kitchen Settings Gets Methods                       */
/* -------------------------------------------------------------------------- */

// GetKitchenAutoAcceptOrdersByKitchenID fetches auto-accept settings by KitchenID.
func (r *kitchenRepository) GetKitchenAutoAcceptOrdersByKitchenID(ctx context.Context, kitchenID uuid.UUID) (*kitchen_models.KitchenAutoAcceptOrders, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var settings kitchen_models.KitchenAutoAcceptOrders
	err := db.WithContext(ctx).Where("kitchen_id = ?", kitchenID).First(&settings).Error
	return &settings, err
}

// GetKitchenPrinterConfigByKitchenID fetches printer configuration by KitchenID.
func (r *kitchenRepository) GetKitchenPrinterConfigByKitchenID(ctx context.Context, kitchenID uuid.UUID) (*kitchen_models.KitchenPrinterConfig, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var config kitchen_models.KitchenPrinterConfig
	err := db.WithContext(ctx).Where("kitchen_id = ?", kitchenID).First(&config).Error
	return &config, err
}

// GetKitchenThrottlingConfigByKitchenID fetches throttling configuration by KitchenID.
func (r *kitchenRepository) GetKitchenThrottlingConfigByKitchenID(ctx context.Context, kitchenID uuid.UUID) (*kitchen_models.KitchenThrottlingConfig, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var config kitchen_models.KitchenThrottlingConfig
	err := db.WithContext(ctx).Where("kitchen_id = ?", kitchenID).First(&config).Error
	return &config, err
}

// ListKitchenOperatingHoursByKitchenID fetches all operating hour records for a specific kitchen.
func (r *kitchenRepository) ListKitchenOperatingHoursByKitchenID(ctx context.Context, kitchenID uuid.UUID) ([]kitchen_models.KitchenOperatingHours, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var hours []kitchen_models.KitchenOperatingHours
	err := db.WithContext(ctx).Where("kitchen_id = ?", kitchenID).Order("day_of_week ASC").Find(&hours).Error
	return hours, err
}

/* -------------------------------------------------------------------------- */
/*                      Kitchens Settings Update Methods                      */
/* -------------------------------------------------------------------------- */

// UpdateKitchenAutoAcceptOrders updates fields dynamically via map.
func (r *kitchenRepository) UpdateKitchenAutoAcceptOrders(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).
		Model(&kitchen_models.KitchenAutoAcceptOrders{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// UpdateKitchenPrinterConfig updates printer fields dynamically.
func (r *kitchenRepository) UpdateKitchenPrinterConfig(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).
		Model(&kitchen_models.KitchenPrinterConfig{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// UpdateKitchenThrottlingConfig updates capacity fields dynamically.
func (r *kitchenRepository) UpdateKitchenThrottlingConfig(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).
		Model(&kitchen_models.KitchenThrottlingConfig{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// UpdateKitchenOperatingHours updates time/status fields dynamically.
func (r *kitchenRepository) UpdateKitchenOperatingHours(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).
		Model(&kitchen_models.KitchenOperatingHours{}).
		Where("id = ?", id).
		Updates(updates).Error
}
