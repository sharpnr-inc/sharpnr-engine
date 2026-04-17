package invoice_engine_repos

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"

	"github.com/google/uuid"
)

func (r *invoiceEngineRepo) CreateInvoice(
	ctx context.Context,
	invoice *invoice_engine_models.Invoice,
) error {
	db, ok := dbtx.ExtractTx(ctx) // Assuming dbtx extracts a *gorm.DB
	if !ok {
		db = r.db
	}

	// Use WithContext to pass request tracing/timeouts to the database
	if err := db.WithContext(ctx).Create(invoice).Error; err != nil {
		return fmt.Errorf("repository: failed to create invoice: %w", err)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *invoiceEngineRepo) UpdateInvoiceStatus(
	ctx context.Context,
	invoiceID uuid.UUID,
	status string,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	gormDB := db.WithContext(ctx)

	// Use Model with Where to perform a single-column update
	result := gormDB.Model(&invoice_engine_models.Invoice{}).
		Where("id = ?", invoiceID).
		Update("status", status)

	if result.Error != nil {
		return fmt.Errorf("repository: failed to update invoice status: %w", result.Error)
	}

	// Optional: Handle case where the ID does not exist
	if result.RowsAffected == 0 {
		return fmt.Errorf("repository: no invoice found with ID %s", invoiceID)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *invoiceEngineRepo) UpdatePaymentStatusOnInvoice(
	ctx context.Context,
	invoiceID uuid.UUID,
	paymentStatus string,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	gormDB := db.WithContext(ctx)

	// Update only the payment_status field for the given ID
	result := gormDB.Model(&invoice_engine_models.Invoice{}).
		Where("id = ?", invoiceID).
		Update("payment_status", paymentStatus)

	if result.Error != nil {
		return fmt.Errorf("repository: failed to update payment status: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("repository: invoice not found: %s", invoiceID)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
