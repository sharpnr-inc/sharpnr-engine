package accounting_service

import (
	"context"
	"fmt"
	"time"

	dbtx "sharpnr-engine/core/dbTx"

	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_engine_helpers "sharpnr-engine/service_engines/accounting_engine/helpers"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/* ---------------------------------- <***> --------------------------------- */

func (s *accountingServices) CreateJournalEntry(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	journalEntryRequest accountingdtos.JournalEntryExternalRequestDto,
) (journalEntryID uuid.UUID, err error) {

	err = s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// Ensure lines exist
		if len(journalEntryRequest.Lines) == 0 {
			return fmt.Errorf("journal entry must contain at least one line")
		}

		// Validate debit/credit balance
		if err := accounting_engine_helpers.ValidateJournalBalance(
			journalEntryRequest.Lines,
		); err != nil {
			return err
		}

		accBook, err := s.accRepo.GetAccountingBookByLegalEntityID(txCtx, legalEntityID)
		if err != nil {
			return err
		}

		// Fetch accounting book using request value (correct source of truth)
		// accBook, err := s.accRepo.GetAccountingBookByID(
		// 	txCtx,
		// 	journalEntryRequest.AccountingBookID,
		// )
		// if err != nil {
		// 	return err
		// }

		if accBook == nil {
			return fmt.Errorf("accounting book not found")
		}

		// Ensure book belongs to legal entity
		if accBook.LegalEntityID != legalEntityID {
			return fmt.Errorf("accounting book does not belong to legal entity")
		}

		// Validate fiscal year
		fy, err := s.accRepo.GetOpenFiscalYearByBookAndDate(
			txCtx,
			accBook.ID,
			time.Now(),
		)
		if err != nil {
			return err
		}

		if fy == nil {
			return fmt.Errorf("no open fiscal year found for accounting book")
		}

		// Corrected logic (previous version was reversed)
		if fy.Status == accountingconstants.Closed {
			return fmt.Errorf("fiscal year is closed for journal entries")
		}

		// Create journal entry header
		newEntry := accounting_models.JournalEntry{
			ID:               uuid.New(),
			AccountingBookID: accBook.ID,
			FiscalYearID:     fy.ID,
			EntryDate:        time.Now(),
			Reference:        journalEntryRequest.Reference,
			Narration:        journalEntryRequest.Narration,
			EntryType:        journalEntryRequest.EntryType,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}

		if err := s.accRepo.CreateJournalEntry(txCtx, &newEntry); err != nil {
			return err
		}

		// assign return value correctly
		journalEntryID = newEntry.ID

		// Insert journal lines
		for _, line := range journalEntryRequest.Lines {

			newLine := accounting_models.JournalLine{
				ID:             uuid.New(),
				JournalEntryID: newEntry.ID,
				AccountID:      line.AccountID,
				Debit:          line.Debit,
				Credit:         line.Credit,
				PartyID:        line.PartyID,
				CreatedAt:      time.Now(),
			}

			if err := s.accRepo.CreateJournalLine(txCtx, &newLine); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return journalEntryID, nil
}

/* ---------------------------------- <***> --------------------------------- */
