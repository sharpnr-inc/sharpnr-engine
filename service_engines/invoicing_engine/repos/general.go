package invoice_engine_repos

import (
	"context"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"

	"github.com/google/uuid"
)

func (r *invoiceEngineRepo) CreateInvoice(ctx context.Context, invoice *invoice_engine_models.Invoice) error {
	panic("Unimplemented")
}

func (r *invoiceEngineRepo) UpdateInvoiceStatus(ctx context.Context, invoiceID uuid.UUID, status string) error {
	panic("Unimplemented")
}

func (r *invoiceEngineRepo) CreateInvoiceItem(ctx context.Context, invoiceItem *invoice_engine_models.InvoiceItem) error {
	panic("Unimplemented")
}

func (r *invoiceEngineRepo) CreateInvoiceTaxBreakUp(ctx context.Context, taxBreakUp *invoice_engine_models.InvoiceTaxBreakup) error {
	panic("Unimplemented")
}
