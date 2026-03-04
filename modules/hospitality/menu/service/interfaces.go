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

	// Queries
	GetMenu(ctx context.Context, menuID uuid.UUID) (*hospitality_menu_contracts.MenuOutput, error)
	GetMenusOfOwners(ctx context.Context, ownersID uuid.UUID) ([]hospitality_menu_contracts.MenuOutput, error)

	MenuKitchenRegisterServiceInterface
}

type MenuKitchenRegisterServiceInterface interface {
	SetKitchenToMenuGroup(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuGroupInput) error
	SetKitchenToMenuCategory(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuCategoryInput) error
	SetKitchenToMenuSubCategory(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuSubCategoryInput) error
	SetKitchenToMenuItem(ctx context.Context, req *hospitality_menu_contracts.SetKitchenToMenuItemInput) error
	/* ---------------------------------- <***> --------------------------------- */
	GetKitchenfOfMenuGroup(ctx context.Context, menuGroupID uuid.UUID) (uuid.UUID, error)
	GetKitchenOfMenuCategory(ctx context.Context, menuCategoryID uuid.UUID) (uuid.UUID, error)
	GetKitchenOfMenuSubCategory(ctx context.Context, menuSubCategoryID uuid.UUID) (uuid.UUID, error)
	GetKitchenOfMenuItem(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error)
}
