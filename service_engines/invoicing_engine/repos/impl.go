package invoice_engine_repos

import (
	"context"

	dbtx "sharpnr-engine/pkg/dbTx"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"

	"github.com/google/uuid"
)

// ─────────────────────────────────────────────────────────────────────────────
// Invoice
// ─────────────────────────────────────────────────────────────────────────────

func (r *invoiceEngineRepo) GetInvoiceByID(
	ctx context.Context,
	invoiceID uuid.UUID,
) (*invoice_engine_models.Invoice, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var invoice invoice_engine_models.Invoice
	if err := db.Where("id = ?", invoiceID).First(&invoice).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *invoiceEngineRepo) UpdateInvoicePaidAmount(
	ctx context.Context,
	invoiceID uuid.UUID,
	paidPaise int64,
	paymentStatus string,
) error {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.Model(&invoice_engine_models.Invoice{}).
		Where("id = ?", invoiceID).
		Updates(map[string]any{
			"paid_paise":     paidPaise,
			"payment_status": paymentStatus,
		}).Error
}

func (r *invoiceEngineRepo) GetInvoicesBySource(
	ctx context.Context,
	source string,
	sourceRefID uuid.UUID,
) ([]*invoice_engine_models.Invoice, error) {

	var invoices []*invoice_engine_models.Invoice

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	if err := db.
		Where("source = ? AND source_ref_id = ?", source, sourceRefID).
		Order("created_at DESC").
		Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Invoice Items
// ─────────────────────────────────────────────────────────────────────────────

func (r *invoiceEngineRepo) GetInvoiceItems(
	ctx context.Context,
	invoiceID uuid.UUID,
) ([]invoice_engine_models.InvoiceItem, error) {

	var items []invoice_engine_models.InvoiceItem
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	if err := db.
		Where("invoice_id = ?", invoiceID).
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Tax Breakup
// ─────────────────────────────────────────────────────────────────────────────

func (r *invoiceEngineRepo) GetInvoiceTaxBreakups(
	ctx context.Context,
	invoiceID uuid.UUID,
) ([]invoice_engine_models.InvoiceTaxBreakup, error) {

	var breakups []invoice_engine_models.InvoiceTaxBreakup
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	if err := db.
		Where("invoice_id = ?", invoiceID).
		Find(&breakups).Error; err != nil {
		return nil, err
	}
	return breakups, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Payment Links
// ─────────────────────────────────────────────────────────────────────────────

func (r *invoiceEngineRepo) GetInvoicePaymentLinkByRef(
	ctx context.Context,
	invoiceID uuid.UUID,
	paymentRef string,
) (*invoice_engine_models.InvoicePaymentLink, error) {

	var link invoice_engine_models.InvoicePaymentLink
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	if err := db.
		Where("invoice_id = ? AND payment_ref = ?", invoiceID, paymentRef).
		First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *invoiceEngineRepo) GetPaymentLinks(
	ctx context.Context,
	invoiceID uuid.UUID,
) ([]invoice_engine_models.InvoicePaymentLink, error) {

	var links []invoice_engine_models.InvoicePaymentLink
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	if err := db.
		Where("invoice_id = ?", invoiceID).
		Order("created_at ASC").
		Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}
