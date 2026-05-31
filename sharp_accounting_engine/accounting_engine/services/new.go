package accounting_service

import (
	"context"

	"github.com/google/uuid"
)

// this is a method for getting the summary of accounts for a legal entity.

func (s *accountingServices) GetAccountsSummaryOfLegalEntity(ctx context.Context, orgID uuid.UUID, emplID uuid.UUID, legalEntityID uuid.UUID) {

	// s.accRepo.GetAccountingBookByLegalEntityID()

	// s.accRepo.GetOpenFiscalYearByBookAndDate()

	// s.accRepo.GetJournalEntriesByBookAndPeriod()

	panic("unimplemented")
}
