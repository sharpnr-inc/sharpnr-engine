package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	"time"

	"github.com/google/uuid"
)

// Create Menu and returns the Created ID of Created Menu.
func (s *menuService) CreateMenu(ctx context.Context, req *hospitality_menu_contracts.CreateMenuInput) (menuID uuid.UUID, err error) {

	menuGrp := &hospitalityMenu_models.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: req.OrganizationID,
		PropertyID:     req.PropertyID,
		Name:           req.MenuName,
		Description:    *req.MenuDescription,
		OwnerModule:    req.OwnerModuleType,
		OwnerID:        req.OwnerID,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := s.menuRepo.CreateMenuGroup(ctx, menuGrp); err != nil {
		return uuid.Nil, err
	}

	return menuGrp.ID, nil
}

func (s *menuService) AddCategoryInMenu(ctx context.Context, req *hospitality_menu_contracts.AddMenuCategoryInput) (categoryID uuid.UUID, err error) {

	cat := &hospitalityMenu_models.MenuCategory{
		ID:          uuid.New(),
		MenuGroupID: req.MenuID,
		Name:        req.CategoryName,
		Description: req.CategoryDescription,
		IsActive:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.menuRepo.CreateCategory(ctx, cat); err != nil {
		return uuid.Nil, err
	}

	return cat.ID, nil
}

func (s *menuService) AddSubCategoryInCategory(ctx context.Context, req *hospitality_menu_contracts.AddSubCategoryInput) (subCategoryID uuid.UUID, err error) {

	subCategory := &hospitalityMenu_models.MenuSubCategory{
		ID:             uuid.New(),
		MenuCategoryID: req.MenuCategoryID,
		Name:           req.SubCatName,
		Description:    req.SubCatDescription,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.menuRepo.CreateSubCategory(ctx, subCategory); err != nil {
		return uuid.Nil, err
	}

	return subCategory.ID, nil
}

func (s *menuService) AddMenuItemInSubCategory(ctx context.Context, req *hospitality_menu_contracts.AddMenuItemInput) (menuItemID uuid.UUID, err error) {

	menuItem := &hospitalityMenu_models.MenuItem{
		ID:                uuid.New(),
		MenuGroupID:       req.MenuID,
		MenuSubCategoryID: req.MenuSubCategoryID,
		Name:              req.Name,
		Description:       req.Description,
		IsVeg:             true,
		BasePrice:         req.BasePrice,
		IsAvailable:       true,
		Unit:              req.Unit,
		IsActive:          true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := s.menuRepo.CreateMenuItem(ctx, menuItem); err != nil {
		return uuid.Nil, err
	}

	return menuItem.ID, nil
}

func (s *menuService) AddVariantGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantGroupInput) (variantGroupID uuid.UUID, err error) {

	variantGrp := &hospitalityMenu_models.MenuVariantGroup{
		ID:            uuid.New(),
		MenuItemID:    req.MenuItemID,
		Name:          req.VariantGroupName,
		IsMultiSelect: true,
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.menuRepo.CreateVariantGroup(ctx, variantGrp); err != nil {
		return uuid.Nil, err
	}

	return variantGrp.ID, nil
}

func (s *menuService) AddVariantsInVariantsGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantInput) (variantID uuid.UUID, err error) {

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

	if err := s.menuRepo.CreateVariant(ctx, variant); err != nil {
		return uuid.Nil, err
	}
	return variant.ID, nil
}

func (s *menuService) AddAddOnInMenu(ctx context.Context, req *hospitality_menu_contracts.AddAddOnInput) (AddonID uuid.UUID, err error) {

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

	if err := s.menuRepo.CreateAddon(ctx, addOn); err != nil {
		return uuid.Nil, err
	}

	return addOn.ID, nil
}
