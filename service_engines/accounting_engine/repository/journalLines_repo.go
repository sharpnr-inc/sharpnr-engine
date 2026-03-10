package accounting_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"time"

	"github.com/google/uuid"
)

// GetJournalLinesByEntryID implements AccountingRepositoryInterface.
func (r *accountingRepository) GetJournalLinesByEntryID(
	ctx context.Context,
	entryID uuid.UUID,
) ([]accounting_models.JournalLine, error) {

	var lines []accounting_models.JournalLine

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Preload("Account").
		Where("journal_entry_id = ?", entryID).
		Order("created_at ASC").
		Find(&lines).
		Error

	if err != nil {
		return nil, err
	}

	return lines, nil
}

// GetJournalLinesByAccountAndPeriod implements AccountingRepositoryInterface.
func (r *accountingRepository) GetJournalLinesByAccountAndPeriod(
	ctx context.Context,
	accountID uuid.UUID,
	from time.Time,
	to time.Time,
) ([]accounting_models.JournalLine, error) {

	var lines []accounting_models.JournalLine

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	err := db.WithContext(ctx).
		Joins("JOIN journal_entries je ON je.id = journal_lines.journal_entry_id").
		Where(
			"journal_lines.account_id = ? AND je.entry_date >= ? AND je.entry_date <= ?",
			accountID,
			from,
			to,
		).
		Order("je.entry_date ASC, journal_lines.created_at ASC").
		Find(&lines).
		Error

	if err != nil {
		return nil, err
	}

	return lines, nil
}
