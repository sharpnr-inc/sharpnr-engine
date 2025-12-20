package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// CreateAddon implements MenuRepository.
func (m *menuRepo) CreateAddon(ctx context.Context, addon *hospitalitymodels.Addon) error {
	return m.db.WithContext(ctx).Create(addon).Error
}

// DeleteAddon implements MenuRepository.
func (m *menuRepo) DeleteAddon(id uuid.UUID) error {
	panic("unimplemented")
}

// UpdateAddon implements MenuRepository.
func (m *menuRepo) UpdateAddon(addon *hospitalitymodels.Addon) error {
	panic("unimplemented")
}
