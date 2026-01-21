package menuservice

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

type MenuService interface {
	// -------------------------------
	// 🌐 MENU GROUPS
	// -------------------------------
	CreateMenuGroup(ctx context.Context, group *hospitalityMenu_models.MenuGroup) error

	GetMenuGroupByID(id uuid.UUID) (*hospitalityMenu_models.MenuGroup, error)
	UpdateMenuGroup(group *hospitalityMenu_models.MenuGroup) error
	DeleteMenuGroup(id uuid.UUID) error

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	CreateCategory(ctx context.Context, orgId uuid.UUID) error

	GetCategoryByID(id uuid.UUID) (*hospitalityMenu_models.MenuCategory, error)
	UpdateCategory(category *hospitalityMenu_models.MenuCategory) error
	DeleteCategory(id uuid.UUID) error

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------

	GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalityMenu_models.MenuSubCategory, error)
	GetSubCategoryByID(id uuid.UUID) (*hospitalityMenu_models.MenuSubCategory, error)
	UpdateSubCategory(subCat *hospitalityMenu_models.MenuSubCategory) error
	DeleteSubCategory(id uuid.UUID) error

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------

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

	// -------------------------------
	// 🧀 ADDONS
	// -------------------------------

	UpdateAddonOfMenuItem(ctx context.Context, addon *hospitalityMenu_models.MenuItemAddon) error
	DeleteAddonOfMenuItem(ctx context.Context, id uuid.UUID) error
}
