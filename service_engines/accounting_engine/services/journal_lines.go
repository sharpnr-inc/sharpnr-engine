package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	accounting_models "sharpnr-engine/service_engines/accounting_engine/models"

	"github.com/google/uuid"
)

// This Method is use for the Adding the Jornal Lines
// This is service method
func (s *accountingServices) AddJournalLines(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	line accountingdtos.JournalLineRequestDto,
) error {

	if err := line.ValidateDto(); err != nil {
		return err
	}

	jouranlLine := &accounting_models.JournalLine{
		ID:             uuid.New(),
		JournalEntryID: line.JournalEntryID,
		AccountID:      line.AccountID,
		Debit:          line.Debit,
		Credit:         line.Credit,
		PartyID:        line.PartyID,
		CreatedAt:      line.CreatedAt,
	}

	if err := s.accRepo.CreateJournalLine(ctx, jouranlLine); err != nil {
		return err
	}

	return nil

}

