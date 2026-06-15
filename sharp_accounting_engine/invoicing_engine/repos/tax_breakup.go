package invoice_engine_repos

import (
	"context"
	"fmt"
	dbtx "sharpnr-engine/pkg/dbTx"
	invoice_engine_models "sharpnr-engine/sharp_accounting_engine/invoicing_engine/models"
)

func (r *invoiceEngineRepo) CreateInvoiceTaxBreakUp(
	ctx context.Context,
	taxBreakUp *invoice_engine_models.InvoiceTaxBreakup,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	gormDB := db.WithContext(ctx)

	if err := gormDB.Create(taxBreakUp).Error; err != nil {
		return fmt.Errorf("repository: failed to create invoice tax breakup: %w", err)
	}

	return nil
}
