package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"

	"github.com/google/uuid"
)

func (s *menuService) GetMenu(ctx context.Context, menuID uuid.UUID) (*hospitality_menu_contracts.MenuOutput, error) {
	// 1. Get the Root Menu
	menu, err := s.menuRepo.GetMenuByID(ctx, menuID)
	if err != nil {
		return nil, err
	}

	output := &hospitality_menu_contracts.MenuOutput{
		ID:              menu.ID,
		MenuName:        menu.Name,
		MenuDescription: &menu.Description,
		Categories:      []hospitality_menu_contracts.CategoryOutput{},
	}

	// 2. Get Categories
	categories, err := s.menuRepo.GetAllCategoriesByMenuID(ctx, menu.ID)
	if err != nil {
		return nil, err
	}

	for _, cat := range categories {
		catOut := hospitality_menu_contracts.CategoryOutput{
			ID:                  cat.ID,
			MenuID:              cat.MenuGroupID,
			CategoryName:        cat.Name,
			CategoryDescription: cat.Description,
			SubCategories:       []hospitality_menu_contracts.SubCategoryOutput{},
		}

		// 3. Get Sub-Categories
		subCats, err := s.menuRepo.GetSubCategoriesByCategoryID(ctx, cat.ID)
		if err != nil {
			return nil, err
		}

		for _, sub := range subCats {
			subOut := hospitality_menu_contracts.SubCategoryOutput{
				ID:                sub.ID,
				MenuCategoryID:    sub.MenuCategoryID,
				SubCatName:        sub.Name,
				SubCatDescription: sub.Description,
				MenuItems:         []hospitality_menu_contracts.MenuItemOutput{},
			}

			// 4. Get Menu Items
			items, err := s.menuRepo.GetItemsBySubCategoryID(ctx, sub.ID)
			if err != nil {
				return nil, err
			}

			for _, item := range items {
				itemOut := hospitality_menu_contracts.MenuItemOutput{
					ID:                item.ID,
					MenuID:            item.MenuGroupID,
					MenuSubCategoryID: item.MenuSubCategoryID,
					Name:              item.Name,
					Description:       item.Description,
					Unit:              item.Unit,
					BasePrice:         item.BasePrice,
					VariantGroups:     []hospitality_menu_contracts.VariantGroupOutput{},
					AddOns:            []hospitality_menu_contracts.AddOnOutput{},
				}

				// 5. Get Variant Groups
				vGroups, err := s.menuRepo.GetVariantGroupsByMenuItemID(ctx, item.ID)
				if err != nil {
					return nil, err
				}

				for _, vg := range vGroups {
					vgOut := hospitality_menu_contracts.VariantGroupOutput{
						ID:               vg.ID,
						MenuItemID:       vg.MenuItemID,
						VariantGroupName: vg.Name,
						Variants:         []hospitality_menu_contracts.VariantOutput{},
					}

					// 6. Get Variants
					vars, err := s.menuRepo.GetVariantsByVariantGroupID(ctx, vg.ID)
					if err != nil {
						return nil, err
					}
					for _, v := range vars {
						vgOut.Variants = append(vgOut.Variants, hospitality_menu_contracts.VariantOutput{
							ID:             v.ID,
							VariantGroupID: v.MenuVariantGroupID,
							MenuItemID:     v.MenuItemID,
							Name:           v.Name,
							Price:          v.Price,
						})
					}
					itemOut.VariantGroups = append(itemOut.VariantGroups, vgOut)
				}

				// 7. Get Add-ons
				addOns, err := s.menuRepo.GetAddOnsByMenuItemID(ctx, item.ID)
				if err != nil {
					return nil, err
				}
				for _, ao := range addOns {
					itemOut.AddOns = append(itemOut.AddOns, hospitality_menu_contracts.AddOnOutput{
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

func (s *menuService) GetMenusOfOwners(ctx context.Context, ownersID uuid.UUID) ([]hospitality_menu_contracts.MenuOutput, error) {
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
		fullMenu, err := s.GetMenu(ctx, m.ID)
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
