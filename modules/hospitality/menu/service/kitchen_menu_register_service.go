package menuservice

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (c *menuService) SetKitchenToWholeMenu(ctx context.Context, ownersID uuid.UUID, req *hospitality_menu_contracts.SetKitchenToWholeMenuInput) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		menus, err := c.GetMenu(txCtx, req.MenuID)
		if err != nil {
			return err
		}

		// 2. Iterate through the levels to find every MenuItem
		for _, category := range menus.Categories {
			for _, subCategory := range category.SubCategories {
				for _, item := range subCategory.MenuItems {

					// 3. Prepare the database model for this item
					entry := &hospitalityMenu_models.MenuItemKitchenRegister{
						ID:         uuid.New(),
						MenuItemID: item.ID,
						KitchenID:  req.KitchenID,
						OwnersID:   ownersID,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}

					// 4. Register the item
					// Note: If you have many items, consider a BulkCreate method in your repo instead
					err := c.menuRepo.CreateEntryInKitchenMenuRegister(txCtx, entry)
					if err != nil {
						// Decide if you want to continue or fail fast
						return err
					}
				}
			}
		}

		return nil
	})

}

func (c *menuService) GetKitchenOfMenuItem(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.GetMenuItemKitchen(ctx, menuItemID)
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}
