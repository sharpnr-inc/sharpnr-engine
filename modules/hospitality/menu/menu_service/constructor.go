package menuservice

import (
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"

	"github.com/google/uuid"
)

type menuService struct {
	menuRepo menurepo.MenuRepositoryInterface
}

// CreateVariant implements MenuService.
func (m *menuService) CreateVariant(variant *hospitalitymodels.MenuVariant) error {
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
