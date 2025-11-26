package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (r *menuRepo) CreateVariantGroup(ctx context.Context, variantGroup *hospitalitymodels.MenuVariantGroup) error {
	return r.db.WithContext(ctx).Create(&variantGroup).Error
}

func (r *menuRepo) UpdateVariantGroup(ctx context.Context, variant *hospitalitymodels.MenuVariantGroup) error {
	panic("Unimplmented")
}

func (r *menuRepo) GetVariantGroupsByMenuItemID(ctx context.Context, menuId uuid.UUID) ([]hospitalitymodels.MenuVariantGroup, error) {
	var groups []hospitalitymodels.MenuVariantGroup

	err := r.db.WithContext(ctx).
		Preload("Variants").
		Where("menu_item_id = ?", menuId).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}
