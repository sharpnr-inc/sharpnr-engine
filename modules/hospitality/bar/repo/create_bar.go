package barrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
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
