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

/* -------------------------------------------------------------------------- */
/*                         Menu Group Register Service                        */
/* -------------------------------------------------------------------------- */
func (s *menuService) SetKitchenToMenuGroup(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuGroupInput) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		entry := &hospitalityMenu_models.MenuGroupKitchenRegister{
			ID:          uuid.New(),
			MenuGroupID: req.MenuID,
			KitchenID:   req.KitchenID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		return s.menuRepo.CreateMenuGroupKitchenRegister(txCtx, entry)

	})
}

func (c *menuService) GetKitchenfOfMenuGroup(ctx context.Context, menuGroupID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.GetMenuGroupPreferredKitchen(ctx, menuGroupID)
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}

/* -------------------------------------------------------------------------- */
/*                   Menu Category Kitchen Register Services                  */
/* -------------------------------------------------------------------------- */

func (s *menuService) SetKitchenToMenuCategory(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuCategoryInput) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		entry := &hospitalityMenu_models.MenuCategoryKitchenRegister{
			ID:             uuid.New(),
			MenuCategoryID: req.MenuCategoryID,
			KitchenID:      req.KitchenID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		return s.menuRepo.CreateMenuCategoryKitchenRegister(txCtx, entry)

	})
}

func (c *menuService) GetKitchenOfMenuCategory(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.GetMenuItemKitchen(ctx, menuItemID)
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}

/* -------------------------------------------------------------------------- */
/*                 Menu sub-category Kitchen Register Services                */
/* -------------------------------------------------------------------------- */

func (s *menuService) SetKitchenToMenuSubCategory(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuSubCategoryInput) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		entry := &hospitalityMenu_models.MenuSubCategoryKitchenRegister{
			ID:                uuid.New(),
			MenuSubCategoryID: req.MenuSubCategoryID,
			KitchenID:         req.KitchenID,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		return s.menuRepo.CreateMenuSubCategoryKitchenRegister(txCtx, entry)

	})
}

func (c *menuService) GetKitchenOfMenuSubCategory(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}

/* -------------------------------------------------------------------------- */
/*                     Menu Item Kitchen Regsiter Services                    */
/* -------------------------------------------------------------------------- */

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
