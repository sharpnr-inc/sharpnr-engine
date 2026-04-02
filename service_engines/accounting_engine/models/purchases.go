package accounting_models

import (
	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type PurchaseInvoice struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`
	FiscalYearID     uuid.UUID `gorm:"type:uuid;not null;index"`
	SupplierID       uuid.UUID `gorm:"type:uuid;not null;index"`

	// link to your existing invoicing engine if needed
	// PurchaseOrderID  *uuid.UUID `gorm:"type:uuid;index"` // optional, if PO raised first

	// Invoice Should Not duplicate With Context of the SuppliersID
	InvoiceNumber string    // supplier's bill number e.g. "RAJ-2026-041"
	OurReference  string    // your internal ref e.g. "PINV-2026-001"
	InvoiceDate   time.Time `gorm:"not null"`
	DueDate       *time.Time

	Status accountingconstants.AccountingPurchaseInvoiceStatus `gorm:"size:20;default:'DRAFT'"`
	// DRAFT | POSTED | PARTIALLY_PAID | PAID | CANCELLED

	PaymentStatus string `gorm:"size:20;default:'UNPAID'"`
	// UNPAID | PARTIAL | PAID

	// accounts
	PayableAccountID uuid.UUID `gorm:"type:uuid;not null"` // Accounts Payable

	// totals
	SubTotal    float64 `gorm:"default:0"`
	TaxAmount   float64 `gorm:"default:0"`
	TotalAmount float64 `gorm:"default:0"`
	AmountPaid  float64 `gorm:"default:0"`
	AmountDue   float64 `gorm:"default:0"` // TotalAmount - AmountPaid

	// accounting linkage (set when POSTED)
	JournalEntryID *uuid.UUID `gorm:"type:uuid"`

	Narration string
	CreatedBy uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PurchaseInvoice) TableName() string {
	return "accounting_engine.purchase_invoice"
}

// ── Purchase Invoice Line ──
type PurchaseInvoiceLine struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	PurchaseInvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	// what was bought
	Description string  `gorm:"not null"`
	Quantity    float64 `gorm:"default:1"`
	UnitPrice   float64 `gorm:"default:0"`

	// accounting
	DebitAccountID uuid.UUID `gorm:"type:uuid;not null"`
	// ↑ COGS account OR Inventory Asset account OR Expense account
	// whoever calls the API decides which account to debit

	TaxCategoryID *uuid.UUID `gorm:"type:uuid"` // links to your existing tax system

	// computed
	LineAmount  float64 `gorm:"default:0"` // quantity * unit_price
	TaxAmount   float64 `gorm:"default:0"` // computed from TaxCategory
	TotalAmount float64 `gorm:"default:0"` // LineAmount + TaxAmount

	// inventory linkage (optional — set if goods go to inventory)
	InventoryItemID *uuid.UUID `gorm:"type:uuid"`
	InventoryID     *uuid.UUID `gorm:"type:uuid"`
	// when set, inventory quantity updates automatically on POST

	CreatedAt time.Time
}

func (PurchaseInvoiceLine) TableName() string {
	return "accounting_engine.purchase_invoice_line"
}

// ── Payment against a Purchase Invoice ──
type PurchasePayment struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	PurchaseInvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`
	SupplierID        uuid.UUID `gorm:"type:uuid;not null;index"`

	PaymentDate   time.Time `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	PaymentMethod string    `gorm:"size:20"` // CASH | BANK | UPI | CHEQUE

	// accounting
	PayFromAccountID uuid.UUID  `gorm:"type:uuid;not null"` // Cash / Bank account
	JournalEntryID   *uuid.UUID `gorm:"type:uuid"`          // auto-created

	Reference string // UTR, cheque number, etc.
	Notes     string

	CreatedBy uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
}

func (PurchasePayment) TableName() string {
	return "accounting_engine.purchase_payment"
}
