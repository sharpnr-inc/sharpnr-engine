package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

func (m *menuRepo) CreateSubCategory(ctx context.Context, sub *hospitalityMenu_models.MenuSubCategory) error {
	return m.db.WithContext(ctx).Create(sub).Error
}

func (m *menuRepo) GetSubCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuSubCategory, error) {
	var sub hospitalityMenu_models.MenuSubCategory
	err := m.db.WithContext(ctx).
		Where("id = ?", id).
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}
