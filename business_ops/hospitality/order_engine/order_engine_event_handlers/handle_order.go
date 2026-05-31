package orderEngineSero_event_handlers

import (
	"context"
	"fmt"
	hospitality_order_engine_service "sharpnr-engine/business_ops/hospitality/order_engine/service"
	restaurant_events "sharpnr-engine/business_ops/hospitality/restaurant/events"
	"sharpnr-engine/sero"
)

func handleOrder(
	serivce hospitality_order_engine_service.HospitalityOrderEngineServiceInterface,

) sero.Handler {

	return func(ctx context.Context, payload interface{}) error {

		// Here we force the type safety
		data, ok := payload.(restaurant_events.OrderCreatedPayload)
		if !ok {
			// If it came from the DB, it's likely a map[string]interface{}
			// We use a helper to convert it
			if err := sero.MapToStruct(payload, &data); err != nil {
				return fmt.Errorf("invalid payload type: %w", err)
			}
		}

		return serivce.ProcessOrder(ctx, data)

	}

}
