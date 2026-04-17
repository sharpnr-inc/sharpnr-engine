package accounting_engine_helpers

import (
	"fmt"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"github.com/google/uuid"
)

// FindSupplierByID searches for a supplier in a slice and returns a pointer to it if found.
func FindSupplierByID(suppliers []org_suppliers_models.Supplier, targetID uuid.UUID) (*org_suppliers_models.Supplier, error) {
	for i := range suppliers {
		if suppliers[i].ID == targetID {
			// Return the address of the element in the slice to avoid copying
			return &suppliers[i], nil
		}
	}
	return nil, fmt.Errorf("supplier with ID %s not found", targetID)
}

func FliterItemToTransferedToInventory(
	purchaseline []accountingdtos.PurchaseInvoiceLineRequestDto,
) (itemToTransfered []accountingdtos.PurchaseInvoiceLineRequestDto, err error) {

	// Pre-allocate the slice capacity to the length of the input.
	// This avoids unnecessary memory reallocations if every line happens to be an inventory item.
	itemToTransfered = make([]accountingdtos.PurchaseInvoiceLineRequestDto, 0, len(purchaseline))

	for i := range purchaseline {
		line := purchaseline[i]

		// 1. Check if InventoryID is present and valid
		hasValidInventoryID := line.InventoryID != nil && *line.InventoryID != uuid.Nil

		// 2. Check if WareHouseInventoryAccountID is present and valid
		hasValidWarehouseID := line.WareHouseInventoryAccountID != nil && *line.WareHouseInventoryAccountID != uuid.Nil

		// If both conditions are met, this item needs to go to inventory
		if hasValidInventoryID && hasValidWarehouseID {
			itemToTransfered = append(itemToTransfered, line)
		}
	}

	return itemToTransfered, nil
}

// ValidatePurchaseInvoiceLines ensures every line has the minimum fields
// required to produce a balanced journal entry.
// Call this AFTER lines are built, BEFORE journaling.
func ValidatePurchaseInvoiceLines(lines []*accounting_models.PurchaseInvoiceLine) error {
	for _, line := range lines {
		if line.DebitAccountID == uuid.Nil {
			return fmt.Errorf(
				"line %q is missing a debit account — all lines must have an account assigned before posting",
				line.Description,
			)
		}
		if line.TotalAmount <= 0 {
			return fmt.Errorf(
				"line %q has a zero or negative total — cannot post invoice",
				line.Description,
			)
		}
	}
	return nil
}
