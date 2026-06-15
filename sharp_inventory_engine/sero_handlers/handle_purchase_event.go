package inventory_engine_sero_handlers

import (
	"context"
	"fmt"
	"sharpnr-engine/sero"
	accounting_engine_events "sharpnr-engine/sharp_accounting_engine/accounting_engine/events"
	inventory_engine_contracts "sharpnr-engine/sharp_inventory_engine/contracts"
	inventory_engine_services "sharpnr-engine/sharp_inventory_engine/services"
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

		input := mapToAddItemsInInventoryInput(parsedPayload)
		if len(input.Lines) == 0 {
			// No inventory-tracked lines in this invoice — nothing to do
			return nil
		}

		if err := inventoryEngineService.AddItemsInInventory(ctx, input); err != nil {
			return fmt.Errorf("failed to process inventory from purchase invoice: %w", err)
		}

		return nil
	}
}

// mapToAddItemsInInventoryInput isolates the accounting → inventory mapping.
// Any future changes to the accounting DTO only touch this function.
func mapToAddItemsInInventoryInput(
	payload accounting_engine_events.PurchaceInvoiceItemTransterToInventoryEventPayload,
) inventory_engine_contracts.AddItemsInInventoryInput {

	var lines []inventory_engine_contracts.AddItemsInInventoryLineInput

	for _, line := range payload.Items {
		// Skip non-inventory-tracked lines
		if line.InventoryID == nil || line.WareHouseInventoryAccountID == nil {
			continue
		}

		lines = append(lines, inventory_engine_contracts.AddItemsInInventoryLineInput{
			InventoryItemID:    line.InventoryItemID,
			InventoryID:        *line.InventoryID,
			InventoryAccountID: *line.WareHouseInventoryAccountID,
			Description:        line.Description,
			Unit:               line.Unit,
			Qty:                line.Qty,
			UnitPrice:          line.UnitPrice,
			COGSAccountID:      &line.DebitAccountID,
		})
	}

	return inventory_engine_contracts.AddItemsInInventoryInput{
		OrganizationID: payload.OrganizationID,
		EmployeeID:     payload.EmployeeID,
		SupplierID:     payload.SupplierID,
		Lines:          lines,
	}
}
