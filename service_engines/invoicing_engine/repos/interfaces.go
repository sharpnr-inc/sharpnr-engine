package invoice_engine_repos

import (
	"context"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"
	"time"

	"github.com/google/uuid"
)

type InvoicingEngineRepositoryInterface interface {

	// ─── Invoice ──────────────────────────────────────────────────────────────

	CreateInvoice(ctx context.Context, invoice *invoice_engine_models.Invoice) error
	GetInvoiceByID(ctx context.Context, invoiceID uuid.UUID) (*invoice_engine_models.Invoice, error)
	UpdateInvoiceStatus(ctx context.Context, invoiceID uuid.UUID, status string) error
	UpdateInvoicePaidAmount(ctx context.Context, invoiceID uuid.UUID, paidPaise int64, paymentStatus string) error
	GetInvoicesBySource(ctx context.Context, source string, sourceRefID uuid.UUID) ([]*invoice_engine_models.Invoice, error)

	// ─── Invoice Items ────────────────────────────────────────────────────────

	CreateInvoiceItem(ctx context.Context, invoiceItem *invoice_engine_models.InvoiceItem) error
	GetInvoiceItems(ctx context.Context, invoiceID uuid.UUID) ([]invoice_engine_models.InvoiceItem, error)

	// ─── Tax Breakup ──────────────────────────────────────────────────────────

	CreateInvoiceTaxBreakUp(ctx context.Context, taxBreakUp *invoice_engine_models.InvoiceTaxBreakup) error
	GetInvoiceTaxBreakups(ctx context.Context, invoiceID uuid.UUID) ([]invoice_engine_models.InvoiceTaxBreakup, error)

	// ─── Payment Links ────────────────────────────────────────────────────────

	CreateInvoicePaymentLink(ctx context.Context, invoicePaymentLink *invoice_engine_models.InvoicePaymentLink) error
	GetInvoicePaymentLinkByInvoiceID(ctx context.Context, invoiceID uuid.UUID) (*invoice_engine_models.InvoicePaymentLink, error)
	GetInvoicePaymentLinkByRef(ctx context.Context, invoiceID uuid.UUID, paymentRef string) (*invoice_engine_models.InvoicePaymentLink, error)
	GetPaymentLinks(ctx context.Context, invoiceID uuid.UUID) ([]invoice_engine_models.InvoicePaymentLink, error)
	UpdateInvoicePaymentLinkStatus(ctx context.Context, invoicePaymentLinkID uuid.UUID, status string, reversedAt *time.Time) error
}
