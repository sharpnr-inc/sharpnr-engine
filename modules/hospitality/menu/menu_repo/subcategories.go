package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (m *menuRepo) CreateSubCategory(ctx context.Context, sub *hospitalitymodels.MenuSubCategory) error {
	return m.db.WithContext(ctx).Create(sub).Error
}

func (m *menuRepo) GetSubCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.MenuSubCategory, error) {
	var sub hospitalitymodels.MenuSubCategory
	err := m.db.WithContext(ctx).
		Where("id = ?", id).
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}
