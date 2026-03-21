package invoice_engine_services

import (
	"context"
	invoice_engine_contracts "sharpnr-engine/service_engines/invoicing_engine/contracts"

	"github.com/google/uuid"
)

type InvoiceEngineServiceInterface interface {
	GenerateInvoice(ctx context.Context, request invoice_engine_contracts.CreateInvoiceRequest) (interface{}, error)
	IssueInvoice(ctx context.Context, invoiceID uuid.UUID) (interface{}, error)
}
