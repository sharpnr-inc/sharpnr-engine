package hospitality_orchestration_service

import (
	"context"
	hotel_contracts "sharpnr-engine/business_ops/hospitality/hotel/contracts"
	restaurant_contracts "sharpnr-engine/business_ops/hospitality/restaurant/contracts"

	"github.com/google/uuid"
)

type OrchestrationServiceInterface interface {
	hospitalityOrchestrationHotelInterface
}

type hospitalityOrchestrationHotelInterface interface {
	ListAllOrgsHotels(ctx context.Context, orgID, emplID uuid.UUID) ([]hotel_contracts.Hotels, error)
	ListAllRestaurant(ctx context.Context, orgID, emplID uuid.UUID) ([]restaurant_contracts.RestaurantDetails, error)
	ListAllBar(ctx context.Context, orgID, emplID uuid.UUID) error
	ListAllInfra(ctx context.Context, orgID, empplID uuid.UUID) error
}
