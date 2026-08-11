package barrepo

import (
	"context"
	bar_models "sharpnr-engine/property_business_ops/hospitality/bar/models"

	"github.com/google/uuid"
)

func (r *barRepository) CreateBar(ctx context.Context, bar *bar_models.Bar) error {
	if err := r.db.WithContext(ctx).Create(bar).Error; err != nil {
		return err
	}
	return nil
}

func (r *barRepository) CreateFloor(ctx context.Context, floor *bar_models.BarFloor) error {
	return r.db.WithContext(ctx).Create(floor).Error
}

func (r *barRepository) UpdateBarDetails(ctx context.Context, barID uuid.UUID, updates map[string]interface{}) error {

	return r.db.WithContext(ctx).
		Model(&bar_models.Bar{}).
		Where("id = ?", barID).
		Updates(updates).
		Error
}
