package stockwaveEngineSeroHandlers

import (
	"sharpnr-engine/sero"
	accounting_engine_events "sharpnr-engine/sharp_accounting_engine/accounting_engine/events"
	stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"
)

func RegisterSeroHandler(
	seroEngine *sero.SEROEngine,
	inventoryEngineService stockwaveEngineServices.InventoryEngineServicesInterface,
) {

	seroEngine.Subscribe(
		accounting_engine_events.PurchaseInvoiceCreated,
		handlePurchaseInvoiceCreationEvent(inventoryEngineService),
		false,
	)
}
