package hospitality_orchestration_service

// This is the Hospitality Orchestration Layer
// This orchestration layer includes the all modules available in Hospitality Domain
// As all Domain exists on the Property it have property and legal Entity Service

import (
	legal_entity_service "sharpnr-engine/core/legal_entity/service"
	org_property_services "sharpnr-engine/core/org_property/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"
)

type hospitalityOrchestrationService struct {
	propertyService    org_property_services.OrganizationPropertyServicesInterfaces
	hotelService       hotelservices.HotelServiceInterface
	restaurantService  restaurantservices.RestaurantServiceInterface
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface
	kitchenService     kitchenservices.KitchenServiceInterface
}

func NewHospitalityOrchestrationService(
	propertyService org_property_services.OrganizationPropertyServicesInterfaces,
	hotelService hotelservices.HotelServiceInterface,
	restaurantService restaurantservices.RestaurantServiceInterface,
	legalEntityService legal_entity_service.OrgLegalEntityServiceInterface,
	kitchenService kitchenservices.KitchenServiceInterface,
) *hospitalityOrchestrationService {
	return &hospitalityOrchestrationService{
		propertyService:    propertyService,
		hotelService:       hotelService,
		restaurantService:  restaurantService,
		legalEntityService: legalEntityService,
		kitchenService:     kitchenService,
	}
}
