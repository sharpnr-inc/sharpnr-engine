package menuservice

import (
	"context"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (s *menuService) CreateSubCategory(
	ctx context.Context,
	req *hospitalitydtos.Menu_CreateSubCategory_REQ,
	orgID uuid.UUID,
) error {

	catID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return fmt.Errorf("invalid category id: %w", err)
	}

	// 1️⃣ Validate category exists and belongs to this org
	category, err := s.menuRepo.GetCategoryByID(ctx, catID)
	if err != nil {
		return fmt.Errorf("category not found: %w", err)
	}

	if category.OrganizationID != orgID {
		return fmt.Errorf("category does not belong to your organization")
	}

	// 2️⃣ Create SubCategory
	sub := &hospitalitymodels.MenuSubCategory{
		ID:             uuid.New(),
		CategoryID:     catID,
		OrganizationID: orgID,
		Name:           req.SubCategoryName,
		Description:    req.SubCategoryDesc,
	}

	if err := s.menuRepo.CreateSubCategory(ctx, sub); err != nil {
		return fmt.Errorf("failed to create subcategory: %w", err)
	}

	return nil
}
