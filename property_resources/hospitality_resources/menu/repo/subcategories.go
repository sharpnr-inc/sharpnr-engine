package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/property_resources/hospitality_resources/menu/models"

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

// GetSubCategoriesByCategory implements MenuRepository.
func (m *menuRepo) GetSubCategoriesByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]hospitalityMenu_models.MenuSubCategory, error) {
	var subCategories []hospitalityMenu_models.MenuSubCategory

	// Use Find instead of First to populate a slice
	err := m.db.WithContext(ctx).
		Where("menu_category_id = ?", categoryID).
		Find(&subCategories).Error

	if err != nil {
		return nil, err
	}

	return subCategories, nil
}
