package menuservice

import (
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"

	"github.com/google/uuid"
)

type menuService struct {
	menuRepo menurepo.MenuRepositoryInterface
}

// CreateAddon implements MenuService.
func (m *menuService) CreateAddon(addon *hospitalitymodels.Addon) error {
	panic("unimplemented")
}

// CreateAddonGroup implements MenuService.
func (m *menuService) CreateAddonGroup(group *hospitalitymodels.AddonGroup) error {
	panic("unimplemented")
}

// CreateCategory implements MenuService.
func (m *menuService) CreateCategory(category *hospitalitymodels.MenuCategory) error {
	panic("unimplemented")
}

// CreateMenuGroup implements MenuService.
func (m *menuService) CreateMenuGroup(group *hospitalitymodels.MenuGroup) error {
	panic("unimplemented")
}

// CreateMenuItem implements MenuService.
func (m *menuService) CreateMenuItem(item *hospitalitymodels.MenuItem) error {
	panic("unimplemented")
}

// CreateSubCategory implements MenuService.
func (m *menuService) CreateSubCategory(subCat *hospitalitymodels.MenuSubCategory) error {
	panic("unimplemented")
}

// CreateVariant implements MenuService.
func (m *menuService) CreateVariant(variant *hospitalitymodels.MenuVariant) error {
	panic("unimplemented")
}

// DeleteAddon implements MenuService.
func (m *menuService) DeleteAddon(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteAddonGroup implements MenuService.
func (m *menuService) DeleteAddonGroup(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteCategory implements MenuService.
func (m *menuService) DeleteCategory(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteMenuGroup implements MenuService.
func (m *menuService) DeleteMenuGroup(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteMenuItem implements MenuService.
func (m *menuService) DeleteMenuItem(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteSubCategory implements MenuService.
func (m *menuService) DeleteSubCategory(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteVariant implements MenuService.
func (m *menuService) DeleteVariant(id uuid.UUID) error {
	panic("unimplemented")
}

// GetAddonGroupsByItem implements MenuService.
func (m *menuService) GetAddonGroupsByItem(itemID uuid.UUID) ([]hospitalitymodels.AddonGroup, error) {
	panic("unimplemented")
}

// GetAddonsByGroup implements MenuService.
func (m *menuService) GetAddonsByGroup(groupID uuid.UUID) ([]hospitalitymodels.Addon, error) {
	panic("unimplemented")
}

// GetAllMenuGroups implements MenuService.
func (m *menuService) GetAllMenuGroups(orgID uuid.UUID) ([]hospitalitymodels.MenuGroup, error) {
	panic("unimplemented")
}

// GetCategoriesByGroup implements MenuService.
func (m *menuService) GetCategoriesByGroup(groupID uuid.UUID) ([]hospitalitymodels.MenuCategory, error) {
	panic("unimplemented")
}

// GetCategoryByID implements MenuService.
func (m *menuService) GetCategoryByID(id uuid.UUID) (*hospitalitymodels.MenuCategory, error) {
	panic("unimplemented")
}

// GetItemsBySubCategory implements MenuService.
func (m *menuService) GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalitymodels.MenuItem, error) {
	panic("unimplemented")
}

// GetMenuGroupByID implements MenuService.
func (m *menuService) GetMenuGroupByID(id uuid.UUID) (*hospitalitymodels.MenuGroup, error) {
	panic("unimplemented")
}

// GetMenuItemByID implements MenuService.
func (m *menuService) GetMenuItemByID(id uuid.UUID) (*hospitalitymodels.MenuItem, error) {
	panic("unimplemented")
}

// GetSubCategoriesByCategory implements MenuService.
func (m *menuService) GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalitymodels.MenuSubCategory, error) {
	panic("unimplemented")
}

// GetSubCategoryByID implements MenuService.
func (m *menuService) GetSubCategoryByID(id uuid.UUID) (*hospitalitymodels.MenuSubCategory, error) {
	panic("unimplemented")
}

// GetVariantsByItem implements MenuService.
func (m *menuService) GetVariantsByItem(itemID uuid.UUID) ([]hospitalitymodels.MenuVariant, error) {
	panic("unimplemented")
}

// UpdateAddon implements MenuService.
func (m *menuService) UpdateAddon(addon *hospitalitymodels.Addon) error {
	panic("unimplemented")
}

// UpdateAddonGroup implements MenuService.
func (m *menuService) UpdateAddonGroup(group *hospitalitymodels.AddonGroup) error {
	panic("unimplemented")
}

// UpdateCategory implements MenuService.
func (m *menuService) UpdateCategory(category *hospitalitymodels.MenuCategory) error {
	panic("unimplemented")
}

// UpdateMenuGroup implements MenuService.
func (m *menuService) UpdateMenuGroup(group *hospitalitymodels.MenuGroup) error {
	panic("unimplemented")
}

// UpdateMenuItem implements MenuService.
func (m *menuService) UpdateMenuItem(item *hospitalitymodels.MenuItem) error {
	panic("unimplemented")
}

// UpdateSubCategory implements MenuService.
func (m *menuService) UpdateSubCategory(subCat *hospitalitymodels.MenuSubCategory) error {
	panic("unimplemented")
}

// UpdateVariant implements MenuService.
func (m *menuService) UpdateVariant(variant *hospitalitymodels.MenuVariant) error {
	panic("unimplemented")
}

func NewHospitalityMenuService(menuRepo menurepo.MenuRepositoryInterface) *menuService {
	return &menuService{
		menuRepo: menuRepo,
	}

}
