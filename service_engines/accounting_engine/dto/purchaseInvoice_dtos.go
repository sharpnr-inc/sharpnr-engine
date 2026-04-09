package accountingdtos

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PurchaseInvoiceRequestDto struct {
	PayableAccountID uuid.UUID                       `json:"payable_account_id"`
	SupplierID       uuid.UUID                       `json:"supplier_id"`
	PaymentStatus    string                          `json:"payment_status"`
	Narration        string                          `json:"narration"`
	Lines            []PurchaseInvoiceLineRequestDto `json:"lines"`
	AmountPaid       *float64                        `json:"amount_paid,omitempty"`
	PurchaseOrderID  *uuid.UUID                      `json:"purchase_order_id,omitempty"`
	InvoiceNumber    string                          `json:"invoice_number"`
	Reference        string                          `json:"reference"`
	InvoiceDate      time.Time                       `json:"invoice_date"`
	DueDate          *time.Time                      `json:"due_date,omitempty"`
}

type PurchaseInvoiceLineRequestDto struct {
	Description                 string     `json:"description"`
	Qty                         float64    `json:"qty"`
	Unit                        string     `json:"unit"`
	UnitPrice                   float64    `json:"unit_price"`
	Tax                         float64    `json:"tax"`
	TaxCategoryID               *uuid.UUID `json:"tax_category_id,omitempty"`
	DebitAccountID              uuid.UUID  `json:"debit_account_id"`
	InventoryItemID             *uuid.UUID `json:"inventory_item_id,omitempty"`
	InventoryID                 *uuid.UUID `json:"inventory_id,omitempty"`
	WareHouseInventoryAccountID *uuid.UUID `json:"warehouse_inventory_account_id,omitempty"`
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
