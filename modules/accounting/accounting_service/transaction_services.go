package accounting_service

import (
	"context"
	"errors"
	"fmt"
	accountingdtos "sharpnr-engine/modules/accounting/accounting_dtos"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"time"

	"github.com/google/uuid"
)

// 🐕‍🦺 this is the service method for fetching the transactions
// ✨ Documentation needed
func (s *accountingServices) GetTransactions(
	ctx context.Context,
	orgID uuid.UUID,
	bookID uuid.UUID,
	from time.Time,
	to time.Time,
	limit int,
	offset int,
) ([]accountingdtos.JournalEntryListDTO, int64, error) {
	book, err := s.accRepo.GetAccountingBookByID(ctx, bookID)
	if err != nil {
		return nil, 0, err
	}

	if book.OrgID != orgID {
		return nil, 0, errors.New("unauthorized book access")
	}

	entries, total, err := s.accRepo.GetJournalEntries(
		ctx,
		bookID,
		from,
		to,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, err
	}
	fmt.Println(entries)

	resp := make([]accountingdtos.JournalEntryListDTO, 0, len(entries))
	for _, e := range entries {
		td, tc := sumJournalLines(e.ID, s.accRepo, ctx)

		resp = append(resp, accountingdtos.JournalEntryListDTO{
			ID:          e.ID,
			EntryDate:   e.EntryDate,
			Reference:   e.Reference,
			Narration:   e.Narration,
			TotalDebit:  td,
			TotalCredit: tc,
			CreatedAt:   e.CreatedAt,
		})
	}

	return resp, total, nil
}

// ⚡⚡⚡ Helper function for the summing the journal entries ⚡⚡⚡
func sumJournalLines(
	entryID uuid.UUID,
	repo accounting_repo.AccountingRepositoryInterface,
	ctx context.Context,
) (float64, float64) {

	totalDebit, totalCredit, err := repo.GetJournalEntryTotals(ctx, entryID)
	if err != nil {
		// In accounting UI, fail-safe = show zero instead of crashing
		return 0, 0
	}

	return totalDebit, totalCredit
}
