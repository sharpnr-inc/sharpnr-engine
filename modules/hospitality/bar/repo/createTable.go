package barrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (r *barRepository) CreateTable(ctx context.Context, barTable *hospitalitymodels.BarTable) error {
	return r.db.WithContext(ctx).Create(barTable).Error
}
