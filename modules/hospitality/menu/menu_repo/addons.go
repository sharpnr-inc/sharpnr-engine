package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

// CreateAddon implements MenuRepository.
func (m *menuRepo) CreateAddon(ctx context.Context, addon *hospitalityMenu_models.Addon) error {
	return m.db.WithContext(ctx).Create(addon).Error
}

// DeleteAddon implements MenuRepository.
func (m *menuRepo) DeleteAddon(id uuid.UUID) error {
	panic("unimplemented")
}

// UpdateAddon implements MenuRepository.
func (m *menuRepo) UpdateAddon(addon *hospitalityMenu_models.Addon) error {
	panic("unimplemented")
}
