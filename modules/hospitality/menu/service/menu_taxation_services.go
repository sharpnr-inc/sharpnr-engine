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

	// 1. Always get the item lineage first
	item, err := s.menuRepo.GetMenuItemByID(ctx, menuItemID)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	// LEVEL 1: ITEM
	mItem, _ := s.menuRepo.GetMenuItemTaxCategoryByMenuItemID(ctx, item.ID)
	if mItem != nil {
		return mItem.TaxCategoryID, mItem.LegalEntityID, nil
	}

	// LEVEL 2: SUBCATEGORY
	if item.MenuSubCategoryID != uuid.Nil {
		mSub, _ := s.menuRepo.GetMenuSubCategoryTaxCategoryBySubCategoryID(ctx, item.MenuSubCategoryID)
		if mSub != nil {
			return mSub.TaxCategoryID, mSub.LegalEntityID, nil
		}

		// LEVEL 3: CATEGORY (Parent of SubCategory)
		subCat, _ := s.menuRepo.GetSubCategoryByID(ctx, item.MenuSubCategoryID)
		if subCat != nil && subCat.MenuCategoryID != uuid.Nil {
			mCat, _ := s.menuRepo.GetMenuCategoryTaxCategoryByMenuCategoryID(ctx, subCat.MenuCategoryID)
			if mCat != nil {
				return mCat.TaxCategoryID, mCat.LegalEntityID, nil
			}
		}
	}

	// LEVEL 4: MENU GROUP (The one you just confirmed works)
	if item.MenuGroupID != uuid.Nil {
		mGroup, err := s.menuRepo.GetMenuGroupTaxCategoryByMenuGroupID(ctx, item.MenuGroupID)
		if err == nil && mGroup != nil {
			return mGroup.TaxCategoryID, mGroup.LegalEntityID, nil
		}
	}

	// If we reached here, absolutely nothing was found
	return uuid.Nil, uuid.Nil, nil
}
