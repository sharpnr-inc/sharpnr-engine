package propertyHospitalityOrchestrator

import (
	bar_contracts "sharpnr-engine/property_business_ops/hospitality/bar/contracts"
	hotel_contracts "sharpnr-engine/property_business_ops/hospitality/hotel/contracts"
	kitchen_contracts "sharpnr-engine/property_infra/hospitality_infra/kitchen/contracts"
	restaurant_contracts "sharpnr-engine/property_business_ops/hospitality/restaurant/contracts"
	hospitality_utilities "sharpnr-engine/property_business_ops/hospitality/utilities"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
)

// Engine contract -> wire DTO mapping (rule 3 in orchestrator/service/doc.go).
// ModuleType is set from the originating service (known & correct).

func mapBarToBusinessModule(b *bar_contracts.BarDetailsReader) orchestrator_dtos.HospitalityBusinessModule {
	return orchestrator_dtos.HospitalityBusinessModule{
		ID:         b.ID,
		Name:       b.Name,
		ModuleType: hospitality_utilities.Bar,
		// Status:       b.Status,
		// IsConfigured: b.IsConfigured,
	}
}

func mapHotelToBusinessModule(h hotel_contracts.HotelReader) orchestrator_dtos.HospitalityBusinessModule {
	return orchestrator_dtos.HospitalityBusinessModule{
		ID:           h.ID,
		Name:         h.Name,
		ModuleType:   hospitality_utilities.Hotel,
		Status:       h.Status,
		IsConfigured: h.IsConfigured,
	}
}

func mapRestaurantToBusinessModule(r *restaurant_contracts.RestaurantDetails) orchestrator_dtos.HospitalityBusinessModule {
	return orchestrator_dtos.HospitalityBusinessModule{
		ID:           r.ID,
		Name:         r.Name,
		ModuleType:   hospitality_utilities.Restaurant,
		Status:       r.Status,
		IsConfigured: r.IsConfigured,
	}
}

// mapKitchenToInfraModule maps a kitchen read-model to the infra DTO.
func mapKitchenToInfraModule(k *kitchen_contracts.Kitchen_Reader) orchestrator_dtos.HospitalityInfraModule {
	return orchestrator_dtos.HospitalityInfraModule{
		ID:        k.ID,
		Name:      k.Name,
		Status:    k.Status,
		InfraType: hospitality_utilities.Kitchen,
		Meta:      nil, // shape TBD
	}
}
