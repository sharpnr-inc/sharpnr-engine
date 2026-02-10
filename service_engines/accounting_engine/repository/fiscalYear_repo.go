package accounting_repo

import (
	"context"
	"errors"
	"fmt"

	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CloseFiscalYear implements AccountingRepositoryInterface.
func (r *accountingRepository) CloseFiscalYear(ctx context.Context, fiscalYearID uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. Check if the fiscal year exists and isn't already closed
		var fiscalYear accounting_models.FiscalYear
		if err := tx.Where("id = ? AND status != ?", fiscalYearID, "closed").First(&fiscalYear).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("fiscal year not found or already closed")
			}
			return err
		}

		// 2. Perform business logic (e.g., set end date, validate balances)
		// This is where you'd trigger balance carry-overs if required.

		// 3. Update the status to 'closed'
		updateResult := tx.Model(&fiscalYear).Update("status", "closed")
		if updateResult.Error != nil {
			return updateResult.Error
		}

		if updateResult.RowsAffected == 0 {
			return errors.New("failed to close fiscal year: no rows updated")
		}

		return nil
	})
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
