package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/business_ops/hospitality/kitchen/models"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
)

// CreateKitchen inserts a new kitchen record into the database
func (r *kitchenRepository) CreateKitchen(ctx context.Context, kitchen *kitchen_models.Kitchen) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	// Use WithContext to respect request timeouts and cancellations
	if err := db.WithContext(ctx).Create(kitchen).Error; err != nil {
		return err
	}
	return nil
}

func (r *kitchenRepository) UpdateKitchenGeneralDetails(ctx context.Context, kitchenID uuid.UUID, updates map[string]interface{}) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	err := db.WithContext(ctx).Model(&kitchen_models.Kitchen{}).
		Where("id = ?", kitchenID).
		Updates(updates).Error

	return err
}

// GetKitchenById implements KitchenRepositoryInterface.
func (r *kitchenRepository) GetKitchenById(ctx context.Context, kitchenId uuid.UUID) (*kitchen_models.Kitchen, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var kitchens *kitchen_models.Kitchen

	// Using Find instead of First because the return type is a slice
	err := db.WithContext(ctx).Where("id = ?", kitchenId).Find(&kitchens).Error
	if err != nil {
		return nil, err
	}

	return kitchens, nil
}

// GetKitchenByOrgId implements KitchenRepositoryInterface.
func (r *kitchenRepository) GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) ([]*kitchen_models.Kitchen, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var kitchens []*kitchen_models.Kitchen

	err := db.WithContext(ctx).Where("organization_id = ?", orgId).Find(&kitchens).Error
	if err != nil {
		return nil, err
	}

	return kitchens, nil
}

// GetKitchenByPropertyId implements KitchenRepositoryInterface.
func (r *kitchenRepository) GetKitchenByPropertyId(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_models.Kitchen, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	var kitchens []*kitchen_models.Kitchen

	err := db.WithContext(ctx).Where("organization_property_id = ?", propertyID).Find(&kitchens).Error
	if err != nil {
		return nil, err
	}

	return kitchens, nil
}

// SwitchKitchenOnOff implements KitchenRepositoryInterface.
func (r *kitchenRepository) SwitchKitchenOnOff(ctx context.Context, kitchenID uuid.UUID, b bool) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	// Updates the 'is_active' or 'status' column. Adjust column name as per your schema.
	err := db.WithContext(ctx).Model(&kitchen_models.Kitchen{}).
		Where("id = ?", kitchenID).
		Update("is_active", b).Error

	return err
}
