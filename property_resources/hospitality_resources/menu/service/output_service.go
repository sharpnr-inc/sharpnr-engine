package menuservice

import (
	"context"
	"errors"
	hospitality_menu_contracts "sharpnr-engine/property_resources/hospitality_resources/menu/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *menuService) GetMenu(ctx context.Context, orgID, emplID, menuID uuid.UUID) (*hospitality_menu_contracts.MenuOutput, error) {

	menu, err := s.menuRepo.GetMenuByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	var menuPreferredKitchen *uuid.UUID
	menuKitchenID, err := s.GetKitchenfOfMenuGroup(ctx, menu.ID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		menuPreferredKitchen = &menuKitchenID
	}

	output := &hospitality_menu_contracts.MenuOutput{
		ID:                 menu.ID,
		MenuName:           menu.Name,
		MenuDescription:    &menu.Description,
		PreferredKitchenID: menuPreferredKitchen,
		Categories:         make([]hospitality_menu_contracts.CategoryOutput, 0),
	}

	categories, err := s.menuRepo.GetAllCategoriesByMenuID(ctx, menu.ID)
	if err != nil {
		return nil, err
	}

	for _, cat := range categories {

		// Preferred Kitchen (Category)
		var categoryPreferredKitchen *uuid.UUID
		categoryKitchenID, err := s.GetKitchenOfMenuCategory(ctx, cat.ID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		} else {
			categoryPreferredKitchen = &categoryKitchenID
		}

		catOut := hospitality_menu_contracts.CategoryOutput{
			ID:                  cat.ID,
			MenuID:              cat.MenuGroupID,
			CategoryName:        cat.Name,
			CategoryDescription: cat.Description,
			PreferredKitchenID:  categoryPreferredKitchen,
			SubCategories:       make([]hospitality_menu_contracts.SubCategoryOutput, 0),
		}

		subCats, err := s.menuRepo.GetSubCategoriesByCategoryID(ctx, cat.ID)
		if err != nil {
			return nil, err
		}

		for _, sub := range subCats {

			// Preferred Kitchen (SubCategory)
			var subCategoryPreferredKitchen *uuid.UUID
			subKitchenID, err := s.GetKitchenOfMenuSubCategory(ctx, sub.ID)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}
			} else {
				subCategoryPreferredKitchen = &subKitchenID
			}

			subOut := hospitality_menu_contracts.SubCategoryOutput{
				ID:                 sub.ID,
				MenuCategoryID:     sub.MenuCategoryID,
				SubCatName:         sub.Name,
				SubCatDescription:  sub.Description,
				PreferredKitchenID: subCategoryPreferredKitchen,
				MenuItems:          make([]hospitality_menu_contracts.MenuItemOutput, 0),
			}

			items, err := s.menuRepo.GetItemsBySubCategoryID(ctx, sub.ID)
			if err != nil {
				return nil, err
			}

			for _, item := range items {

				// 1. Force initialization to an empty slice (This kills the "null" in JSON)
				itemTaxDetails := make([]hospitality_menu_contracts.ItemTaxDetails, 0)

				// 2. Resolve Tax (Tracing through the hierarchy)
				// taxCatID, _, err := s.GetTaxCategoryOfMenuItem(ctx, orgID, emplID, item.ID)

				// Only fetch breakdown if a Category ID was actually found
				// if err == nil && taxCatID != uuid.Nil {
				// 	breakdown, err := s.taxEngine.GetTaxCategoryTaxBreakdown(ctx, taxCatID, item.BasePrice)
				// 	if err == nil && breakdown != nil {
				// 		for _, b := range breakdown {
				// 			itemTaxDetails = append(itemTaxDetails, hospitality_menu_contracts.ItemTaxDetails{
				// 				TaxComponentName: b.Name,
				// 				TaxRate:          b.Rate,
				// 			})
				// 		}
				// 	}
				// }

				kitchenReg, err := s.menuRepo.GetMenuItemKitchen(ctx, item.ID)
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}

				var servedBy uuid.UUID
				if kitchenReg != nil {
					servedBy = kitchenReg.KitchenID
				} else {
					servedBy = uuid.Nil
				}

				itemOut := hospitality_menu_contracts.MenuItemOutput{
					ID:                item.ID,
					MenuID:            item.MenuGroupID,
					MenuSubCategoryID: item.MenuSubCategoryID,
					Name:              item.Name,
					Description:       item.Description,
					Unit:              item.Unit,
					BasePrice:         item.BasePrice,
					TaxDetails:        itemTaxDetails,
					DietaryType:       item.DietaryType,
					ServedBy:          &servedBy,
					VariantGroups:     make([]hospitality_menu_contracts.VariantGroupOutput, 0),
					AddOns:            make([]hospitality_menu_contracts.AddOnOutput, 0),
				}

				vGroups, err := s.menuRepo.GetVariantGroupsByMenuItemID(ctx, item.ID)
				if err != nil {
					return nil, err
				}

				for _, vg := range vGroups {

					vgOut := hospitality_menu_contracts.VariantGroupOutput{
						ID:               vg.ID,
						MenuItemID:       vg.MenuItemID,
						VariantGroupName: vg.Name,
						Variants:         make([]hospitality_menu_contracts.VariantOutput, 0),
					}

					vars, err := s.menuRepo.GetVariantsByVariantGroupID(ctx, vg.ID)
					if err != nil {
						return nil, err
					}

					for _, v := range vars {
						vgOut.Variants = append(vgOut.Variants,
							hospitality_menu_contracts.VariantOutput{
								ID:             v.ID,
								VariantGroupID: v.MenuVariantGroupID,
								MenuItemID:     v.MenuItemID,
								Name:           v.Name,
								Price:          v.Price,
							})
					}

					itemOut.VariantGroups = append(itemOut.VariantGroups, vgOut)
				}

				addOns, err := s.menuRepo.GetAddOnsByMenuItemID(ctx, item.ID)
				if err != nil {
					return nil, err
				}

				for _, ao := range addOns {
					itemOut.AddOns = append(itemOut.AddOns,
						hospitality_menu_contracts.AddOnOutput{
							ID:         ao.ID,
							MenuItemID: ao.MenuItemID,
							Name:       ao.Name,
							Price:      ao.Price,
						})
				}

				subOut.MenuItems = append(subOut.MenuItems, itemOut)
			}

			catOut.SubCategories = append(catOut.SubCategories, subOut)
		}

		output.Categories = append(output.Categories, catOut)
	}

	return output, nil
}

func (s *menuService) GetMenusOfOwners(ctx context.Context, orgID, emplID, ownersID uuid.UUID) ([]hospitality_menu_contracts.MenuOutput, error) {
	// Initialize as an empty slice to ensure JSON returns [] instead of null
	finalMenus := []hospitality_menu_contracts.MenuOutput{}

	// 1. Get all root menu groups for this owner
	menus, err := s.menuRepo.GetMenuByOwnersID(ctx, ownersID)
	if err != nil {
		// Return empty slice and error if DB call fails
		return finalMenus, err
	}

	// 2. For each menu group, build the full nested tree
	for _, m := range menus {
		// We reuse the existing GetMenu logic
		fullMenu, err := s.GetMenu(ctx, orgID, emplID, m.ID)
		if err != nil {
			// If one menu fails to load, we can skip it or return error
			continue
		}

		if fullMenu != nil {
			finalMenus = append(finalMenus, *fullMenu)
		}
	}

	return finalMenus, nil
}
