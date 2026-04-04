package inventory_engine_sero_handlers

import (
	"context"
	"sharpnr-engine/sero"
	inventory_engine_services "sharpnr-engine/service_engines/inventory_engine/services"
)

func handlePurchaseInvoiceCreationEvent(
	inventoryEngineService inventory_engine_services.InventoryEngineServicesInterface,
) sero.Handler {
	return func(ctx context.Context, payload interface{}) error {
		return nil
	}
}
