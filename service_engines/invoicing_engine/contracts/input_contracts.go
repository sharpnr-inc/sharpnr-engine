package invoice_engine_contracts

import (
	"time"

	"github.com/google/uuid"
)

type CreateInvoiceRequest struct {
	OrganizationID uuid.UUID
	LegalEntityID  *uuid.UUID

	CustomerID        *uuid.UUID
	CustomerName      string
	CustomerGST       *string
	CustomerStateCode *string // snapshotted — drives CGST/SGST vs IGST

	Source      string     // RESTAURANT_SESSION, HOTEL_STAY, STANDALONE
	SourceRefID *uuid.UUID // session ID, stay ID

	Items []InvoiceItemRequest
}

type InvoiceItemRequest struct {
	InvoiceID     uuid.UUID
	Description   string
	Quantity      float64
	Rate          float64
	TaxCategoryID *uuid.UUID
	HSNCode       *string
}

type RecordPaymentRequest struct {
	InvoiceID      uuid.UUID
	PaymentRef     string // Payment.ID from payment engine — opaque
	PaymentMethod  string // snapshot: CASH, UPI etc.
	AllocatedPaise int64
	PaidAt         time.Time
}

type ReversePaymentRequest struct {
	InvoiceID  uuid.UUID
	PaymentRef string
}
