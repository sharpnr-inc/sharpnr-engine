package hospitality_orchestration_service

import (
	"context"

	"github.com/google/uuid"
)

type OrchestrationServiceInterface interface {
	hospitalityOrchestrationHotelInterface
}

type hospitalityOrchestrationHotelInterface interface {
	ListAllOrgsHotels(ctx context.Context, orgID, emplID uuid.UUID) (string, error)
	ListAllRestaurant(ctx context.Context, orgID, emplID uuid.UUID) error
	ListAllBar(ctx context.Context, orgID, emplID uuid.UUID) error
	ListAllInfra(ctx context.Context, orgID, empplID uuid.UUID) error
}
