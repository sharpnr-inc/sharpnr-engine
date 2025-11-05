package kitchenrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

// CreateKitchen inserts a new kitchen record into the database
func (d *kitchenRepository) CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error {
	// Use WithContext to respect request timeouts and cancellations
	if err := d.db.WithContext(ctx).Create(kitchen).Error; err != nil {
		return err
	}
	return nil
}
