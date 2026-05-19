package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"
	dbtx "sharpnr-engine/pkg/dbTx"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *menuService) GetMenuItemParents(
	ctx context.Context,
	menuItemID uuid.UUID,
) (hospitality_menu_contracts.MenuItemParentsDetails, error) {

	var res hospitality_menu_contracts.MenuItemParentsDetails

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		menuItem, err := s.menuRepo.GetMenuItemByID(txCtx, menuItemID)
		if err != nil {
			return err
		}

		subCategory, err := s.menuRepo.GetSubCategoryByID(txCtx, menuItem.MenuSubCategoryID)
		if err != nil {
			return err
		}

		category, err := s.menuRepo.GetCategoryByID(txCtx, subCategory.MenuCategoryID)
		if err != nil {
			return err
		}

		res = hospitality_menu_contracts.MenuItemParentsDetails{
			MenuCategoryID:    category.ID,
			MenuSubCategoryID: subCategory.ID,
			MenuGroupID:       category.MenuGroupID,
		}

		return nil
	})

	if err != nil {
		return res, err
	}

	return res, nil
}
