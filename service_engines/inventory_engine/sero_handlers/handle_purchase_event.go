package inventory_engine_sero_handlers

import (
	"context"
	"fmt"
	"sharpnr-engine/sero"
	accounting_engine_events "sharpnr-engine/service_engines/accounting_engine/events"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
)

func handlePurchaseInvoiceCreationEvent(
	inventoryEngineService inventory_engine_services.InventoryEngineServicesInterface,
) sero.Handler {
	return func(ctx context.Context, payload interface{}) error {

		parsedPayload, ok := payload.(accounting_engine_events.PurchaceInvoiceItemTransterToInventoryEventPayload)
		if !ok {
			if err := sero.MapToStruct(payload, &parsedPayload); err != nil {
				return fmt.Errorf("invalid payload type: %w", err)
			}
		}

		// for i, pp := range parsedPayload.Items{

		// }

		return nil
	}
}
