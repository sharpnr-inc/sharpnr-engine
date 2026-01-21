package accounting_repo

import (
	"context"
	"errors"
	accounting_models "sharpnr-engine/modules/accounting/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateJournalEntry implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateJournalEntry(ctx context.Context, entry *accounting_models.JournalEntry) error {
	return r.db.WithContext(ctx).Create(entry).Error
}

// GetJournalEntriesByBookAndPeriod implements AccountingRepositoryInterface.
func (r *accountingRepository) GetJournalEntriesByBookAndPeriod(ctx context.Context, bookID uuid.UUID, from time.Time, to time.Time) ([]accounting_models.JournalEntry, error) {
	var entries []accounting_models.JournalEntry

	err := r.db.WithContext(ctx).
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

// GetJournalEntryByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetJournalEntryByID(ctx context.Context, entryID uuid.UUID) (*accounting_models.JournalEntry, error) {
	var entry accounting_models.JournalEntry

	err := r.db.WithContext(ctx).
		Preload("Lines").
		Preload("Lines.Account").
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

func (r *accountingRepository) GetJournalEntries(
	ctx context.Context,
	bookID uuid.UUID,
	from time.Time,
	to time.Time,
	limit int,
	offset int,
) ([]accounting_models.JournalEntry, int64, error) {

	var entries []accounting_models.JournalEntry
	var total int64

	baseQuery := r.db.
		WithContext(ctx).
		Model(&accounting_models.JournalEntry{}).
		Where(
			"accounting_book_id = ? AND entry_date BETWEEN ? AND ?",
			bookID,
			from,
			to,
		)

	// Count
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch
	err := baseQuery.
		Order("entry_date DESC, created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&entries).
		Error

	if err != nil {
		return nil, 0, err
	}

	return entries, total, nil
}

func (r *accountingRepository) GetJournalEntryTotals(
	ctx context.Context,
	entryID uuid.UUID,
) (float64, float64, error) {

	type result struct {
		TotalDebit  float64
		TotalCredit float64
	}

	var res result

	err := r.db.
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

func (r *accountingRepository) GetJournalEntryWithLines(
	ctx context.Context,
	entryID uuid.UUID,
) (*accounting_models.JournalEntry, []accounting_models.JournalLine, error) {

	var entry accounting_models.JournalEntry
	var lines []accounting_models.JournalLine

	// 1️⃣ Fetch Journal Entry
	if err := r.db.
		WithContext(ctx).
		First(&entry, "id = ?", entryID).
		Error; err != nil {
		return nil, nil, err
	}

	// 2️⃣ Fetch Journal Lines with Account info
	if err := r.db.
		WithContext(ctx).
		Preload("Account").
		Where("journal_entry_id = ?", entryID).
		Order("created_at ASC").
		Find(&lines).
		Error; err != nil {
		return nil, nil, err
	}

	return &entry, lines, nil
}
