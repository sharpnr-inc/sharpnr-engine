package menuservice

import (
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

type menuService struct {
	menuRepo menurepo.MenuRepositoryInterface
}

// CreateVariant implements MenuService.
func (m *menuService) CreateVariant(variant *hospitalityMenu_models.MenuVariant) error {
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
func (m *menuService) GetCategoryByID(id uuid.UUID) (*hospitalityMenu_models.MenuCategory, error) {
	panic("unimplemented")
}

// GetItemsBySubCategory implements MenuService.
func (m *menuService) GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalityMenu_models.MenuItem, error) {
	panic("unimplemented")
}

// GetMenuGroupByID implements MenuService.
func (m *menuService) GetMenuGroupByID(id uuid.UUID) (*hospitalityMenu_models.MenuGroup, error) {
	panic("unimplemented")
}

// GetMenuItemByID implements MenuService.
func (m *menuService) GetMenuItemByID(id uuid.UUID) (*hospitalityMenu_models.MenuItem, error) {
	panic("unimplemented")
}

// GetSubCategoriesByCategory implements MenuService.
func (m *menuService) GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalityMenu_models.MenuSubCategory, error) {
	panic("unimplemented")
}

// GetSubCategoryByID implements MenuService.
func (m *menuService) GetSubCategoryByID(id uuid.UUID) (*hospitalityMenu_models.MenuSubCategory, error) {
	panic("unimplemented")
}

// UpdateCategory implements MenuService.
func (m *menuService) UpdateCategory(category *hospitalityMenu_models.MenuCategory) error {
	panic("unimplemented")
}

// UpdateMenuGroup implements MenuService.
func (m *menuService) UpdateMenuGroup(group *hospitalityMenu_models.MenuGroup) error {
	panic("unimplemented")
}

// UpdateMenuItem implements MenuService.
func (m *menuService) UpdateMenuItem(item *hospitalityMenu_models.MenuItem) error {
	panic("unimplemented")
}

// UpdateSubCategory implements MenuService.
func (m *menuService) UpdateSubCategory(subCat *hospitalityMenu_models.MenuSubCategory) error {
	panic("unimplemented")
}

// UpdateVariant implements MenuService.
func (m *menuService) UpdateVariant(variant *hospitalityMenu_models.MenuVariant) error {
	panic("unimplemented")
}

func NewHospitalityMenuService(menuRepo menurepo.MenuRepositoryInterface) *menuService {
	return &menuService{
		menuRepo: menuRepo,
	}

}
