package invoice_engine_services

import (
	"context"
	invoice_engine_contracts "sharpnr-engine/sharp_accounting_engine/invoicing_engine/contracts"
	invoice_engine_models "sharpnr-engine/sharp_accounting_engine/invoicing_engine/models"

	"github.com/google/uuid"
)

type InvoiceEngineServiceInterface interface {

	// Invoice lifecycle
	GenerateInvoice(ctx context.Context, req invoice_engine_contracts.CreateInvoiceRequest) (*invoice_engine_models.Invoice, error)
	IssueInvoice(ctx context.Context, invoiceID uuid.UUID) (*invoice_engine_models.Invoice, error)
	CancelInvoice(ctx context.Context, invoiceID uuid.UUID, reason string) error

	// Line items
	AddInvoiceItems(ctx context.Context, invoiceID uuid.UUID, items []invoice_engine_contracts.InvoiceItemRequest) error

	// Payment recording
	RecordPayment(ctx context.Context, req invoice_engine_contracts.RecordPaymentRequest) error
	ReversePayment(ctx context.Context, req invoice_engine_contracts.ReversePaymentRequest) error

	// Queries
	GetInvoiceByID(ctx context.Context, invoiceID uuid.UUID) (*invoice_engine_models.Invoice, error)
	GetInvoiceWithDetails(ctx context.Context, invoiceID uuid.UUID) (*invoice_engine_contracts.InvoiceWithDetails, error) // ← contract not model
	GetInvoicesBySource(ctx context.Context, source string, sourceRefID uuid.UUID) ([]*invoice_engine_models.Invoice, error)
	GetOutstandingAmount(ctx context.Context, invoiceID uuid.UUID) (int64, error)
}
