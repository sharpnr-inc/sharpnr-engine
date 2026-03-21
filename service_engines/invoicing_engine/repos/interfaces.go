package invoice_engine_repos

import (
	"context"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"

	"github.com/google/uuid"
)

type InvoicingEngineRepositoryInterface interface {
	CreateInvoice(ctx context.Context, invoice *invoice_engine_models.Invoice) error
	UpdateInvoiceStatus(ctx context.Context, invoiceID uuid.UUID, status string) error

	CreateInvoiceItem(ctx context.Context, invoiceItem *invoice_engine_models.InvoiceItem) error
	CreateInvoiceTaxBreakUp(ctx context.Context, taxBreakUp *invoice_engine_models.InvoiceTaxBreakup) error

}
