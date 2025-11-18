package menuservice

import (
	"context"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (s *menuService) CreateCategory(ctx context.Context, req *hospitalitydtos.Menu_CreateCategoryINMenuGroup_REQ, orgID uuid.UUID) error {

	groupUUID, err := uuid.Parse(req.MenuGroupId)
	if err != nil {
		return fmt.Errorf("invalid menu group id: %w", err)
	}

	// ✔ Ensure this group belongs to SAME ORG
	group, err := s.menuRepo.GetMenuGroupByID(ctx, groupUUID)
	if err != nil {
		return fmt.Errorf("menu group not found: %w", err)
	}

	if group.OrganizationID != orgID {
		return fmt.Errorf("this menu group does not belong to your organization")
	}

	category := &hospitalitymodels.MenuCategory{
		ID:             uuid.New(),
		GroupID:        groupUUID,
		OrganizationID: orgID,
		Name:           req.CategoryName,
		Description:    req.CategoryDescription,
	}

	if err := s.menuRepo.CreateCategory(ctx, category); err != nil {
		return fmt.Errorf("failed to create category: %w", err)
	}

	return nil
}
