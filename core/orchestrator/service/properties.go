package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/core/orchestrator/dtos"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func (s *orchestratorService) GetAllPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*orchestrator_dtos.PropertiesOfOrganization_ResDto, error) {

	res, err := s.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	readers := make([]orchestrator_dtos.PropertiesReaderSlice, 0, len(res.Properties))

	for _, p := range res.Properties {
		readers = append(readers, orchestrator_dtos.PropertiesReaderSlice{
			ID:     p.ID,
			Name:   p.Name,
			City:   p.City,
			State:  p.State,
			Domain: p.Domain,
			Code:   p.Code,
		})
	}

	return &orchestrator_dtos.PropertiesOfOrganization_ResDto{
		Properties: readers,
	}, nil
}

func (c *orchestratorService) GetPropertyDetail(ctx context.Context, orgID, propertyID uuid.UUID) (*orchestrator_dtos.PropertyDetails_res_dto, error) {

	propDetail, err := c.orgPropertyService.GetPropertyDetailsByID(ctx, orgID, propertyID)
	if err != nil {
		return nil, err
	}

	// Map the base details to your response DTO
	response := &orchestrator_dtos.PropertyDetails_res_dto{
		Name:        propDetail.Name,
		Code:        propDetail.Code,
		City:        propDetail.City,
		AddressLine: propDetail.AddressLine,
		State:       propDetail.State,
		Country:     propDetail.Country,
		Domain:      &propDetail.Domain,
	}

	// Case 1: Empty Domain - return the basic details immediately
	if propDetail.Domain == "" {
		return response, nil
	}

	// Case 2: Hospitality Domain - enrich with sub-modules
	if propDetail.Domain == "HOSPITALITY" {
		err := c.enrichHospitalityData(ctx, orgID, propertyID, response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (c *orchestratorService) enrichHospitalityData(
	ctx context.Context,
	orgID, propertyID uuid.UUID,
	res *orchestrator_dtos.PropertyDetails_res_dto,
) error {
	// 1. Initialize the Modules container
	res.Modules = &orchestrator_dtos.HospitalityPropertyModules{}

	g, ctx := errgroup.WithContext(ctx)

	// --- Hotel Module ---
	g.Go(func() error {
		hotel, err := c.hospitalityDomainSvc.HotelService().GetHotelDetails(ctx, orgID, propertyID)
		if err != nil {
			if isNotFoundError(err) {
				return nil
			}
			return err
		}

		if hotel != nil && len(hotel.Hotels) > 0 {
			hotelDTOs := make([]orchestrator_dtos.HotelDetails_dto, len(hotel.Hotels))

			for i, h := range hotel.Hotels {
				hotelDTOs[i] = orchestrator_dtos.HotelDetails_dto{
					ID:     h.ID,
					Name:   h.Name,
					Status: h.Status,
				}
			}

			// Since the struct expects *[]HotelDetails_dto, we take the address of the slice
			res.Modules.HotelDetails = &hotelDTOs
		}

		return nil
	})

	// --- Future: Add Bar and Restaurant g.Go blocks here ---

	// 2. Wait for all background tasks to finish
	return g.Wait()
}

// Utility to check for 404/Not Found errors
// Adjust this logic to match your specific DB or Error library (e.g., GORM, sql)
func isNotFoundError(err error) bool {
	// Example: return errors.Is(err, gorm.ErrRecordNotFound)
	// Example: return strings.Contains(err.Error(), "not found")
	return err != nil && strings.Contains(strings.ToLower(err.Error()), "not found")
}
