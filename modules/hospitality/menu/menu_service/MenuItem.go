package menuservice

import (
	"context"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// CreateMenuItem implements MenuService.
func (s *menuService) CreateMenuItem(
	ctx context.Context,
	req *hospitalitydtos.Menu_CreateMenuItem_REQ,
	orgID uuid.UUID,
) error {

	subID, err := uuid.Parse(req.SubCategoryID)
	if err != nil {
		return fmt.Errorf("invalid subcategory id: %w", err)
	}

	// Validate Subcategory exists + belongs to org
	sub, err := s.menuRepo.GetSubCategoryByID(ctx, subID)
	if err != nil {
		return fmt.Errorf("subcategory not found: %w", err)
	}

	if sub.OrganizationID != orgID {
		return fmt.Errorf("subcategory does not belong to your organization")
	}

	item := &hospitalitymodels.MenuItem{
		ID:             uuid.New(),
		SubCategoryID:  &subID,
		OrganizationID: orgID,
		Name:           req.Name,
		Description:    req.Description,
		IsVeg:          req.IsVeg,
		BasePrice:      req.BasePrice,
		ServedBy:       req.ServedBy,
	}

	if err := s.menuRepo.CreateMenuItem(ctx, item); err != nil {
		return fmt.Errorf("failed to create menu item: %w", err)
	}

	return nil
}
