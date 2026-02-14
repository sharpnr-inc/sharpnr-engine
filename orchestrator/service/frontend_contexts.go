package orchestrator_service

import (
	"context"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"

	"github.com/google/uuid"
)

func (s *orchestratorService) GetPropertiesAndRestaurant(
	ctx context.Context,
	orgID uuid.UUID,
	empID uuid.UUID,
) (*orchestrator_dtos.PropertyOfResturants, error) {

	// 1. Fetch all properties for the Org
	prop, err := s.orgPropertyService.GetPropertiesOfOrg(ctx, orgID)
	if err != nil {
		return nil, err
	}

	// 2. Initialize the response container
	result := &orchestrator_dtos.PropertyOfResturants{
		Properties: []orchestrator_dtos.PropertiesOfOrg{},
	}

	// 3. Iterate and Filter
	for _, p := range prop.Properties {
		// Fetch restaurants for this specific property
		resList, err := s.hospitalityDomainSvc.RestaurantService().GetAllRestaurantOfProperty(ctx, p.ID)
		if err != nil {
			return nil, err
		}

		// Only "push" if the property has at least one restaurant
		if len(resList) > 0 {
			restaurantIDs := make([]uuid.UUID, 0, len(resList))
			for _, r := range resList {
				restaurantIDs = append(restaurantIDs, r.ID) // Extracting IDs
			}

			result.Properties = append(result.Properties, orchestrator_dtos.PropertiesOfOrg{
				PropertyID:   p.ID,
				PropertyName: p.Name,
				Restaurents:  restaurantIDs,
			})
		}
	}

	return result, nil
}
