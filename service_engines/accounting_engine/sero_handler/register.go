package accounting_sero_event_handler

import (
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
)

func Register(
	seroEngine *sero.SEROEngine,
	accountingEngine accounting_service.AccountingServiceInterface,
) {

	seroEngine.Subscribe(
		restaurant_events.RestaurantPaymentSettled,
		handleRestaurantPaymentSettlment(accountingEngine),
		false,
	)
}
