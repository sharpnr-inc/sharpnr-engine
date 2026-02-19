package orderEngineSero_event_handlers

import (
	hospitality_order_engine_service "sharpnr-engine/modules/hospitality/order_engine/service"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	"sharpnr-engine/sero"
)

func Register(
	seroEngine *sero.SEROEngine,
	service hospitality_order_engine_service.HospitalityOrderEngineServiceInterface,

) {

	seroEngine.Subscribe(
		restaurant_events.OrderCreatedInRestaurant,
		handleOrder(service),
		true,
	)
}
