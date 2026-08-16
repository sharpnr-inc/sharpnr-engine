package barrepo

import (
	"context"
	bar_models "sharpnr-engine/property_business_ops/hospitality/bar/models"
)

func (r *barRepository) CreateTable(ctx context.Context, barTable *bar_models.BarTable) error {
	return r.db.WithContext(ctx).Create(barTable).Error
}
