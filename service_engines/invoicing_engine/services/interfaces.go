package invoice_engine_services

import (
	"context"
	invoice_engine_contracts "sharpnr-engine/service_engines/invoicing_engine/contracts"

	"github.com/google/uuid"
)

type InvoiceEngineServiceInterface interface {
	GenerateInvoice(ctx context.Context, request invoice_engine_contracts.CreateInvoiceRequest) error
	AddInvoiceItems(ctx context.Context, invoiceitems invoice_engine_contracts.InvoiceItemRequest) error
	IssueInvoice(ctx context.Context, invoiceID uuid.UUID) error
}
