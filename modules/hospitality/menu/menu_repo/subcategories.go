package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (m *menuRepo) CreateSubCategory(ctx context.Context, sub *hospitalitymodels.MenuSubCategory) error {
	return m.db.WithContext(ctx).Create(sub).Error
}
