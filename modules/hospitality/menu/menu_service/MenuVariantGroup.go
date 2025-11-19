package menuservice

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (s *menuService) CreateVariantGroup(
	ctx context.Context,
	menuItemId uuid.UUID,
	req *hospitalitydtos.Menu_CreateVariantGroupAndVariant_REQ,
) error {

	// 1️⃣ Create Variant Group Model
	group := &hospitalitymodels.MenuVariantGroup{
		ID:         uuid.New(),
		MenuItemID: menuItemId,
		Name:       req.VariantName,
	}

	// 2️⃣ Save Variant Group
	if err := s.menuRepo.CreateVariantGroup(ctx, group); err != nil {
		return err
	}

	// Now group.ID contains the newly created VariantGroupID
	var groupID = group.ID

	// 3️⃣ Create each variant
	for _, v := range req.Variant {
		variant := &hospitalitymodels.MenuVariant{
			ID:         uuid.New(),
			GroupID:    groupID,
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
