package kitchenrepo

import (
	"context"
	"errors"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateKitchen inserts a new kitchen record into the database
func (d *kitchenRepository) CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error {
	// Use WithContext to respect request timeouts and cancellations
	if err := d.db.WithContext(ctx).Create(kitchen).Error; err != nil {
		return err
	}
	return nil
}

func (d *kitchenRepository) GetKitchenByOrgId(
	ctx context.Context,
	orgId uuid.UUID,
) (*hospitalitymodels.Kitchen, error) {

	var kitchen hospitalitymodels.Kitchen

	err := d.db.WithContext(ctx).
		Where("organization_id = ?", orgId).
		Preload("Organization").
		Preload("MenuItems").
		First(&kitchen).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("kitchen not found for org: %s", orgId)
		}
		return nil, err
	}

	return &kitchen, nil
}
