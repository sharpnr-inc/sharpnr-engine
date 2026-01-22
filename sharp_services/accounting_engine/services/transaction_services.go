package accounting_service

import (
	"context"
	"errors"
	"fmt"
	accountingdtos "sharpnr-engine/sharp_services/accounting_engine/dto"
	accounting_repo "sharpnr-engine/sharp_services/accounting_engine/repository"

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

	if book.OrganizationID != orgID {
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

func (s *accountingServices) GetTransactionDetails(
	ctx context.Context,
	orgID uuid.UUID,
	entryID uuid.UUID,
) (*accountingdtos.JournalEntryDetailDTO, error) {

	entry, lines, err := s.accRepo.GetJournalEntryWithLines(ctx, entryID)
	if err != nil {
		return nil, err
	}

	// 🔐 Optional safety: validate org via book
	book, err := s.accRepo.GetAccountingBookByID(ctx, entry.AccountingBookID)
	if err != nil {
		return nil, err
	}
	if book.OrganizationID != orgID {
		return nil, errors.New("unauthorized access")
	}

	resp := &accountingdtos.JournalEntryDetailDTO{
		ID:        entry.ID,
		EntryDate: entry.EntryDate,
		Reference: entry.Reference,
		Narration: entry.Narration,
		Lines:     []accountingdtos.JournalLineDTO{},
	}

	for _, l := range lines {
		resp.Lines = append(resp.Lines, accountingdtos.JournalLineDTO{
			AccountID:   l.AccountID,
			AccountName: l.Account.Name,
			AccountCode: l.Account.Code,
			Debit:       l.Debit,
			Credit:      l.Credit,
		})
		resp.TotalDebit += l.Debit
		resp.TotalCredit += l.Credit
	}

	return resp, nil
}
