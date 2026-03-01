package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
)

// CreateKitchen inserts a new kitchen record into the database
func (d *kitchenRepository) CreateKitchen(ctx context.Context, kitchen *kitchen_models.Kitchen) error {
	// Use WithContext to respect request timeouts and cancellations
	if err := d.db.WithContext(ctx).Create(kitchen).Error; err != nil {
		return err
	}
	return nil
}

// GetKitchenById implements KitchenRepositoryInterface.
func (d *kitchenRepository) GetKitchenById(ctx context.Context, kitchenId uuid.UUID) (*kitchen_models.Kitchen, error) {
	var kitchens *kitchen_models.Kitchen

	// Using Find instead of First because the return type is a slice
	err := d.db.WithContext(ctx).Where("id = ?", kitchenId).Find(&kitchens).Error
	if err != nil {
		return nil, err
	}

	return kitchens, nil
}

// GetKitchenByOrgId implements KitchenRepositoryInterface.
func (d *kitchenRepository) GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) ([]*kitchen_models.Kitchen, error) {
	var kitchens []*kitchen_models.Kitchen

	err := d.db.WithContext(ctx).Where("organization_id = ?", orgId).Find(&kitchens).Error
	if err != nil {
		return nil, err
	}

	return kitchens, nil
}

// GetKitchenByPropertyId implements KitchenRepositoryInterface.
func (d *kitchenRepository) GetKitchenByPropertyId(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_models.Kitchen, error) {
	var kitchens []*kitchen_models.Kitchen

	err := d.db.WithContext(ctx).Where("organization_property_id = ?", propertyID).Find(&kitchens).Error
	if err != nil {
		return nil, err
	}

	return kitchens, nil
}

// SwitchKitchenOnOff implements KitchenRepositoryInterface.
func (d *kitchenRepository) SwitchKitchenOnOff(ctx context.Context, kitchenID uuid.UUID, b bool) error {
	// Updates the 'is_active' or 'status' column. Adjust column name as per your schema.
	err := d.db.WithContext(ctx).Model(&kitchen_models.Kitchen{}).
		Where("id = ?", kitchenID).
		Update("is_active", b).Error

	return err
}
