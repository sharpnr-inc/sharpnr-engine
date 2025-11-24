package barservices

import (
	"context"
	"errors"
	"fmt"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

func (s *barServices) GetBarLayout(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetBarDetails_RESPONSE, error) {
	bar, err := s.barRepo.GetBarByOrgId(ctx, orgId)
	if err != nil {
		return nil, err
	}
	if bar == nil {
		return nil, errors.New("bar not found")
	}

	response := &hospitalitydtos.GetBarDetails_RESPONSE{
		BarId:   bar.ID.String(),
		BarName: bar.DisplayName,
	}

	// --------------------------------
	// ALWAYS LOAD FLOORS
	// --------------------------------
	floors := make([]hospitalitydtos.BarFloorResponse, 0, len(bar.Floors))

	for _, floor := range bar.Floors {
		floorRes := hospitalitydtos.BarFloorResponse{
			FloorId:   floor.ID,
			FloorName: floor.Name,
			Tables:    make([]hospitalitydtos.BarTablesResponse, 0, len(floor.Tables)),
		}

		for _, table := range floor.Tables {
			floorRes.Tables = append(floorRes.Tables, hospitalitydtos.BarTablesResponse{
				TableId:     table.ID,
				TableName:   table.Name,
				TableStatus: table.Status,
			})
		}

		floors = append(floors, floorRes)
	}

	response.Floors = floors

	return response, nil
}

func (s *barServices) UpdateBar(
	ctx context.Context,
	orgId uuid.UUID,
	req hospitalitydtos.BAR_BarUpdateReq,
) error {

	// Step 1: Fetch bar by OrganizationID
	bar, err := s.barRepo.GetBarByOrgId(ctx, orgId)
	if err != nil {
		return errors.New("bar not found for this organization")
	}

	// Step 2: Build dynamic update map
	updates := map[string]interface{}{}

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.DisplayName != nil {
		updates["display_name"] = *req.DisplayName
	}
	if req.LegalName != nil {
		updates["legal_name"] = *req.LegalName
	}
	if req.RegisteredAddress != nil {
		updates["registered_address"] = *req.RegisteredAddress
	}

	if req.OpeningTime != nil {
		updates["opening_time"] = *req.OpeningTime
	}
	if req.ClosingTime != nil {
		updates["closing_time"] = *req.ClosingTime
	}

	if req.IsKitchen != nil {
		updates["is_kitchen"] = *req.IsKitchen
	}
	if req.KitchenID != nil {
		updates["kitchen_id"] = *req.KitchenID
	}

	if req.IsShareLegalDetailsWithMainOrg != nil {
		updates["is_share_legal_details_with_main_org"] = *req.IsShareLegalDetailsWithMainOrg
	}
	if req.GSTNumber != nil {
		updates["gst_number"] = *req.GSTNumber
	}
	if req.CINNumber != nil {
		updates["cin_number"] = *req.CINNumber
	}
	if req.FSSAINumber != nil {
		updates["fssai_number"] = *req.FSSAINumber
	}

	if req.IsMultifloor != nil {
		updates["is_multifloor"] = *req.IsMultifloor
	}
	if req.AddressLine != nil {
		updates["address_line"] = *req.AddressLine
	}
	if req.City != nil {
		updates["city"] = *req.City
	}
	if req.State != nil {
		updates["state"] = *req.State
	}
	if req.Country != nil {
		updates["country"] = *req.Country
	}
	if req.Pincode != nil {
		updates["pincode"] = *req.Pincode
	}

	// Step 3: If nothing to update, don't hit DB
	if len(updates) == 0 {
		return nil
	}

	// Step 4: Update the bar
	err = s.barRepo.UpdateBarDetails(ctx, bar.ID, updates)
	if err != nil {
		return fmt.Errorf("failed to update bar: %w", err)
	}

	return nil
}
