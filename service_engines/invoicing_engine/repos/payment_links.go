package invoice_engine_repos

import (
	"context"
	"errors"
	"fmt"
	dbtx "sharpnr-engine/core/dbTx"
	invoice_engine_models "sharpnr-engine/service_engines/invoicing_engine/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *invoiceEngineRepo) CreateInvoicePaymentLink(
	ctx context.Context,
	invoicePaymentLink *invoice_engine_models.InvoicePaymentLink,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	gormDB := db.WithContext(ctx)

	if err := gormDB.Create(invoicePaymentLink).Error; err != nil {
		return fmt.Errorf("repository: failed to create invoice payment link: %w", err)
	}

	return nil
}

func (r *invoiceEngineRepo) GetInvoicePaymentLinkByInvoiceID(
	ctx context.Context,
	invoiceID uuid.UUID,
) (*invoice_engine_models.InvoicePaymentLink, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var link invoice_engine_models.InvoicePaymentLink
	gormDB := db.WithContext(ctx)

	err := gormDB.Where("invoice_id = ?", invoiceID).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Or return a specific sentinel error if preferred
		}
		return nil, fmt.Errorf("repository: failed to fetch payment link for invoice %s: %w", invoiceID, err)
	}

	return &link, nil
}

func (r *invoiceEngineRepo) UpdateInvoicePaymentLinkStatus(
	ctx context.Context,
	invoicePaymentLinkID uuid.UUID,
	status string,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	gormDB := db.WithContext(ctx)

	// We use a map for updates if we need to conditionally handle the ReversedAt timestamp
	updates := map[string]interface{}{
		"status": status,
	}

	if status == "REVERSED" {
		now := time.Now()
		updates["reversed_at"] = &now
	}

	result := gormDB.Model(&invoice_engine_models.InvoicePaymentLink{}).
		Where("id = ?", invoicePaymentLinkID).
		Updates(updates)

	if result.Error != nil {
		return fmt.Errorf("repository: failed to update payment link status: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("repository: invoice payment link not found: %s", invoicePaymentLinkID)
	}

	return nil
}
