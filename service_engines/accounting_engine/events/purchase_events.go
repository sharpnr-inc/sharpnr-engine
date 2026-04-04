package accounting_engine_events

import (
	"sharpnr-engine/sero"

	"github.com/google/uuid"
)

const (
	PurchaseInvoiceCreated sero.SeroEngineEventType = "PURCHASE_INVOICE_CREATED"
)

type PurchaceInvoiceItemTransterToInventoryEventPayload struct {
	InventoryAccountID uuid.UUID
	InventoryItemID    uuid.UUID
	InventoryID        uuid.UUID

	// other to decided
}
