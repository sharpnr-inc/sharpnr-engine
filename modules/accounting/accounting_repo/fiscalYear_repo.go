package accounting_repo

import (
	"context"
	"sharpnr-engine/modules/accounting/accounting_models"
	"time"

	"github.com/google/uuid"
)

// CloseFiscalYear implements AccountingRepositoryInterface.
func (r *accountingRepository) CloseFiscalYear(ctx context.Context, fiscalYearID uuid.UUID) error {
	panic("unimplemented")
}

// CreateFiscalYear implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateFiscalYear(ctx context.Context, fy *accounting_models.FiscalYear) error {
	return r.db.WithContext(ctx).Create(fy).Error
}

// GetOpenFiscalYearByBookAndDate implements AccountingRepositoryInterface.
func (r *accountingRepository) GetOpenFiscalYearByBookAndDate(ctx context.Context, bookID uuid.UUID, date time.Time) (*accounting_models.FiscalYear, error) {
	panic("unimplemented")
}
