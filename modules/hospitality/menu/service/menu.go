package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	"time"

	"github.com/google/uuid"
)

// CreateMenu is the entry point that handles the recursive creation of the entire menu tree.
func (s *menuService) CreateMenu(ctx context.Context, req *hospitality_menu_contracts.CreateMenuInput) (uuid.UUID, error) {
	// 1. Create the Root Menu Group
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

	// 2. Process Categories if they exist
	if req.Categories != nil {
		for _, catInput := range *req.Categories {
			catInput.MenuID = menuGrp.ID // Inject parent ID
			if _, err := s.AddCategoryInMenu(ctx, &catInput); err != nil {
				return menuGrp.ID, err
			}
		}
	}

	return menuGrp.ID, nil
}

func (s *menuService) AddCategoryInMenu(ctx context.Context, req *hospitality_menu_contracts.AddMenuCategoryInput) (uuid.UUID, error) {
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

	// Process Sub-Categories
	if req.SubCategories != nil {
		for _, subInput := range *req.SubCategories {
			subInput.MenuCategoryID = cat.ID
			if _, err := s.AddSubCategoryInCategory(ctx, &subInput); err != nil {
				return cat.ID, err
			}
		}
	}

	return cat.ID, nil
}

func (s *menuService) AddSubCategoryInCategory(ctx context.Context, req *hospitality_menu_contracts.AddSubCategoryInput) (uuid.UUID, error) {
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

	// Process Menu Items
	if req.MenuItem != nil {
		for _, itemInput := range *req.MenuItem {
			itemInput.MenuSubCategoryID = subCategory.ID
			// Inherit MenuID from the parent logic if your DB schema requires it
			if _, err := s.AddMenuItemInSubCategory(ctx, &itemInput); err != nil {
				return subCategory.ID, err
			}
		}
	}

	return subCategory.ID, nil
}

func (s *menuService) AddMenuItemInSubCategory(ctx context.Context, req *hospitality_menu_contracts.AddMenuItemInput) (uuid.UUID, error) {
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

	// Process Variant Groups
	if req.VariantGroup != nil {
		for _, vgInput := range *req.VariantGroup {
			vgInput.MenuItemID = menuItem.ID
			if _, err := s.AddVariantGroup(ctx, &vgInput); err != nil {
				return menuItem.ID, err
			}
		}
	}

	// Process Add-Ons
	if req.AddOns != nil {
		for _, addOnInput := range *req.AddOns {
			addOnInput.MenuItemID = menuItem.ID
			if _, err := s.AddAddOnInMenu(ctx, &addOnInput); err != nil {
				return menuItem.ID, err
			}
		}
	}

	return menuItem.ID, nil
}

func (s *menuService) AddVariantGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantGroupInput) (uuid.UUID, error) {
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

	// Process Individual Variants
	if req.Variant != nil {
		for _, vInput := range *req.Variant {
			vInput.VariantGroupID = variantGrp.ID
			vInput.MenuItemID = variantGrp.MenuItemID
			if _, err := s.AddVariantsInVariantsGroup(ctx, &vInput); err != nil {
				return variantGrp.ID, err
			}
		}
	}

	return variantGrp.ID, nil
}

func (s *menuService) AddVariantsInVariantsGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantInput) (uuid.UUID, error) {
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

func (s *menuService) AddAddOnInMenu(ctx context.Context, req *hospitality_menu_contracts.AddAddOnInput) (uuid.UUID, error) {
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
