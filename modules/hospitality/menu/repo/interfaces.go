package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

type MenuRepositoryInterface interface {
	// -------------------------------
	// 🌐 MENU GROUPS
	// -------------------------------
	CreateMenuGroup(ctx context.Context, group *hospitalityMenu_models.MenuGroup) error
	GetAllMenuGroups(ctx context.Context, propertyID uuid.UUID) ([]*hospitalityMenu_models.MenuGroup, error)

	GetMenuGroupByID(ctx context.Context, groupID uuid.UUID) (*hospitalityMenu_models.MenuGroup, error)
	UpdateMenuGroup(group *hospitalityMenu_models.MenuGroup) error
	DeleteMenuGroup(id uuid.UUID) error

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	CreateCategory(ctx context.Context, category *hospitalityMenu_models.MenuCategory) error
	GetCategoriesByGroup(ctx context.Context, groupID uuid.UUID) ([]hospitalityMenu_models.MenuCategory, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuCategory, error)
	UpdateCategory(category *hospitalityMenu_models.MenuCategory) error
	DeleteCategory(id uuid.UUID) error

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------
	CreateSubCategory(ctx context.Context, sub *hospitalityMenu_models.MenuSubCategory) error
	GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalityMenu_models.MenuSubCategory, error)
	GetSubCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuSubCategory, error)
	UpdateSubCategory(subCat *hospitalityMenu_models.MenuSubCategory) error
	DeleteSubCategory(id uuid.UUID) error

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------
	CreateMenuItem(ctx context.Context, item *hospitalityMenu_models.MenuItem) error
	GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalityMenu_models.MenuItem, error)
	GetMenuItemByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuItem, error)
	UpdateMenuItem(item *hospitalityMenu_models.MenuItem) error
	DeleteMenuItem(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS
	// -------------------------------
	CreateVariant(ctx context.Context, variant *hospitalityMenu_models.MenuVariant) error
	GetVariantsByItem(itemID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error)
	UpdateVariant(variant *hospitalityMenu_models.MenuVariant) error
	DeleteVariant(id uuid.UUID) error

	// -------------------------------
	// 🧩 VARIANTS GROUP
	// -------------------------------
	CreateVariantGroup(ctx context.Context, variant *hospitalityMenu_models.MenuVariantGroup) error
	GetVariantGroupsByMenuItemID(ctx context.Context, menuId uuid.UUID) ([]hospitalityMenu_models.MenuVariantGroup, error)
	UpdateVariantGroup(ctx context.Context, variant *hospitalityMenu_models.MenuVariantGroup) error

	// -------------------------------
	// 🧀 ADDONS
	// -------------------------------
	CreateAddon(ctx context.Context, addon *hospitalityMenu_models.MenuItemAddon) error
	UpdateAddon(addon *hospitalityMenu_models.MenuItemAddon) error
	DeleteAddon(id uuid.UUID) error
}
