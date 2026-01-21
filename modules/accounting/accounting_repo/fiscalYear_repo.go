package accounting_repo

import (
	"context"
	"errors"
	"fmt"
	accountingconstants "sharpnr-engine/modules/accounting/accounting_constants"
	accounting_models "sharpnr-engine/modules/accounting/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
func (r *accountingRepository) GetOpenFiscalYearByBookAndDate(
	ctx context.Context,
	bookID uuid.UUID,
	date time.Time,
) (*accounting_models.FiscalYear, error) {

	var fy accounting_models.FiscalYear

	err := r.db.
		WithContext(ctx).
		Where(`
			accounting_book_id = ?
			AND status = ?
			AND start_date <= ?
			AND end_date >= ?
		`,
			bookID,
			accountingconstants.Open,
			date,
			date,
		).
		First(&fy).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(
				"no open fiscal year found for book %s on date %s",
				bookID,
				date.Format("2006-01-02"),
			)
		}
		return nil, err
	}

	return &fy, nil
}
