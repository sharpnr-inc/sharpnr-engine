package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/property_business_ops/hospitality/menu/models"

	"github.com/google/uuid"
)

// CreateCategory implements MenuRepository.
func (m *menuRepo) CreateCategory(ctx context.Context, category *hospitalityMenu_models.MenuCategory) error {
	return m.db.WithContext(ctx).Create(category).Error
}

// GetCategoriesByGroup implements MenuRepository.
func (m *menuRepo) GetAllCategoriesByMenuID(ctx context.Context, menuID uuid.UUID) ([]hospitalityMenu_models.MenuCategory, error) {
	var categories []hospitalityMenu_models.MenuCategory

	err := m.db.WithContext(ctx).
		Where("menu_group_id = ?", menuID).
		Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByID implements MenuRepository.
func (m *menuRepo) GetCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuCategory, error) {
	var category hospitalityMenu_models.MenuCategory

	err := m.db.WithContext(ctx).
		Where("id = ?", id).
		First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
