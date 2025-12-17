package accounting_repo

import (
	"context"
	"sharpnr-engine/modules/accounting/accounting_models"
	"time"

	"github.com/google/uuid"
)

// CreateJournalEntry implements AccountingRepositoryInterface.
func (r *accountingRepository) CreateJournalEntry(ctx context.Context, entry *accounting_models.JournalEntry) error {
	panic("unimplemented")
}

// GetJournalEntriesByBookAndPeriod implements AccountingRepositoryInterface.
func (r *accountingRepository) GetJournalEntriesByBookAndPeriod(ctx context.Context, bookID uuid.UUID, from time.Time, to time.Time) ([]accounting_models.JournalEntry, error) {
	panic("unimplemented")
}

// GetJournalEntryByID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetJournalEntryByID(ctx context.Context, entryID uuid.UUID) (*accounting_models.JournalEntry, error) {
	panic("unimplemented")
}
