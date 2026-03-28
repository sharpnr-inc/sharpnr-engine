package accounting_service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	dbtx "sharpnr-engine/core/dbTx"

	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_engine_helpers "sharpnr-engine/service_engines/accounting_engine/helpers"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *accountingServices) GetLegalEntityLedgerEntry(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
	legalEntityID uuid.UUID,
) (accountingdtos.LedgerPaginationResponse, error) {

	var res accountingdtos.LedgerPaginationResponse

	// Hardcoded pagination
	page := 1
	limit := 10

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// 1️⃣ Get Accounting Book from LegalEntity
		book, err := s.accRepo.GetAccountingBookByLegalEntityID(txCtx, legalEntityID)
		if err != nil {
			return err
		}

		// 2️⃣ Get Fiscal Year based on current date
		fy, err := s.accRepo.GetOpenFiscalYearByBookAndDate(txCtx, book.ID, time.Now())
		if err != nil {
			return err
		}

		// 3️⃣ Fetch Ledger Entries
		ledger, totalRows, err := s.accRepo.GetLedgerEntries(
			txCtx,
			book.ID,
			fy.ID,
			nil, // no account filter
			page,
			limit,
		)
		if err != nil {
			return err
		}

		// 4️⃣ Calculate total pages
		totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

		// 5️⃣ Build response
		res = accountingdtos.LedgerPaginationResponse{
			Data:       ledger,
			Page:       page,
			Limit:      limit,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		}

		return nil
	})

	if err != nil {
		return res, err
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *accountingServices) GetJournalEntryDetails(
	ctx context.Context,
	entryID uuid.UUID,
) (accountingdtos.JournalEntryDetailsDTO, error) {

	var res accountingdtos.JournalEntryDetailsDTO

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		entry, err := s.accRepo.GetJournalEntryByID(txCtx, entryID)
		if err != nil {
			return err
		}

		if entry == nil {
			return fmt.Errorf("journal entry not found")
		}

		lines, err := s.accRepo.GetJournalLinesWithAccount(txCtx, entryID)
		if err != nil {
			return err
		}

		res = accountingdtos.JournalEntryDetailsDTO{
			ID:        entry.ID,
			EntryDate: entry.EntryDate,
			Reference: entry.Reference,
			Narration: entry.Narration,
			Lines:     lines,
		}

		return nil
	})

	if err != nil {
		return res, err
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *accountingServices) AddJournalEntry(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	entryRequest accountingdtos.JournalEntryRequestDto,
) error {

	if err := entryRequest.ValidateRequest(); err != nil {
		return err
	}

	if len(entryRequest.JournalLines) == 0 {
		return errors.New("journal entry must contain at least one journal line")
	}

	// Validate each journal line
	for _, line := range entryRequest.JournalLines {
		if err := line.ValidateDto(); err != nil {
			return err
		}
	}

	// Validate debit-credit balance
	if err := accounting_engine_helpers.ValidateJournalBalance(entryRequest.JournalLines); err != nil {
		return err
	}

	journalEntry := &accounting_models.JournalEntry{
		ID:               uuid.New(),
		AccountingBookID: entryRequest.AccountingBookID,
		FiscalYearID:     entryRequest.FiscalYearID,
		EntryDate:        entryRequest.EntryDate,
		Reference:        entryRequest.Reference,
		Narration:        entryRequest.Narration,
		EntryType:        entryRequest.EntryType,
		CreatedAt:        entryRequest.CreatedAt,
		UpdatedAt:        time.Now(),
	}

	if err := s.accRepo.CreateJournalEntry(ctx, journalEntry); err != nil {
		return err
	}

	// Prepare journal lines

	for _, l := range entryRequest.JournalLines {

		lines := &accounting_models.JournalLine{
			ID:             uuid.New(),
			JournalEntryID: journalEntry.ID,
			AccountID:      l.AccountID,
			Debit:          l.Debit,
			Credit:         l.Credit,
			PartyID:        l.PartyID,
			CreatedAt:      time.Now(),
		}

		if err := s.accRepo.CreateJournalLine(ctx, lines); err != nil {
			return err
		}
	}

	return nil

}
