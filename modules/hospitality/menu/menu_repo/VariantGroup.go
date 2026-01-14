package menurepo

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

func (r *menuRepo) CreateVariantGroup(ctx context.Context, variantGroup *hospitalityMenu_models.MenuVariantGroup) error {
	return r.db.WithContext(ctx).Create(&variantGroup).Error
}

func (r *menuRepo) UpdateVariantGroup(ctx context.Context, variant *hospitalityMenu_models.MenuVariantGroup) error {
	panic("Unimplmented")
}

func (r *menuRepo) GetVariantGroupsByMenuItemID(ctx context.Context, menuId uuid.UUID) ([]hospitalityMenu_models.MenuVariantGroup, error) {
	var groups []hospitalityMenu_models.MenuVariantGroup

	err := r.db.WithContext(ctx).
		Preload("Variants").
		Where("menu_item_id = ?", menuId).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}
