package menurepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (r *menuRepo) CreateVariantGroup(ctx context.Context, variantGroup *hospitalitymodels.MenuVariantGroup) error {
	return r.db.WithContext(ctx).Create(&variantGroup).Error
}

func (r *menuRepo) UpdateVariantGroup(ctx context.Context, variant *hospitalitymodels.MenuVariantGroup) error {
	panic("Unimplmented")
}
