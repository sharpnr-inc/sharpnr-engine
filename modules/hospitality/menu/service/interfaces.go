package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(ctx context.Context, req *hospitality_menu_contracts.CreateMenuInput) (menuID uuid.UUID, err error)
	AddCategoryInMenu(ctx context.Context, req *hospitality_menu_contracts.AddMenuCategoryInput) (categoryID uuid.UUID, err error)
	AddSubCategoryInCategory(ctx context.Context, req *hospitality_menu_contracts.AddSubCategoryInput) (subCategoryID uuid.UUID, err error)
	AddMenuItemInSubCategory(ctx context.Context, req *hospitality_menu_contracts.AddMenuItemInput) (menuItemID uuid.UUID, err error)
	AddVariantGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantGroupInput) (variantGroupID uuid.UUID, err error)
	AddVariantsInVariantsGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantInput) (variantID uuid.UUID, err error)
	AddAddOnInMenu(ctx context.Context, req *hospitality_menu_contracts.AddAddOnInput) (AddonID uuid.UUID, err error)
	SetKitchenToMenuItem(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuItemInput) error
	SetKitchenToWholeMenu(ctx context.Context, ownersID uuid.UUID, req *hospitality_menu_contracts.SetKitchenToWholeMenuInput) error

	// Queries
	GetMenu(ctx context.Context, menuID uuid.UUID) (*hospitality_menu_contracts.MenuOutput, error)
	GetMenusOfOwners(ctx context.Context, ownersID uuid.UUID) ([]hospitality_menu_contracts.MenuOutput, error)
}
