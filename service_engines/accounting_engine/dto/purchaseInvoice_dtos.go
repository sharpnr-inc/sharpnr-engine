package accountingdtos

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PurchaseInvoiceRequestDto struct {
	PayableAccountID uuid.UUID
	SupplierID       uuid.UUID
	PaymentStatus    string
	Narration        string
	Lines            []PurchaseInvoiceLineRequestDto
	AmountPaid       *float64
	PurchaseOrderID  *uuid.UUID
	InvoiceNumber    string
	Reference        string
	InvoiceDate      time.Time
	DueDate          *time.Time
}

type PurchaseInvoiceLineRequestDto struct {
	Description string
	Qty         float64
	UnitPrice   float64
	Tax         float64

	// ↑ COGS account OR Inventory Asset account OR Expense account
	// whoever calls the API decides which account to debit
	DebitAccountID uuid.UUID

	// IF this available then Added in the inventory
	InventoryItemID             *uuid.UUID
	InventoryID                 *uuid.UUID
	WareHouseInventoryAccountID *uuid.UUID
}

// Subtotal calculates the amount before tax
func (pi *PurchaseInvoiceLineRequestDto) Subtotal() float64 {
	return pi.Qty * pi.UnitPrice
}

// Total calculates the final amount including tax
func (pi *PurchaseInvoiceLineRequestDto) Total() float64 {
	return pi.Subtotal() + pi.Tax
}

// IsInventoryTracked checks if the line item involves inventory management
func (pi *PurchaseInvoiceLineRequestDto) IsInventoryTracked() bool {
	return pi.InventoryItemID != nil && pi.InventoryID != nil
}

func (pi *PurchaseInvoiceLineRequestDto) Validate() error {
	if pi.Qty <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if pi.UnitPrice < 0 {
		return errors.New("unit price cannot be negative")
	}
	if pi.DebitAccountID == uuid.Nil {
		return errors.New("debit account ID is required")
	}
	return nil
}

func (pi *PurchaseInvoiceLineRequestDto) GetLineTotals() (subtotal float64, tax float64, total float64) {
	subtotal = pi.Qty * pi.UnitPrice
	tax = pi.Tax
	total = subtotal + tax
	return
}

// HasPurchaseOrder returns true if a valid, non-nil PurchaseOrderID is present
func (p *PurchaseInvoiceRequestDto) HasPurchaseOrder() bool {
	if p.PurchaseOrderID == nil {
		return false
	}
	// Check if the underlying UUID is the "Nil" UUID
	return *p.PurchaseOrderID != uuid.Nil
}
