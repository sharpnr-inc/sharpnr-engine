package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (m *menuRepo) CreateMenuItem(ctx context.Context, item *hospitalitymodels.MenuItem) error {
	return m.db.WithContext(ctx).Create(item).Error
}

// GetMenuItemByID implements MenuRepository.
func (m *menuRepo) GetMenuItemByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.MenuItem, error) {
	var item hospitalitymodels.MenuItem

	if err := m.db.WithContext(ctx).First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}
