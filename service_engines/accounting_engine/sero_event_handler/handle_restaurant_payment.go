package accounting_sero_event_handler

import (
	"context"
	"fmt"
	restaurant_events "sharpnr-engine/modules/hospitality/restaurant/events"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/service_engines/accounting_engine/services"
)

func handleRestaurantPaymentSettlment(
	accountingService accounting_service.AccountingServiceInterface,
) sero.Handler {

	return func(ctx context.Context, payload interface{}) error {

		data, ok := payload.(restaurant_events.RestaurantPaymentSettledDto)
		if !ok {
			if err := sero.MapToStruct(payload, &data); err != nil {
				return fmt.Errorf("invalid payload type: %w", err)
			}
		}

		return nil
	}
}
