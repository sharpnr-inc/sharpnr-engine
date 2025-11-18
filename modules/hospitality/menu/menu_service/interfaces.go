package menuservice

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type MenuService interface {
	// -------------------------------
	// 🌐 MENU GROUPS
	// -------------------------------
	CreateMenuGroup(ctx context.Context, group *hospitalitymodels.MenuGroup) error
	GetAllMenuGroups(ctx context.Context, orgID uuid.UUID) ([]hospitalitydtos.Menu_GetAllMenuGroup_RESPONSE, error)
	GetMenuGroupByID(id uuid.UUID) (*hospitalitymodels.MenuGroup, error)
	UpdateMenuGroup(group *hospitalitymodels.MenuGroup) error
	DeleteMenuGroup(id uuid.UUID) error

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	CreateCategory(ctx context.Context, req *hospitalitydtos.Menu_CreateCategoryINMenuGroup_REQ, orgId uuid.UUID) error
	GetCategoriesByGroup(ctx context.Context, orgID uuid.UUID, groupID uuid.UUID) ([]hospitalitydtos.Menu_Category_RESPONSE, error)
	GetCategoryByID(id uuid.UUID) (*hospitalitymodels.MenuCategory, error)
	UpdateCategory(category *hospitalitymodels.MenuCategory) error
	DeleteCategory(id uuid.UUID) error

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------
	CreateSubCategory(subCat *hospitalitymodels.MenuSubCategory) error
	GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalitymodels.MenuSubCategory, error)
	GetSubCategoryByID(id uuid.UUID) (*hospitalitymodels.MenuSubCategory, error)
	UpdateSubCategory(subCat *hospitalitymodels.MenuSubCategory) error
	DeleteSubCategory(id uuid.UUID) error

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------
	CreateMenuItem(item *hospitalitymodels.MenuItem) error
	GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalitymodels.MenuItem, error)
	GetMenuItemByID(id uuid.UUID) (*hospitalitymodels.MenuItem, error)
	UpdateMenuItem(item *hospitalitymodels.MenuItem) error
	DeleteMenuItem(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS
	// -------------------------------
	CreateVariant(variant *hospitalitymodels.MenuVariant) error
	GetVariantsByItem(itemID uuid.UUID) ([]hospitalitymodels.MenuVariant, error)
	UpdateVariant(variant *hospitalitymodels.MenuVariant) error
	DeleteVariant(id uuid.UUID) error

	// -------------------------------
	// ➕ ADDON GROUPS
	// -------------------------------
	CreateAddonGroup(group *hospitalitymodels.AddonGroup) error
	GetAddonGroupsByItem(itemID uuid.UUID) ([]hospitalitymodels.AddonGroup, error)
	UpdateAddonGroup(group *hospitalitymodels.AddonGroup) error
	DeleteAddonGroup(id uuid.UUID) error

	// -------------------------------
	// 🧀 ADDONS
	// -------------------------------
	CreateAddon(addon *hospitalitymodels.Addon) error
	GetAddonsByGroup(groupID uuid.UUID) ([]hospitalitymodels.Addon, error)
	UpdateAddon(addon *hospitalitymodels.Addon) error
	DeleteAddon(id uuid.UUID) error
}
