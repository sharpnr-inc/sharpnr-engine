package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// CreateCategory implements MenuRepository.
func (m *menuRepo) CreateCategory(ctx context.Context, category *hospitalitymodels.MenuCategory) error {
	return m.db.WithContext(ctx).Create(category).Error
}

// GetCategoriesByGroup implements MenuRepository.
func (m *menuRepo) GetCategoriesByGroup(ctx context.Context, groupID uuid.UUID) ([]hospitalitymodels.MenuCategory, error) {
	var categories []hospitalitymodels.MenuCategory

	err := m.db.WithContext(ctx).
		Where("group_id = ?", groupID).
		Preload("SubCategories").
		Preload("SubCategories.MenuItems").
		Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}
