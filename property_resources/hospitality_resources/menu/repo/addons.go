package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/property_resources/hospitality_resources/menu/models"

	"github.com/google/uuid"
)

// CreateAddon implements MenuRepository.
func (m *menuRepo) CreateAddon(ctx context.Context, addon *hospitalityMenu_models.MenuItemAddon) error {
	return m.db.WithContext(ctx).Create(addon).Error
}

func (m *menuRepo) GetAddOnsByMenuItemID(ctx context.Context, menuItemID uuid.UUID) ([]hospitalityMenu_models.MenuItemAddon, error) {
	var addons []hospitalityMenu_models.MenuItemAddon

	err := m.db.WithContext(ctx).
		Where("menu_item_id = ?", menuItemID).
		Find(&addons).Error

	if err != nil {
		return nil, err
	}

	return addons, nil
}

// DeleteAddon implements MenuRepository.
func (m *menuRepo) DeleteAddon(id uuid.UUID) error {
	panic("unimplemented")
}

// UpdateAddon implements MenuRepository.
func (m *menuRepo) UpdateAddon(addon *hospitalityMenu_models.MenuItemAddon) error {
	panic("unimplemented")
}
