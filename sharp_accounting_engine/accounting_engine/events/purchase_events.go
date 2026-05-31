package accounting_engine_events

import (
	"sharpnr-engine/sero"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"

	"github.com/google/uuid"
)

const (
	PurchaseInvoiceCreated sero.SeroEngineEventType = "PURCHASE_INVOICE_CREATED"
)

type PurchaceInvoiceItemTransterToInventoryEventPayload struct {
	OrganizationID uuid.UUID
	EmployeeID     uuid.UUID
	SupplierID     uuid.UUID
	Items          []accountingdtos.PurchaseInvoiceLineRequestDto
}
