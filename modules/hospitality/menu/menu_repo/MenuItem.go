package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (m *menuRepo) CreateMenuItem(ctx context.Context, item *hospitalitymodels.MenuItem) error {
	return m.db.WithContext(ctx).Create(item).Error
}
