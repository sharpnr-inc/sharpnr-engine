package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

// CreateMenuGroup implements MenuRepository.
func (m *menuRepo) CreateMenuGroup(ctx context.Context, group *hospitalityMenu_models.MenuGroup) error {
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
func (m *menuRepo) GetAllMenuGroups(ctx context.Context, orgID uuid.UUID) ([]hospitalityMenu_models.MenuGroup, error) {
	var groups []hospitalityMenu_models.MenuGroup

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

// Add helper to get MenuGroup for org validation
func (m *menuRepo) GetMenuGroupByID(ctx context.Context, groupID uuid.UUID) (*hospitalityMenu_models.MenuGroup, error) {
	var group hospitalityMenu_models.MenuGroup
	err := m.db.WithContext(ctx).
		Where("id = ?", groupID).
		First(&group).Error

	if err != nil {
		return nil, err
	}
	return &group, nil
}

// UpdateMenuGroup implements MenuRepository.
func (m *menuRepo) UpdateMenuGroup(group *hospitalityMenu_models.MenuGroup) error {
	panic("unimplemented")
}
