package menuservice

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalityMenu_models "sharpnr-engine/modules/hospitality/menu/models"

	"github.com/google/uuid"
)

func (s *menuService) CreateVariantGroup(
	ctx context.Context,
	menuItemId uuid.UUID,
	req *hospitalitydtos.Menu_CreateVariantGroupAndVariant_REQ,
) error {

	// 1️⃣ Create Variant Group Model
	group := &hospitalityMenu_models.MenuVariantGroup{
		ID:         uuid.New(),
		MenuItemID: menuItemId,
		Name:       req.VariantName,
	}

	// 2️⃣ Save Variant Group
	if err := s.menuRepo.CreateVariantGroup(ctx, group); err != nil {
		return err
	}

	// Now group.ID contains the newly created VariantGroupID

	// 3️⃣ Create each variant
	for _, v := range req.Variant {
		variant := &hospitalityMenu_models.MenuVariant{
			ID:         uuid.New(),
			MenuItemID: menuItemId,
			Name:       v.Name,
			Price:      v.Price,
		}

		if err := s.menuRepo.CreateVariant(ctx, variant); err != nil {
			return err
		}
	}

	return nil
}

func (m *menuService) GetVariantGroupByMenuId(
	ctx context.Context,
	orgid, menuItemId uuid.UUID,
) ([]hospitalitydtos.Menu_VariantGroup_RESPONSE, error) {

	// Validate org
	menuItem, err := m.menuRepo.GetMenuItemByID(ctx, menuItemId)
	if err != nil {
		return nil, err
	}

	if menuItem.OrganizationID != orgid {
		return nil, errors.New("organization mismatch")
	}

	// Fetch all groups
	// groups, err := m.menuRepo.GetVariantGroupsByMenuItemID(ctx, menuItemId)
	// if err != nil {
	// 	return nil, err
	// }

	// Create response list
	resp := make([]hospitalitydtos.Menu_VariantGroup_RESPONSE, 0)

	// for _, g := range groups {
	// 	grpDto := hospitalitydtos.Menu_VariantGroup_RESPONSE{
	// 		Id:            g.ID,
	// 		Name:          g.Name,
	// 		IsMultiSelect: g.IsMultiSelect,
	// 		Variants:      []hospitalitydtos.Menu_Variant_RESPONSE{},
	// 	}

	// 	// Add variants
	// 	for _, v := range g.Variants {
	// 		grpDto.Variants = append(grpDto.Variants, hospitalitydtos.Menu_Variant_RESPONSE{
	// 			Id:    v.ID,
	// 			Name:  v.Name,
	// 			Price: v.Price,
	// 		})
	// 	}

	// 	resp = append(resp, grpDto)
	// }

	return resp, nil
}
