package propertyHospitalityOrchestrator

import (
	"context"
	restaurant_contracts "sharpnr-engine/property_business_ops/hospitality/restaurant/contracts"

	"github.com/google/uuid"
)

func (s *HospitalityOrchestrator) GetRestaurantDetailsByRestaurantID(ctx context.Context, orgID, employeeID, restaurantID uuid.UUID) (restaurant_contracts.RestaurantGeneralInfo_resp_dto, error) {

	return s.domainSvc.RestaurantService().GetGeneralDetails(ctx, orgID, employeeID, restaurantID)

}

func (s *HospitalityOrchestrator) GetRestaurantFloorLayouts(ctx context.Context, orgID, employeeID, restaurantID uuid.UUID) (restaurant_contracts.RestaurantLayout_Response_dto, error) {
	return s.domainSvc.RestaurantService().GetRestaurantLayout(ctx, orgID, restaurantID)
}
