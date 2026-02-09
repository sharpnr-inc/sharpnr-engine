package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(ctx context.Context, req *hospitality_menu_contracts.CreateMenuInput) (menuID uuid.UUID, err error)
	// AddCategoryInMenu(ctx context.Context)
}
