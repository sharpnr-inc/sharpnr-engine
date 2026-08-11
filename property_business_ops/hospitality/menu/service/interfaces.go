package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/property_business_ops/hospitality/menu/contracts"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(ctx context.Context, orgID, emplID uuid.UUID, req *hospitality_menu_contracts.CreateMenuInput) (menuID uuid.UUID, err error)
	AddCategoryInMenu(ctx context.Context, req *hospitality_menu_contracts.AddMenuCategoryInput) (categoryID uuid.UUID, err error)
	AddSubCategoryInCategory(ctx context.Context, req *hospitality_menu_contracts.AddSubCategoryInput) (subCategoryID uuid.UUID, err error)
	AddMenuItemInSubCategory(ctx context.Context, req *hospitality_menu_contracts.AddMenuItemInput) (menuItemID uuid.UUID, err error)
	AddVariantGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantGroupInput) (variantGroupID uuid.UUID, err error)
	AddVariantsInVariantsGroup(ctx context.Context, req *hospitality_menu_contracts.AddVariantInput) (variantID uuid.UUID, err error)
	AddAddOnInMenu(ctx context.Context, req *hospitality_menu_contracts.AddAddOnInput) (AddonID uuid.UUID, err error)

	// Queries
	GetMenu(ctx context.Context, orgID, emplID, menuID uuid.UUID) (*hospitality_menu_contracts.MenuOutput, error)
	GetMenusOfOwners(ctx context.Context, orgID, emplID, ownersID uuid.UUID) ([]hospitality_menu_contracts.MenuOutput, error)
	GetMenuItemParents(ctx context.Context, menuItemID uuid.UUID) (parents hospitality_menu_contracts.MenuItemParentsDetails, err error)

	MenuKitchenRegisterServiceInterface
	MenuTaxationRegisterServiceInterface
}

type MenuKitchenRegisterServiceInterface interface {
	SetKitchenToMenuGroup(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID, req *hospitality_menu_contracts.SetKitchenToMenuGroupInput) error
	SetKitchenToMenuCategory(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID, req *hospitality_menu_contracts.SetKitchenToMenuCategoryInput) error
	SetKitchenToMenuSubCategory(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID, req *hospitality_menu_contracts.SetKitchenToMenuSubCategoryInput) error
	SetKitchenToMenuItem(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, ownerID uuid.UUID, req *hospitality_menu_contracts.SetKitchenToMenuItemInput) error
	/* ---------------------------------- <***> --------------------------------- */
	GetKitchenfOfMenuGroup(ctx context.Context, menuGroupID uuid.UUID) (uuid.UUID, error)
	GetKitchenOfMenuCategory(ctx context.Context, menuCategoryID uuid.UUID) (uuid.UUID, error)
	GetKitchenOfMenuSubCategory(ctx context.Context, menuSubCategoryID uuid.UUID) (uuid.UUID, error)
	GetKitchenOfMenuItem(ctx context.Context, menuItemID uuid.UUID) (uuid.UUID, error)
}

type MenuTaxationRegisterServiceInterface interface {
	AddTaxCategoryInMenuGroup(
		ctx context.Context,
		orgID, emplID, menuGroupOwnerID uuid.UUID,
		req hospitality_menu_contracts.MenuGroupTaxCategoryRequestDto,
	) error

	// This returns the TaxCategoryID and LegalEntityID
	GetTaxCategoryOfMenuItem(
		ctx context.Context,
		orgID, emplID, menuItemID uuid.UUID,
	) (uuid.UUID, uuid.UUID, error)
}
