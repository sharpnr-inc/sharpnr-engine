package accounting_sero_event_handler

import (
	"context"
	"fmt"
	restaurant_events "sharpnr-engine/business_ops/hospitality/restaurant/events"
	"sharpnr-engine/sero"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"
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

		if err := data.ValidateDto(); err != nil {
			return err
		}

		return nil
	}
}
