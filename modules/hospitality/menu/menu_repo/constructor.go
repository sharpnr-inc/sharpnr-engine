package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func (r *menuRepo) CreateVariant(ctx context.Context, variant *hospitalitymodels.MenuVariant) error {
	return r.db.WithContext(ctx).Create(&variant).Error
}

// DeleteCategory implements MenuRepository.
func (m *menuRepo) DeleteCategory(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteMenuGroup implements MenuRepository.
func (m *menuRepo) DeleteMenuGroup(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteMenuItem implements MenuRepository.
func (m *menuRepo) DeleteMenuItem(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteSubCategory implements MenuRepository.
func (m *menuRepo) DeleteSubCategory(id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteVariant implements MenuRepository.
func (m *menuRepo) DeleteVariant(id uuid.UUID) error {
	panic("unimplemented")
}

// GetItemsBySubCategory implements MenuRepository.
func (m *menuRepo) GetItemsBySubCategory(subCatID uuid.UUID) ([]hospitalitymodels.MenuItem, error) {
	panic("unimplemented")
}

// GetSubCategoriesByCategory implements MenuRepository.
func (m *menuRepo) GetSubCategoriesByCategory(categoryID uuid.UUID) ([]hospitalitymodels.MenuSubCategory, error) {
	panic("unimplemented")
}

// GetVariantsByItem implements MenuRepository.
func (m *menuRepo) GetVariantsByItem(itemID uuid.UUID) ([]hospitalitymodels.MenuVariant, error) {
	panic("unimplemented")
}

// UpdateCategory implements MenuRepository.
func (m *menuRepo) UpdateCategory(category *hospitalitymodels.MenuCategory) error {
	panic("unimplemented")
}

// UpdateMenuItem implements MenuRepository.
func (m *menuRepo) UpdateMenuItem(item *hospitalitymodels.MenuItem) error {
	panic("unimplemented")
}

// UpdateSubCategory implements MenuRepository.
func (m *menuRepo) UpdateSubCategory(subCat *hospitalitymodels.MenuSubCategory) error {
	panic("unimplemented")
}

// UpdateVariant implements MenuRepository.
func (m *menuRepo) UpdateVariant(variant *hospitalitymodels.MenuVariant) error {
	panic("unimplemented")
}

func NewHospitalityMenuRepository(db *gorm.DB) *menuRepo {
	return &menuRepo{
		db: db,
	}
}
