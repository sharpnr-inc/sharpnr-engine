package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/property_resources/hospitality_resources/menu/contracts"
	menu_helpers "sharpnr-engine/property_resources/hospitality_resources/menu/helpers"
	hospitalityMenu_models "sharpnr-engine/property_resources/hospitality_resources/menu/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *menuService) CreateMenu(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	req *hospitality_menu_contracts.CreateMenuInput,
) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		menuGrp := &hospitalityMenu_models.MenuGroup{
			ID:          uuid.New(),
			Name:        req.MenuName,
			Description: *req.MenuDescription,
			OwnerModule: req.OwnerModuleType,
			OwnerID:     req.OwnerID,
			IsActive:    false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := s.menuRepo.CreateMenuGroup(txCtx, menuGrp); err != nil {
			return err
		}

		menuGrpRegister := &hospitality_menu_contracts.SetKitchenToMenuGroupInput{
			MenuID:    menuGrp.ID,
			KitchenID: req.PreferredKitchen,
		}

		if err := s.SetKitchenToMenuGroup(txCtx, orgID, emplID, req.OwnerID, menuGrpRegister); err != nil {
			return nil
		}

		// 🔹 Choose categories source
		var categories []hospitality_menu_contracts.AddMenuCategoryInput

		if req.Categories == nil || len(*req.Categories) == 0 {
			categories = menu_helpers.DefaultMenuCategories
		} else {
			categories = *req.Categories
		}

		// 🔹 Create categories
		for i := range categories {

			cat := &categories[i]
			cat.MenuID = menuGrp.ID

			if _, err := s.AddCategoryInMenu(txCtx, cat); err != nil {
				return err
			}
		}

		res = menuGrp.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

func (s *menuService) AddCategoryInMenu(ctx context.Context, req *hospitality_menu_contracts.AddMenuCategoryInput) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		cat := &hospitalityMenu_models.MenuCategory{
			ID:          uuid.New(),
			MenuGroupID: req.MenuID,
			Name:        req.CategoryName,
			Description: req.CategoryDescription,
			IsActive:    false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := s.menuRepo.CreateCategory(txCtx, cat); err != nil {
			return err
		}

		if req.SubCategories != nil {
			for i := range *req.SubCategories {
				sub := &(*req.SubCategories)[i]
				sub.MenuCategoryID = cat.ID

				if _, err := s.AddSubCategoryInCategory(txCtx, sub); err != nil {
					return err
				}
			}
		}

		res = cat.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

func (s *menuService) AddSubCategoryInCategory(ctx context.Context, req *hospitality_menu_contracts.AddSubCategoryInput) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		subCategory := &hospitalityMenu_models.MenuSubCategory{
			ID:             uuid.New(),
			MenuCategoryID: req.MenuCategoryID,
			Name:           req.SubCatName,
			Description:    req.SubCatDescription,
			IsActive:       false,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := s.menuRepo.CreateSubCategory(txCtx, subCategory); err != nil {
			return err
		}

		if req.MenuItem != nil {
			for i := range *req.MenuItem {
				item := &(*req.MenuItem)[i]
				item.MenuSubCategoryID = subCategory.ID

				if _, err := s.AddMenuItemInSubCategory(txCtx, item); err != nil {
					return err
				}
			}
		}

		res = subCategory.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

func (s *menuService) AddMenuItemInSubCategory(ctx context.Context, req *hospitality_menu_contracts.AddMenuItemInput) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		menuItem := &hospitalityMenu_models.MenuItem{
			ID:                uuid.New(),
			MenuGroupID:       req.MenuID,
			MenuSubCategoryID: req.MenuSubCategoryID,
			Name:              req.Name,
			Description:       req.Description,
			BasePrice:         req.BasePrice,
			IsAvailable:       true,
			Unit:              req.Unit,
			DietaryType:       req.DietaryType,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		if err := s.menuRepo.CreateMenuItem(txCtx, menuItem); err != nil {
			return err
		}

		if req.VariantGroup != nil {
			for i := range *req.VariantGroup {
				vg := &(*req.VariantGroup)[i]
				vg.MenuItemID = menuItem.ID

				if _, err := s.AddVariantGroup(txCtx, vg); err != nil {
					return err
				}
			}
		}

		if req.AddOns != nil {
			for i := range *req.AddOns {
				addon := &(*req.AddOns)[i]
				addon.MenuItemID = menuItem.ID

				if _, err := s.AddAddOnInMenu(txCtx, addon); err != nil {
					return err
				}
			}
		}

		res = menuItem.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

func (s *menuService) AddVariantGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantGroupInput) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		variantGrp := &hospitalityMenu_models.MenuVariantGroup{
			ID:            uuid.New(),
			MenuItemID:    req.MenuItemID,
			Name:          req.VariantGroupName,
			IsMultiSelect: true,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := s.menuRepo.CreateVariantGroup(txCtx, variantGrp); err != nil {
			return err
		}

		if req.Variant != nil {
			for i := range *req.Variant {
				v := &(*req.Variant)[i]
				v.VariantGroupID = variantGrp.ID
				v.MenuItemID = variantGrp.MenuItemID

				if _, err := s.AddVariantsInVariantsGroup(txCtx, v); err != nil {
					return err
				}
			}
		}

		res = variantGrp.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

func (s *menuService) AddVariantsInVariantsGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantInput) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		variant := &hospitalityMenu_models.MenuVariant{
			ID:                 uuid.New(),
			MenuVariantGroupID: req.VariantGroupID,
			MenuItemID:         req.MenuItemID,
			Name:               req.Name,
			Price:              req.Price,
			IsActive:           true,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		if err := s.menuRepo.CreateVariant(txCtx, variant); err != nil {
			return err
		}

		res = variant.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}

func (s *menuService) AddAddOnInMenu(ctx context.Context, req *hospitality_menu_contracts.AddAddOnInput) (uuid.UUID, error) {

	var res uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		addOn := &hospitalityMenu_models.MenuItemAddon{
			ID:         uuid.New(),
			MenuItemID: req.MenuItemID,
			Name:       req.Name,
			Price:      req.Price,
			IsVeg:      true,
			IsActive:   true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := s.menuRepo.CreateAddon(txCtx, addOn); err != nil {
			return err
		}

		res = addOn.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return res, nil
}
