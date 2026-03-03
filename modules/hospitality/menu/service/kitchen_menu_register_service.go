package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

func (s *menuService) SetKitchenToMenuItem(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuItemInput) error {
	// Create a new entry in the MenuItemKitchenRegister table
	entry := &hospitalityMenu_models.MenuItemKitchenRegister{
		ID:         uuid.New(),
		MenuItemID: req.MenuItemID,
		KitchenID:  req.KitchenID,
	}

	return s.menuRepo.CreateEntryInKitchenMenuRegister(ctx, entry)
}

func (c *menuService) GetKitchenOfMenuItem(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.GetMenuItemKitchen(ctx, menuItemID)
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}
