package hospitality_orchestration_service

import (
	org_property_services "sharpnr-engine/core/org_property/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"
)

type hospitalityOrchestrationService struct {
	propertyService   org_property_services.OrganizationPropertyServicesInterfaces
	hotelService      hotelservices.HotelServiceInterface
	restaurantService restaurantservices.RestaurantServiceInterface
}

func NewHospitalityOrchestrationService(propertyService org_property_services.OrganizationPropertyServicesInterfaces, hotelService hotelservices.HotelServiceInterface, restaurantService restaurantservices.RestaurantServiceInterface) *hospitalityOrchestrationService {
	return &hospitalityOrchestrationService{
		propertyService:   propertyService,
		hotelService:      hotelService,
		restaurantService: restaurantService,
	}
}
