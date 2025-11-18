package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// CreateMenuGroup implements MenuRepository.
func (m *menuRepo) CreateMenuGroup(ctx context.Context, group *hospitalitymodels.MenuGroup) error {
	// Ensure ID is set
	if group.ID == uuid.Nil {
		group.ID = uuid.New()
	}

	// Use context with DB
	if err := m.db.WithContext(ctx).Create(group).Error; err != nil {
		return err
	}

	return nil
}

// GetAllMenuGroups implements MenuRepository.
func (m *menuRepo) GetAllMenuGroups(ctx context.Context, orgID uuid.UUID) ([]hospitalitymodels.MenuGroup, error) {
	var groups []hospitalitymodels.MenuGroup

	err := m.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Preload("Categories").
		Find(&groups).
		Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}

// GetMenuGroupByID implements MenuRepository.
func (m *menuRepo) GetMenuGroupByID(id uuid.UUID) (*hospitalitymodels.MenuGroup, error) {
	panic("unimplemented")
}

// UpdateMenuGroup implements MenuRepository.
func (m *menuRepo) UpdateMenuGroup(group *hospitalitymodels.MenuGroup) error {
	panic("unimplemented")
}
