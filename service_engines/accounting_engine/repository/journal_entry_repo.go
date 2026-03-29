package accounting_repo

import (
	"context"
	"errors"
	dbtx "sharpnr-engine/core/dbTx"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) CreateJournalEntry(ctx context.Context, entry *accounting_models.JournalEntry) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}
	return db.WithContext(ctx).Create(entry).Error
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetJournalEntriesByBookAndPeriod(ctx context.Context, bookID uuid.UUID, from time.Time, to time.Time) ([]accounting_models.JournalEntry, error) {
	var entries []accounting_models.JournalEntry

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where(
			"accounting_book_id = ? AND entry_date >= ? AND entry_date <= ?",
			bookID,
			from,
			to,
		).
		Order("entry_date ASC, created_at ASC").
		Find(&entries).
		Error

	if err != nil {
		return nil, err
	}

	return entries, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetJournalEntryByID(ctx context.Context, entryID uuid.UUID) (*accounting_models.JournalEntry, error) {
	var entry accounting_models.JournalEntry

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Where("id = ?", entryID).
		First(&entry).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &entry, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetJournalEntryTotals(
	ctx context.Context,
	entryID uuid.UUID,
) (float64, float64, error) {

	type result struct {
		TotalDebit  float64
		TotalCredit float64
	}

	var res result

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.
		WithContext(ctx).
		Model(&accounting_models.JournalLine{}).
		Select(`
            COALESCE(SUM(debit), 0) AS total_debit,
            COALESCE(SUM(credit), 0) AS total_credit
        `).
		Where("journal_entry_id = ?", entryID).
		Scan(&res).
		Error

	if err != nil {
		return 0, 0, err
	}

	return res.TotalDebit, res.TotalCredit, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *accountingRepository) GetJournalEntriesByAccountingBookID(
	ctx context.Context,
	bookID uuid.UUID,
) ([]accounting_models.JournalEntry, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var entries []accounting_models.JournalEntry

	err := db.
		Where("accounting_book_id = ?", bookID).
		Order("entry_date DESC").
		Find(&entries).
		Error

	if err != nil {
		return nil, err
	}

	return entries, nil
}
