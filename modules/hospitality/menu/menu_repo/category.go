package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

// CreateCategory implements MenuRepository.
func (m *menuRepo) CreateCategory(ctx context.Context, category *hospitalitymodels.MenuCategory) error {
	return m.db.WithContext(ctx).Create(category).Error
}
