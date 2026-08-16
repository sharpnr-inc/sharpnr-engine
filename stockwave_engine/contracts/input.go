package stockwaveEngineContracts

import "github.com/google/uuid"

// AddItemsInInventoryLineInput represents a single inventory-tracked line
// coming from any upstream event (purchase invoice, GRN, etc.)
type AddItemsInInventoryLineInput struct {
	// Item identity
	InventoryItemID    *uuid.UUID // nil = new item, non-nil = existing item
	InventoryID        uuid.UUID
	InventoryAccountID uuid.UUID // WareHouseInventoryAccountID renamed for clarity

	// Item metadata — only used when creating a new item
	Description string
	Unit        string
	Category    string // optional, caller can pass "" → defaults to "OTHER"

	// Transaction details
	Qty       float64
	UnitPrice float64

	// Accounting linkage
	COGSAccountID *uuid.UUID // DebitAccountID mapped here
}

type AddItemsInInventoryInput struct {
	OrganizationID uuid.UUID
	EmployeeID     uuid.UUID
	SupplierID     uuid.UUID
	Lines          []AddItemsInInventoryLineInput
}
