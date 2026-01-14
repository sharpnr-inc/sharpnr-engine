package barrepo

import (
	"context"
	bar_models "sharpnr-engine/modules/hospitality/bar/models"
)

func (r *barRepository) CreateTable(ctx context.Context, barTable *bar_models.BarTable) error {
	return r.db.WithContext(ctx).Create(barTable).Error
}
