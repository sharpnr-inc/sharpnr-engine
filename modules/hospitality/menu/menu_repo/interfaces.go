package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type MenuRepositoryInterface interface {
	// -------------------------------
	// 🌐 MENU GROUPS
	// -------------------------------
	CreateMenuGroup(ctx context.Context, group *hospitalitymodels.MenuGroup) error
	GetAllMenuGroups(ctx context.Context, orgID uuid.UUID) ([]hospitalitymodels.MenuGroup, error)
	GetMenuGroupByID(ctx context.Context, groupID uuid.UUID) (*hospitalitymodels.MenuGroup, error)
	UpdateMenuGroup(group *hospitalitymodels.MenuGroup) error
	DeleteMenuGroup(id uuid.UUID) error

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	CreateCategory(ctx context.Context, category *hospitalitymodels.MenuCategory) error
	GetCategoriesByGroup(ctx context.Context, groupID uuid.UUID) ([]hospitalitymodels.MenuCategory, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.MenuCategory, error)
	UpdateCategory(category *hospitalitymodels.MenuCategory) error
	DeleteCategory(id uuid.UUID) error

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------
	CreateSubCategory(ctx context.Context, sub *hospitalitymodels.MenuSubCategory) error
	GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalitymodels.MenuSubCategory, error)
	GetSubCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.MenuSubCategory, error)
	UpdateSubCategory(subCat *hospitalitymodels.MenuSubCategory) error
	DeleteSubCategory(id uuid.UUID) error

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------
	CreateMenuItem(ctx context.Context, item *hospitalitymodels.MenuItem) error
	GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalitymodels.MenuItem, error)
	GetMenuItemByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.MenuItem, error)
	UpdateMenuItem(item *hospitalitymodels.MenuItem) error
	DeleteMenuItem(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS
	// -------------------------------
	CreateVariant(ctx context.Context, variant *hospitalitymodels.MenuVariant) error
	GetVariantsByItem(itemID uuid.UUID) ([]hospitalitymodels.MenuVariant, error)
	UpdateVariant(variant *hospitalitymodels.MenuVariant) error
	DeleteVariant(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS GROUP
	// -------------------------------
	CreateVariantGroup(ctx context.Context, variant *hospitalitymodels.MenuVariantGroup) error
	GetVariantGroupsByMenuItemID(ctx context.Context, menuId uuid.UUID) ([]hospitalitymodels.MenuVariantGroup, error)
	UpdateVariantGroup(ctx context.Context, variant *hospitalitymodels.MenuVariantGroup) error

	// -------------------------------
	// 🧀 ADDONS
	// -------------------------------
	CreateAddon(ctx context.Context, addon *hospitalitymodels.Addon) error
	UpdateAddon(addon *hospitalitymodels.Addon) error
	DeleteAddon(id uuid.UUID) error
}
