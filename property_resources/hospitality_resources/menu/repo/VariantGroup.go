package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/property_resources/hospitality_resources/menu/models"

	"github.com/google/uuid"
)

func (r *menuRepo) CreateVariantGroup(ctx context.Context, variantGroup *hospitalityMenu_models.MenuVariantGroup) error {
	return r.db.WithContext(ctx).Create(&variantGroup).Error
}

func (r *menuRepo) UpdateVariantGroup(ctx context.Context, variant *hospitalityMenu_models.MenuVariantGroup) error {
	panic("Unimplemented")
}

func (r *menuRepo) GetVariantGroupsByMenuItemID(ctx context.Context, menuId uuid.UUID) ([]hospitalityMenu_models.MenuVariantGroup, error) {
	var groups []hospitalityMenu_models.MenuVariantGroup

	err := r.db.WithContext(ctx).
		Where("menu_item_id = ?", menuId).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *menuRepo) GetVariantsByVariantGroupID(ctx context.Context, variantGroupID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error) {
	var groups []hospitalityMenu_models.MenuVariant

	err := r.db.WithContext(ctx).
		Where("menu_variant_group_id = ?", variantGroupID).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}
