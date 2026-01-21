package menuservice

import (
	"context"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

// CreateMenuItem implements MenuService.
func (s *menuService) CreateMenuItem(
	ctx context.Context,
	req *hospitalitydtos.Menu_CreateMenuItem_REQ,
	orgID uuid.UUID,
) (*hospitalitydtos.Menu_CreateMenuItem_RES, error) {

	subID, err := uuid.Parse(req.SubCategoryID)
	if err != nil {
		return nil, fmt.Errorf("invalid subcategory id: %w", err)
	}

	// Validate Subcategory exists + belongs to org
	sub, err := s.menuRepo.GetSubCategoryByID(ctx, subID)
	if err != nil {
		return nil, fmt.Errorf("subcategory not found: %w", err)
	}

	if sub.OrganizationID != orgID {
		return nil, fmt.Errorf("subcategory does not belong to your organization")
	}

	item := &hospitalityMenu_models.MenuItem{
		ID:             uuid.New(),
		OrganizationID: orgID,
		Name:           req.Name,
		Description:    req.Description,
		IsVeg:          req.IsVeg,
		BasePrice:      req.BasePrice,
	}

	if err := s.menuRepo.CreateMenuItem(ctx, item); err != nil {
		return nil, fmt.Errorf("failed to create menu item: %w", err)
	}

	res := hospitalitydtos.Menu_CreateMenuItem_RES{
		ID:          item.ID.String(),
		CategoryID:  item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		IsVeg:       item.IsVeg,
		BasePrice:   item.BasePrice,
	}

	return &res, nil
}
