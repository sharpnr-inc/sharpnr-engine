package menuservice

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

type MenuService interface {
	// -------------------------------
	// 🌐 MENU GROUPS
	// -------------------------------
	CreateMenuGroup(ctx context.Context, group *hospitalityMenu_models.MenuGroup) error
	GetAllMenuGroups(ctx context.Context, orgID uuid.UUID) ([]hospitalitydtos.Menu_GetAllMenuGroup_RESPONSE, error)
	GetMenuGroupByID(id uuid.UUID) (*hospitalityMenu_models.MenuGroup, error)
	UpdateMenuGroup(group *hospitalityMenu_models.MenuGroup) error
	DeleteMenuGroup(id uuid.UUID) error

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	CreateCategory(ctx context.Context, req *hospitalitydtos.Menu_CreateCategoryINMenuGroup_REQ, orgId uuid.UUID) error
	GetCategoriesByGroup(ctx context.Context, orgID uuid.UUID, groupID uuid.UUID) ([]hospitalitydtos.Menu_Category_RESPONSE, error)
	GetCategoryByID(id uuid.UUID) (*hospitalityMenu_models.MenuCategory, error)
	UpdateCategory(category *hospitalityMenu_models.MenuCategory) error
	DeleteCategory(id uuid.UUID) error

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------
	CreateSubCategory(
		ctx context.Context,
		req *hospitalitydtos.Menu_CreateSubCategory_REQ,
		orgID uuid.UUID,
	) error
	GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalityMenu_models.MenuSubCategory, error)
	GetSubCategoryByID(id uuid.UUID) (*hospitalityMenu_models.MenuSubCategory, error)
	UpdateSubCategory(subCat *hospitalityMenu_models.MenuSubCategory) error
	DeleteSubCategory(id uuid.UUID) error

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------
	CreateMenuItem(context.Context, *hospitalitydtos.Menu_CreateMenuItem_REQ, uuid.UUID) (*hospitalitydtos.Menu_CreateMenuItem_RES, error)
	GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalityMenu_models.MenuItem, error)
	GetMenuItemByID(id uuid.UUID) (*hospitalityMenu_models.MenuItem, error)
	UpdateMenuItem(item *hospitalityMenu_models.MenuItem) error
	DeleteMenuItem(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS
	// -------------------------------
	CreateVariant(variant *hospitalityMenu_models.MenuVariant) error
	GetVariantsByItem(ctx context.Context, itemID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error)
	UpdateVariant(variant *hospitalityMenu_models.MenuVariant) error
	DeleteVariant(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS GROUP
	// -------------------------------
	CreateVariantGroup(ctx context.Context, menuItemId uuid.UUID, req *hospitalitydtos.Menu_CreateVariantGroupAndVariant_REQ) error
	GetVariantGroupByMenuId(ctx context.Context, orgid, menuItemId uuid.UUID) ([]hospitalitydtos.Menu_VariantGroup_RESPONSE, error)

	// -------------------------------
	// 🧀 ADDONS
	// -------------------------------
	CreateAddForMenuItem(ctx context.Context, orgid uuid.UUID, req hospitalitydtos.Menu_CreateAddonsInMenu_REQUEST) error
	UpdateAddonOfMenuItem(ctx context.Context, addon *hospitalityMenu_models.MenuItemAddon) error
	DeleteAddonOfMenuItem(ctx context.Context, id uuid.UUID) error
}
