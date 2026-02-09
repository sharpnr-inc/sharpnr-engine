package menuservice

import (
	"context"
	hospitality_menu_contracts "sharpnr-engine/modules/hospitality/menu/contracts"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"
	"time"

	"github.com/google/uuid"
)

func (s *menuService) CreateMenu(ctx context.Context, req *hospitality_menu_contracts.CreateMenuInput) (menuID uuid.UUID, err error) {

	menuGrp := &hospitalityMenu_models.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: req.OrganizationID,
		PropertyID:     req.PropertyID,
		Name:           req.MenuName,
		Description:    *req.MenuDescription,
		OwnerModule:    req.OwnerModuleType,
		OwnerID:        req.OwnerID,
		IsActive:       false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := s.menuRepo.CreateMenuGroup(ctx, menuGrp); err != nil {
		return uuid.Nil, err
	}

	return menuGrp.ID, nil
}
