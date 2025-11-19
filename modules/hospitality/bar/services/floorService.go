package barservices

import (
	"context"
	"errors"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (s *barServices) CreateFloor(ctx context.Context, orgId uuid.UUID, req hospitalitydtos.BAR_Add_Floor_Req) error {

	// 1️⃣ Validate bar_id
	if req.BarId == uuid.Nil {
		return errors.New("bar_id is required")
	}

	// 2️⃣ Check if bar exists for this org
	bar, err := s.barRepo.GetBarByOrgId(ctx, orgId)
	if err != nil {
		return err
	}
	if bar == nil {
		return errors.New("bar not found")
	}

	// 3️⃣ Ensure bar matches provided BarId
	if bar.ID != req.BarId {
		return errors.New("bar does not belong to this organization")
	}

	// 4️⃣ Create new floor object
	newFloor := &hospitalitymodels.BarFloor{
		ID:             uuid.New(),
		BarID:          bar.ID,
		OrganizationID: orgId,
		Name:           req.FloorName,
		FloorNumber:    len(bar.Floors) + 1, // auto increment (optional)
	}

	// 5️⃣ Save to DB
	if err := s.barRepo.CreateFloor(ctx, newFloor); err != nil {
		return err
	}

	return nil
}
