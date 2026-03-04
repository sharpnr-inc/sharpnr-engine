package menuservice

import (
	"context"
	"fmt"
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
func (s *menuService) SetKitchenToMenuGroup(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID,
	req *hospitality_menu_contracts.SetKitchenToMenuGroupInput,
) error {

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

func (s *menuService) SetKitchenToMenuCategory(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID,
	req *hospitality_menu_contracts.SetKitchenToMenuCategoryInput) error {

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

func (c *menuService) GetKitchenOfMenuCategory(ctx context.Context, menuCategoryID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.GetMenuItemKitchen(ctx, menuCategoryID)
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}

/* -------------------------------------------------------------------------- */
/*                 Menu sub-category Kitchen Register Services                */
/* -------------------------------------------------------------------------- */

func (s *menuService) SetKitchenToMenuSubCategory(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID,
	req *hospitality_menu_contracts.SetKitchenToMenuSubCategoryInput) error {

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

func (c *menuService) GetKitchenOfMenuSubCategory(ctx context.Context, menuSubCategoryID uuid.UUID) (uuid.UUID, error) {
	kitchen, err := c.menuRepo.GetSubCategoryPreferredKitchen(ctx, menuSubCategoryID)
	if err != nil {
		return uuid.Nil, err
	}

	return kitchen.KitchenID, nil
}

/* -------------------------------------------------------------------------- */
/*                     Menu Item Kitchen Regsiter Services                    */
/* -------------------------------------------------------------------------- */

func (s *menuService) SetKitchenToMenuItem(
	ctx context.Context,
	orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID,
	req *hospitality_menu_contracts.SetKitchenToMenuItemInput,
) error {
	// Create a new entry in the MenuItemKitchenRegister table
	entry := &hospitalityMenu_models.MenuItemKitchenRegister{
		ID:         uuid.New(),
		MenuItemID: req.MenuItemID,
		KitchenID:  req.KitchenID,
	}

	return s.menuRepo.CreateEntryInKitchenMenuRegister(ctx, entry)
}

func (c *menuService) GetKitchenOfMenuItem(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error) {
	// 1. Check MenuItem Priority (Highest Priority)
	itemKitchen, err := c.menuRepo.GetMenuItemKitchen(ctx, menuItemID)
	// If we found a specific kitchen for this item, return it immediately
	if err == nil && itemKitchen.KitchenID != uuid.Nil {
		return itemKitchen.KitchenID, nil
	}

	// 2. Fetch the Parents (Category and Sub-Category IDs)
	// Fixed: Added uuid.Nil and err to the return statement
	menuItemParent, err := c.GetMenuItemParents(ctx, menuItemID)
	if err != nil {
		return uuid.Nil, err
	}

	// 3. Check Category Priority (Second Priority)
	if menuItemParent.MenuCategoryID != uuid.Nil {
		catKitchen, err := c.menuRepo.GetCategoryPreferredKitchen(ctx, menuItemParent.MenuCategoryID)
		if err == nil && catKitchen.KitchenID != uuid.Nil {
			return catKitchen.KitchenID, nil
		}
	}

	// 4. Check Sub-Category Priority (Third Priority)
	if menuItemParent.MenuSubCategoryID != uuid.Nil {
		subCatKitchen, err := c.menuRepo.GetSubCategoryPreferredKitchen(ctx, menuItemParent.MenuSubCategoryID)
		if err == nil && subCatKitchen.KitchenID != uuid.Nil {
			return subCatKitchen.KitchenID, nil
		}
	}

	// If no kitchen is found at any level
	return uuid.Nil, fmt.Errorf("no kitchen found for item %s or its parents", menuItemID)
}
