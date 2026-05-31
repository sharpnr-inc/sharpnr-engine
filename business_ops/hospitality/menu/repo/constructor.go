package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/business_ops/hospitality/menu/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func (r *menuRepo) CreateVariant(ctx context.Context, variant *hospitalityMenu_models.MenuVariant) error {
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

// GetVariantsByItem implements MenuRepository.
func (m *menuRepo) GetVariantsByItem(itemID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error) {
	panic("unimplemented")
}

// UpdateCategory implements MenuRepository.
func (m *menuRepo) UpdateCategory(category *hospitalityMenu_models.MenuCategory) error {
	panic("unimplemented")
}

// UpdateMenuItem implements MenuRepository.
func (m *menuRepo) UpdateMenuItem(item *hospitalityMenu_models.MenuItem) error {
	panic("unimplemented")
}

// UpdateSubCategory implements MenuRepository.
func (m *menuRepo) UpdateSubCategory(subCat *hospitalityMenu_models.MenuSubCategory) error {
	panic("unimplemented")
}

// UpdateVariant implements MenuRepository.
func (m *menuRepo) UpdateVariant(variant *hospitalityMenu_models.MenuVariant) error {
	panic("unimplemented")
}

func NewHospitalityMenuRepository(db *gorm.DB) *menuRepo {
	return &menuRepo{
		db: db,
	}
}
