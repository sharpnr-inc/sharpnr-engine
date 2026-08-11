package hospitality_orchestration_service

import (
	"context"
	hotel_contracts "sharpnr-engine/business_ops/hospitality/hotel/contracts"

	"github.com/google/uuid"
)

func (s *hospitalityOrchestrationService) ListAllOrgsHotels(ctx context.Context, orgID, emplID uuid.UUID) ([]hotel_contracts.Hotels, error) {

	return s.hotelService.GetHotelsByOrgID(ctx, orgID, emplID)
}

func (s *hospitalityOrchestrationService) ListAllInfra(ctx context.Context, orgID, empplID uuid.UUID) error {
	return nil
}
