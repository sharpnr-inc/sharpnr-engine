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

func (s *menuService) AddTaxCategoryInMenuGroup(
	ctx context.Context,
	orgID, emplID, menuGroupOwnerID uuid.UUID,
	req hospitality_menu_contracts.MenuGroupTaxCategoryRequestDto,
) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		menGrpTax := &hospitalityMenu_models.MenuGroupTaxCategory{
			ID:            uuid.New(),
			MenuGroupID:   req.MenuGroupID,
			LegalEntityID: req.LegalEntityID,
			TaxCategoryID: req.TaxCategoryID,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := s.menuRepo.CreateMenuGroupTaxCategory(txCtx, menGrpTax); err != nil {
			return err
		}

		return nil
	})

}

func (s *menuService) GetTaxCategoryOfMenuItem(
	ctx context.Context,
	orgID, emplID, menuItemID uuid.UUID,
) (uuid.UUID, uuid.UUID, error) {

	// 1. Get the item
	item, err := s.menuRepo.GetMenuItemByID(ctx, menuItemID)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	// LEVEL 1: ITEM LEVEL
	if mItem, _ := s.menuRepo.GetMenuItemTaxCategoryByMenuItemID(ctx, item.ID); mItem != nil {
		return mItem.TaxCategoryID, mItem.LegalEntityID, nil
	}

	// LEVEL 2 & 3 & 4 (TRACING THROUGH SUBCATEGORY)
	if item.MenuSubCategoryID != uuid.Nil {
		// LEVEL 2: SUBCATEGORY
		if mSub, _ := s.menuRepo.GetMenuSubCategoryTaxCategoryBySubCategoryID(ctx, item.MenuSubCategoryID); mSub != nil {
			return mSub.TaxCategoryID, mSub.LegalEntityID, nil
		}

		// Trace up to Category
		subCat, _ := s.menuRepo.GetSubCategoryByID(ctx, item.MenuSubCategoryID)
		if subCat != nil && subCat.MenuCategoryID != uuid.Nil {

			// LEVEL 3: CATEGORY
			if mCat, _ := s.menuRepo.GetMenuCategoryTaxCategoryByMenuCategoryID(ctx, subCat.MenuCategoryID); mCat != nil {
				return mCat.TaxCategoryID, mCat.LegalEntityID, nil
			}

			// Trace up to Group from Category
			cat, _ := s.menuRepo.GetCategoryByID(ctx, subCat.MenuCategoryID)
			if cat != nil && cat.MenuGroupID != uuid.Nil {
				// LEVEL 4: MENU GROUP
				if mGroup, err := s.menuRepo.GetMenuGroupTaxCategoryByMenuGroupID(ctx, cat.MenuGroupID); err == nil && mGroup != nil {
					return mGroup.TaxCategoryID, mGroup.LegalEntityID, nil
				}
			}
		}
	}

	// LEVEL 4 (DIRECT FALLBACK): If item has a direct GroupID but no SubCategory
	if item.MenuGroupID != uuid.Nil {
		if mGroup, err := s.menuRepo.GetMenuGroupTaxCategoryByMenuGroupID(ctx, item.MenuGroupID); err == nil && mGroup != nil {
			return mGroup.TaxCategoryID, mGroup.LegalEntityID, nil
		}
	}

	return uuid.Nil, uuid.Nil, nil
}
