package menuservice

import (
	"context"
	"errors"
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
) (*uuid.UUID, error) {
	var foundID uuid.UUID
	txCtx := dbtx.InjectTx(ctx, nil) // Assuming tx is managed by manager

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx = dbtx.InjectTx(ctx, tx)

		// Helper to handle the "RecordNotFound" noise
		tryFetch := func(fetchFn func() (uuid.UUID, error)) (bool, error) {
			id, err := fetchFn()
			if err == nil {
				foundID = id
				return true, nil // Found it
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, nil // Keep hunting
			}
			return false, err // Real DB error, stop everything
		}

		// 1. Fetch Item Lineage
		item, err := s.menuRepo.GetMenuItemByID(txCtx, menuItemID)
		if err != nil {
			return err
		}

		// 2. The Chain of Command
		// Level: Item
		if found, err := tryFetch(func() (uuid.UUID, error) {
			m, e := s.menuRepo.GetMenuItemTaxCategoryByMenuItemID(txCtx, item.ID)
			return m.TaxCategoryID, e
		}); found || err != nil {
			return err
		}

		// Level: SubCategory
		if item.MenuSubCategoryID != uuid.Nil {
			if found, err := tryFetch(func() (uuid.UUID, error) {
				m, e := s.menuRepo.GetMenuSubCategoryTaxCategoryBySubCategoryID(txCtx, item.MenuSubCategoryID)
				return m.TaxCategoryID, e
			}); found || err != nil {
				return err
			}
		}

		// Level: Category
		// (Assuming you fetch the CategoryID from the SubCategory here)
		subCat, _ := s.menuRepo.GetSubCategoryByID(txCtx, item.MenuSubCategoryID)
		if subCat.MenuCategoryID != uuid.Nil {
			if found, err := tryFetch(func() (uuid.UUID, error) {
				m, e := s.menuRepo.GetMenuCategoryTaxCategoryByMenuCategoryID(txCtx, subCat.MenuCategoryID)
				return m.TaxCategoryID, e
			}); found || err != nil {
				return err
			}
		}

		// Level: Group (The Final Stand)
		if found, err := tryFetch(func() (uuid.UUID, error) {
			m, e := s.menuRepo.GetMenuGroupTaxCategoryByMenuGroupID(txCtx, item.MenuGroupID)
			return m.TaxCategoryID, e
		}); found || err != nil {
			return err
		}

		return gorm.ErrRecordNotFound
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &foundID, err
}
