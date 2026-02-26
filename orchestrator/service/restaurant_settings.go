package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

func (s *orchestratorService) GetAllRestaurant(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
) ([]orchestrator_dtos.AllRestaurantsResponseDto, error) {

	results := make([]orchestrator_dtos.AllRestaurantsResponseDto, 0)

	restaurants, err := s.hospitalityDomainSvc.
		RestaurantService().
		GetRestaurantsOfOrg(ctx, orgID)
	if err != nil {
		return results, err
	}

	// optional optimization
	results = make([]orchestrator_dtos.AllRestaurantsResponseDto, 0, len(restaurants))

	for _, r := range restaurants {

		// assuming restaurant has PropertyID
		prop, err := s.orgPropertyService.GetPropertyDetailsByID(
			ctx,
			orgID,
			r.PropertyID,
		)
		if err != nil {
			return results, err
		}

		dto := orchestrator_dtos.AllRestaurantsResponseDto{
			ID:           r.ID,
			Name:         r.Name,
			OrgID:        orgID,
			IsActive:     r.Status,
			IsConfigured: r.IsConfigured,
		}

		// nil-safe mapping
		if prop != nil {
			dto.Property = orchestrator_dtos.RestaurantProperty{
				ID:          r.PropertyID,
				Name:        prop.Name,
				AddressLine: prop.AddressLine,
				City:        prop.City,
				State:       prop.State,
				Country:     prop.Country,
			}
		}

		results = append(results, dto)
	}

	return results, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orchestratorService) GetAllKitchen(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID) ([]orchestrator_dtos.AllKitchenResponseDto, error) {

	results := make([]orchestrator_dtos.AllKitchenResponseDto, 0)
	kitchen, err := s.hospitalityDomainSvc.KitchenService().GetKitchensByOrgID(ctx, orgID)
	if err != nil {
		return results, err
	}

	results = make([]orchestrator_dtos.AllKitchenResponseDto, 0, len(kitchen))

	for _, k := range kitchen {

		prop, err := s.orgPropertyService.GetPropertyDetailsByID(
			ctx,
			orgID,
			k.PropertyID,
		)
		if err != nil {
			return results, err
		}

		dto := orchestrator_dtos.AllKitchenResponseDto{
			ID:   k.ID,
			Name: k.Name,
		}

		// nil-safe mapping
		if prop != nil {
			dto.Property = orchestrator_dtos.RestaurantProperty{
				ID:          k.PropertyID,
				Name:        prop.Name,
				AddressLine: prop.AddressLine,
				City:        prop.City,
				State:       prop.State,
				Country:     prop.Country,
			}
		}

		results = append(results, dto)

	}

	return results, nil

}
