package invoice_engine_models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// -------------------------------
// 1️⃣ INVOICE HEADER
// -------------------------------
type Invoice struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey"`
	OrganizationID uuid.UUID  `gorm:"type:uuid;not null;index"`
	LegalEntityID  *uuid.UUID `gorm:"type:uuid;index"`

	InvoiceNumber string     `gorm:"type:varchar(50);not null;uniqueIndex:idx_inv_num_org"`
	InvoiceDate   time.Time  `gorm:"type:date;not null"`
	DueDate       *time.Time `gorm:"type:date"`

	// Customer snapshot — never fetch live for issued invoices
	CustomerID        *uuid.UUID `gorm:"type:uuid;index"` // FK to CRM — opaque
	CustomerName      string     `gorm:"type:varchar(200);not null"`
	CustomerGST       *string    `gorm:"type:varchar(20)"`
	CustomerStateCode *string    `gorm:"type:varchar(3)"` // drives CGST/SGST vs IGST

	// Legal entity GSTIN snapshot — which registration applied at invoice time
	LegalEntityGSTIN *string `gorm:"type:varchar(20)"`

	// Financials — all stored as float, paise cached separately for calculations
	SubTotal float64 `gorm:"type:decimal(15,2);not null"`
	TaxTotal float64 `gorm:"type:decimal(15,2);not null"`
	Total    float64 `gorm:"type:decimal(15,2);not null"`

	// Cached payment state — updated by RecordPayment / ReversePayment
	PaidPaise  int64   `gorm:"default:0"`
	PaidAmount float64 `gorm:"-"` // derived: PaidPaise / 100

	// DRAFT, ISSUED, PAID, CANCELLED
	Status string `gorm:"type:varchar(20);not null;default:'DRAFT';index"`

	// UNPAID, PARTIALLY_PAID, PAID, OVERPAID
	PaymentStatus string `gorm:"type:varchar(20);not null;default:'UNPAID';index"`

	// Source — opaque to invoice engine
	Source      string     `gorm:"type:varchar(50);index"` // RESTAURANT_SESSION, HOTEL_STAY, STANDALONE
	SourceRefID *uuid.UUID `gorm:"type:uuid;index"`        // session ID, stay ID

	CreatedBy string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Invoice) TableName() string {
	return "invoice_engine.invoice"
}

// -------------------------------
// 2️⃣ INVOICE LINE ITEMS
// -------------------------------
type InvoiceItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	Description string  `gorm:"type:text;not null"`
	Quantity    float64 `gorm:"type:decimal(12,2);not null;default:1"`
	Rate        float64 `gorm:"type:decimal(15,2);not null"` // pre-tax unit price
	Amount      float64 `gorm:"type:decimal(15,2);not null"` // Quantity * Rate

	// Tax
	TaxCategoryID *uuid.UUID `gorm:"type:uuid;index"`
	HSNCode       *string    `gorm:"type:varchar(10)"`  // GST HSN / SAC code
	TaxRate       float64    `gorm:"type:decimal(5,2)"` // e.g. 18.00
	TaxAmount     float64    `gorm:"type:decimal(15,2)"`
	TotalAmount   float64    `gorm:"type:decimal(15,2);not null"` // Amount + TaxAmount

	CreatedAt time.Time
}

func (InvoiceItem) TableName() string {
	return "invoice_engine.invoice_item"
}

// -------------------------------
// 3️⃣ TAX BREAKUP
// Header level — one row per tax component
// CGST, SGST, IGST broken out separately for GST filing
// -------------------------------
type InvoiceTaxBreakup struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	// CGST, SGST, IGST, VAT
	TaxType string  `gorm:"type:varchar(20);not null"`
	Rate    float64 `gorm:"type:decimal(5,2);not null"`
	Amount  float64 `gorm:"type:decimal(15,2);not null"`

	CreatedAt time.Time
}

func (InvoiceTaxBreakup) TableName() string {
	return "invoice_engine.invoice_tax_breakup"
}

// -------------------------------
// 4️⃣ INVOICE PAYMENT LINK
// Invoice engine's side of the payment ↔ invoice relationship
// PaymentRef is opaque — no FK to payments schema
// -------------------------------
type InvoicePaymentLink struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Opaque reference to payment engine
	// Stores Payment.ID as string — no FK constraint
	PaymentRef    string `gorm:"type:varchar(100);not null;index"`
	PaymentMethod string `gorm:"type:varchar(30);not null"` // snapshot: CASH, UPI, PHONEPE

	CurrencyCode    string  `gorm:"type:char(3);not null;default:'INR'"`
	AllocatedPaise  int64   `gorm:"not null"`
	AllocatedAmount float64 `gorm:"-"` // derived

	// APPLIED, REVERSED, REFUNDED
	Status string `gorm:"type:varchar(20);not null;default:'APPLIED';index"`

	// Snapshot — so invoice engine never queries payment engine for display
	PaidAt time.Time `gorm:"not null"`

	CreatedAt  time.Time
	ReversedAt *time.Time
}

func (InvoicePaymentLink) TableName() string {
	return "invoice_engine.invoice_payment_link"
}
