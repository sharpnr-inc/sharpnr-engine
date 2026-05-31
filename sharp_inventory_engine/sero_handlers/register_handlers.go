package inventory_engine_sero_handlers

import (
	"sharpnr-engine/sero"
	accounting_engine_events "sharpnr-engine/sharp_accounting_engine/accounting_engine/events"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"
)

func RegisterSeroHandler(
	seroEngine *sero.SEROEngine,
	inventoryEngineService inventory_engine_services.InventoryEngineServicesInterface,
) {

	seroEngine.Subscribe(
		accounting_engine_events.PurchaseInvoiceCreated,
		handlePurchaseInvoiceCreationEvent(inventoryEngineService),
		false,
	)
}
