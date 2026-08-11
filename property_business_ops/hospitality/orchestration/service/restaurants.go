package hospitality_orchestration_service

import (
	"context"
	restaurant_contracts "sharpnr-engine/property_business_ops/hospitality/restaurant/contracts"

	"github.com/google/uuid"
)

func (s *hospitalityOrchestrationService) ListAllRestaurant(ctx context.Context, orgID, emplID uuid.UUID) ([]restaurant_contracts.RestaurantDetails, error) {

	return s.restaurantService.GetRestaurantsOfOrg(ctx, orgID)

}
