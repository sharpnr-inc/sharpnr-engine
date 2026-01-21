package menuservice

import (
	"context"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

// GetVariantsByItem implements MenuService.
// this return the Group
func (m *menuService) GetVariantsByItem(ctx context.Context, itemID uuid.UUID) ([]hospitalityMenu_models.MenuVariant, error) {
	panic("unimplemented")
}
