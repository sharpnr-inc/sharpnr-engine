package invoice_engine_models

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	OrganizationID   uuid.UUID
	AccountingBookID uuid.UUID

	LegalEntityID *uuid.UUID // optional

	InvoiceNumber string
	InvoiceDate   time.Time
	DueDate       *time.Time

	CustomerName string
	CustomerGST  *string

	SubTotal float64
	TaxTotal float64
	Total    float64

	Status string
	// DRAFT, ISSUED, PAID, CANCELLED

	Source string
	// STANDALONE, HOTEL, RESTAURANT

	SourceRefID *uuid.UUID

	CreatedAt time.Time
}

type InvoiceItem struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	InvoiceID uuid.UUID

	Description string
	Quantity    float64
	Rate        float64

	Amount float64

	TaxCategoryID *uuid.UUID // legal entity owned
}

type InvoiceTaxBreakup struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	InvoiceID uuid.UUID
	TaxType   string // CGST, SGST, IGST
	Rate      float64
	Amount    float64
}

type InvoicePaymentLink struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	InvoiceID uuid.UUID
	PaymentID uuid.UUID // accounting

	AppliedAmount float64
}
