package hospitality_orchestration_service

import (
	"context"

	"github.com/google/uuid"
)

func (s *hospitalityOrchestrationService) ListAllOrgsHotels(ctx context.Context, orgID, emplID uuid.UUID) (string, error) {

	// // First we will list all the Properties of the Organization
	// properties, err := s.propertyService.GetPropertiesOfOrg(ctx, orgID)
	// if err != nil {
	// 	return "bye", err
	// }

	// // Now Based on this Property we will Fetch the Hotels
	// for _, p := range properties.Properties {
	// 	// hotel, err := s.hotelService.GetHotelByPropertyID(ctx, p.ID)
	// }

	// // for hotels we need to find the legal Entity details

	

	return "Hi", nil
}
