package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

// CreateMenuGroup implements MenuRepository.
func (m *menuRepo) CreateMenuGroup(ctx context.Context, group *hospitalityMenu_models.MenuGroup) error {
	if err := m.db.WithContext(ctx).Create(group).Error; err != nil {
		return err
	}

	return nil
}

// GetAllMenuGroups implements MenuRepository.
func (m *menuRepo) GetAllMenuGroups(ctx context.Context, propertyID uuid.UUID) ([]*hospitalityMenu_models.MenuGroup, error) {
	var groups []*hospitalityMenu_models.MenuGroup

	err := m.db.WithContext(ctx).
		Where("property_id = ?", propertyID).
		Find(&groups).
		Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}

// Add helper to get MenuGroup for org validation
func (m *menuRepo) GetMenuByID(ctx context.Context, groupID uuid.UUID) (*hospitalityMenu_models.MenuGroup, error) {
	var group hospitalityMenu_models.MenuGroup
	err := m.db.WithContext(ctx).
		Where("id = ?", groupID).
		First(&group).Error

	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (m *menuRepo) GetMenuByOwnersID(ctx context.Context, ownersID uuid.UUID) (*hospitalityMenu_models.MenuGroup, error) {
	var group hospitalityMenu_models.MenuGroup
	err := m.db.WithContext(ctx).
		Where("owner_id = ?", ownersID).
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
