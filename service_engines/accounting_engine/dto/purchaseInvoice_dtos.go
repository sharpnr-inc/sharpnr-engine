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
	Description   string
	Qty           float64
	Unit          string
	UnitPrice     float64
	Tax           float64    // percentage e.g. 18.0 for 18%
	TaxCategoryID *uuid.UUID // ✅ added — links to tax system, stored on the line model

	// COGS account OR Inventory Asset account OR Expense account
	DebitAccountID uuid.UUID

	// If all three are set the line is inventory-tracked
	InventoryItemID             *uuid.UUID
	InventoryID                 *uuid.UUID
	WareHouseInventoryAccountID *uuid.UUID
}

// Subtotal returns qty * unitPrice before tax
func (pi *PurchaseInvoiceLineRequestDto) Subtotal() float64 {
	return pi.Qty * pi.UnitPrice
}

// TaxAmount returns the tax amount (not the rate)
func (pi *PurchaseInvoiceLineRequestDto) TaxAmount() float64 {
	return pi.Subtotal() * pi.Tax / 100
}

// TotalAmount returns subtotal + tax amount
func (pi *PurchaseInvoiceLineRequestDto) TotalAmount() float64 {
	return pi.Subtotal() + pi.TaxAmount()
}

// IsInventoryTracked returns true if this line should update inventory on POST
func (pi *PurchaseInvoiceLineRequestDto) IsInventoryTracked() bool {
	return pi.InventoryItemID != nil &&
		pi.InventoryID != nil &&
		pi.WareHouseInventoryAccountID != nil
}

// Validate checks the line has the minimum required fields
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

// GetLineTotals returns subtotal, taxAmount, total — all correctly computed
func (pi *PurchaseInvoiceLineRequestDto) GetLineTotals() (subtotal float64, taxAmount float64, total float64) {
	subtotal = pi.Qty * pi.UnitPrice
	taxAmount = subtotal * pi.Tax / 100 // ✅ amount not rate
	total = subtotal + taxAmount        // ✅ correct
	return
}

// HasPurchaseOrder returns true if a valid non-nil PurchaseOrderID is present
func (p *PurchaseInvoiceRequestDto) HasPurchaseOrder() bool {
	if p.PurchaseOrderID == nil {
		return false
	}
	return *p.PurchaseOrderID != uuid.Nil
}
