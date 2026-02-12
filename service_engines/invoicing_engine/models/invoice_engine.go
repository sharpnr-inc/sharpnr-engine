package invoice_engine_models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------------------
// 1️⃣ INVOICE HEADER
// -------------------------------
type Invoice struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	AccountingBookID uuid.UUID  `gorm:"type:uuid;not null;index"`
	LegalEntityID    *uuid.UUID `gorm:"type:uuid;index"`

	InvoiceNumber string     `gorm:"type:varchar(50);uniqueIndex;not null"`
	InvoiceDate   time.Time  `gorm:"type:date;not null"`
	DueDate       *time.Time `gorm:"type:date"`

	CustomerName string  `gorm:"type:varchar(200);not null"`
	CustomerGST  *string `gorm:"type:varchar(20)"`

	// Financial Totals
	SubTotal float64 `gorm:"type:decimal(15,2);not null"`
	TaxTotal float64 `gorm:"type:decimal(15,2);not null"`
	Total    float64 `gorm:"type:decimal(15,2);not null"`

	// Status: DRAFT, ISSUED, PAID, CANCELLED
	Status string `gorm:"type:varchar(20);index;default:'DRAFT'"`

	// Source: STANDALONE, HOTEL_STAY, RESTAURANT_SESSION
	Source      string     `gorm:"type:varchar(50);index"`
	SourceRefID *uuid.UUID `gorm:"type:uuid;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Invoice) TableName() string {
	return "invoicing_engine.invoice"
}

// -------------------------------
// 2️⃣ INVOICE LINE ITEMS
// -------------------------------
type InvoiceItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	Description string  `gorm:"type:text;not null"`
	Quantity    float64 `gorm:"type:decimal(12,2);default:1"`
	Rate        float64 `gorm:"type:decimal(15,2);not null"`
	Amount      float64 `gorm:"type:decimal(15,2);not null"`

	TaxCategoryID *uuid.UUID `gorm:"type:uuid;index"`
}

func (InvoiceItem) TableName() string {
	return "invoicing_engine.invoice_item"
}

// -------------------------------
// 3️⃣ TAX BREAKUP (Statutory Compliance)
// -------------------------------
type InvoiceTaxBreakup struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	// TaxType: CGST, SGST, IGST, VAT, SERVICE_TAX
	TaxType string  `gorm:"type:varchar(20);not null"`
	Rate    float64 `gorm:"type:decimal(5,2);not null"` // e.g., 9.00
	Amount  float64 `gorm:"type:decimal(15,2);not null"`

	CreatedAt time.Time
}

func (InvoiceTaxBreakup) TableName() string {
	return "invoicing_engine.invoice_tax_breakup"
}

// -------------------------------
// 4️⃣ PAYMENT MAPPING (Reconciliation)
// -------------------------------
type InvoicePaymentLink struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	InvoiceID uuid.UUID `gorm:"type:uuid;not null;index"`

	// Link to the generic payment record in your accounting module
	PaymentID uuid.UUID `gorm:"type:uuid;not null;index"`

	AppliedAmount float64 `gorm:"type:decimal(15,2);not null"`
}

func (InvoicePaymentLink) TableName() string {
	return "invoicing_engine.invoice_payment_link"
}
