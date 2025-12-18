package accounting_repo

import (
	"context"
	"errors"
	"sharpnr-engine/modules/accounting/accounting_models"
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
