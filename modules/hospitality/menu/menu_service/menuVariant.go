package menuservice

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// GetVariantsByItem implements MenuService.
// this return the Group
func (m *menuService) GetVariantsByItem(ctx context.Context, itemID uuid.UUID) ([]hospitalitymodels.MenuVariant, error) {
	panic("unimplemented")
}
