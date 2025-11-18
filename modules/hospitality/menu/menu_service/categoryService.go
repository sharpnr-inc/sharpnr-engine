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

func (s *menuService) GetCategoriesByGroup(ctx context.Context, orgID uuid.UUID, groupID uuid.UUID) ([]hospitalitydtos.Menu_Category_RESPONSE, error) {

	// Validate the menu group exists and belongs to the same organization
	group, err := s.menuRepo.GetMenuGroupByID(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("menu group not found: %w", err)
	}

	if group.OrganizationID != orgID {
		return nil, fmt.Errorf("this menu group does not belong to your organization")
	}

	categories, err := s.menuRepo.GetCategoriesByGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}

	// Map models → DTO
	resp := make([]hospitalitydtos.Menu_Category_RESPONSE, 0)

	for _, c := range categories {
		cat := hospitalitydtos.Menu_Category_RESPONSE{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		}

		// Subcategories
		for _, sub := range c.SubCategories {
			subDto := hospitalitydtos.Menu_SubCategory_RESPONSE{
				Id:          sub.ID,
				Name:        sub.Name,
				Description: sub.Description,
			}

			// Menu Items
			for _, item := range sub.MenuItems {
				subDto.MenuItems = append(subDto.MenuItems, hospitalitydtos.Menu_Item_RESPONSE{
					Id:          item.ID,
					Name:        item.Name,
					Description: item.Description,
					Price:       float32(item.BasePrice),
				})
			}

			cat.SubCategories = append(cat.SubCategories, subDto)
		}

		resp = append(resp, cat)
	}

	return resp, nil
}
