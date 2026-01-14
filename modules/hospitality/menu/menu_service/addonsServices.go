package menuservice

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

func (m *menuService) CreateAddForMenuItem(ctx context.Context, orgid uuid.UUID, req hospitalitydtos.Menu_CreateAddonsInMenu_REQUEST) error {

	menuItem, err := m.menuRepo.GetMenuItemByID(ctx, req.MenuItemId)
	if err != nil {
		return err
	}

	// Validate org
	if orgid != menuItem.OrganizationID {
		return errors.New("organization mismatch")
	}

	addons := &hospitalityMenu_models.Addon{
		ID:         uuid.New(),
		MenuItemID: menuItem.ID,
		MenuItem:   *menuItem,
		Name:       req.Name,
		Price:      float64(req.Price),
		IsVeg:      true,
		IsActive:   true,
	}

	if err := m.menuRepo.CreateAddon(ctx, addons); err != nil {
		return err
	}

	return nil
}

// UpdateAddon implements MenuService.
func (m *menuService) UpdateAddonOfMenuItem(ctx context.Context, addon *hospitalityMenu_models.Addon) error {
	panic("unimplemented")
}

// DeleteAddon implements MenuService.
func (m *menuService) DeleteAddonOfMenuItem(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
