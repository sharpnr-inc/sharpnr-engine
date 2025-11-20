package barrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *barRepository) CreateBar(ctx context.Context, bar *hospitalitymodels.Bar) error {
	if err := r.db.WithContext(ctx).Create(bar).Error; err != nil {
		return err
	}
	return nil
}

func (r *barRepository) CreateFloor(ctx context.Context, floor *hospitalitymodels.BarFloor) error {
	return r.db.WithContext(ctx).Create(floor).Error
}

func (r *barRepository) UpdateBarDetails(ctx context.Context, barID uuid.UUID, updates map[string]interface{}) error {

	return r.db.WithContext(ctx).
		Model(&hospitalitymodels.Bar{}).
		Where("id = ?", barID).
		Updates(updates).
		Error
}
