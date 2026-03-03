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
	GetMenuByOwnersID(ctx context.Context, ownersID uuid.UUID) ([]*hospitalityMenu_models.MenuGroup, error)
	GetMenuByID(ctx context.Context, menuID uuid.UUID) (*hospitalityMenu_models.MenuGroup, error)
	UpdateMenuGroup(group *hospitalityMenu_models.MenuGroup) error
	DeleteMenuGroup(id uuid.UUID) error

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	CreateCategory(ctx context.Context, category *hospitalityMenu_models.MenuCategory) error
	GetAllCategoriesByMenuID(ctx context.Context, menuID uuid.UUID) ([]hospitalityMenu_models.MenuCategory, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuCategory, error)
	UpdateCategory(category *hospitalityMenu_models.MenuCategory) error
	DeleteCategory(id uuid.UUID) error
	/* -------------------------------- Register -------------------------------- */
	CreateMenuCategoryKitchenRegister(ctx context.Context, entry *hospitalityMenu_models.MenuCategoryKitchenRegister) error
	GetCategoryKitchenByMenuItemID(ctx context.Context, menuCategoryID uuid.UUID, kitchenID uuid.UUID) (*hospitalityMenu_models.MenuCategoryKitchenRegister, error)
	GetCategoryPreferredKitchen(ctx context.Context, menuCategoryID uuid.UUID) (*hospitalityMenu_models.MenuCategoryKitchenRegister, error)

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------
	CreateSubCategory(ctx context.Context, sub *hospitalityMenu_models.MenuSubCategory) error
	GetSubCategoriesByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]hospitalityMenu_models.MenuSubCategory, error)
	GetSubCategoryByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuSubCategory, error)
	UpdateSubCategory(subCat *hospitalityMenu_models.MenuSubCategory) error
	DeleteSubCategory(id uuid.UUID) error
	/* -------------------------------- Register -------------------------------- */
	CreateMenuSubCategoryKitchenRegister(ctx context.Context, entry *hospitalityMenu_models.MenuSubCategoryKitchenRegister) error
	GetSubCategoryKitchenByMenuItemID(ctx context.Context, menuSubCategoryID uuid.UUID, kitchenID uuid.UUID) (*hospitalityMenu_models.MenuSubCategoryKitchenRegister, error)
	GetSubCategoryPreferredKitchen(ctx context.Context, menuSubCategoryID uuid.UUID) (*hospitalityMenu_models.MenuSubCategoryKitchenRegister, error)

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------
	CreateMenuItem(ctx context.Context, item *hospitalityMenu_models.MenuItem) error
	GetItemsBySubCategoryID(ctx context.Context, subCatID uuid.UUID) ([]hospitalityMenu_models.MenuItem, error)
	GetMenuItemByID(ctx context.Context, id uuid.UUID) (*hospitalityMenu_models.MenuItem, error)
	UpdateMenuItem(item *hospitalityMenu_models.MenuItem) error
	DeleteMenuItem(id uuid.UUID) error
	/* -------------------------------- Register -------------------------------- */
	CreateEntryInKitchenMenuRegister(ctx context.Context, entry *hospitalityMenu_models.MenuItemKitchenRegister) error
	GetKitchenMenuRegisterByMenuItemID(ctx context.Context, menuItemID uuid.UUID, kitchenID uuid.UUID) (*hospitalityMenu_models.MenuItemKitchenRegister, error)
	GetMenuItemKitchen(ctx context.Context, menuItemID uuid.UUID) (*hospitalityMenu_models.MenuItemKitchenRegister, error)

	// -------------------------------
	// 🧩 VARIANTS
	// -------------------------------
	CreateVariant(ctx context.Context, variant *hospitalityMenu_models.MenuVariant) error
	GetVariantsByItem(itemID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error)
	GetVariantsByVariantGroupID(ctx context.Context, variantGroupID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error)
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
	GetAddOnsByMenuItemID(ctx context.Context, menuItemID uuid.UUID) ([]hospitalityMenu_models.MenuItemAddon, error)
	UpdateAddon(addon *hospitalityMenu_models.MenuItemAddon) error
	DeleteAddon(id uuid.UUID) error
}
