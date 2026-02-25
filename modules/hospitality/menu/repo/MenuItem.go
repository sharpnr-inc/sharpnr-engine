package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

func (m *menuRepo) CreateMenuItem(ctx context.Context, item *hospitalityMenu_models.MenuItem) error {
	return m.db.WithContext(ctx).Create(item).Error
}

// GetMenuItemByID implements MenuRepository.
func (m *menuRepo) GetMenuItemByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuItem, error) {
	var item hospitalityMenu_models.MenuItem

	if err := m.db.WithContext(ctx).First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

// GetItemsBySubCategory implements MenuRepository.
func (m *menuRepo) GetItemsBySubCategoryID(ctx context.Context, subCatID uuid.UUID) ([]hospitalityMenu_models.MenuItem, error) {
	var items []hospitalityMenu_models.MenuItem

	err := m.db.WithContext(ctx).
		Where("menu_sub_category_id = ?", subCatID).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	return items, nil
}
