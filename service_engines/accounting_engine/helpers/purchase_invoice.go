package accounting_engine_helpers

import (
	"fmt"
	org_suppliers_models "sharpnr-engine/core/org_supplier/models"

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
