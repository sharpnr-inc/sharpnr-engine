package accounting_service

import (
	"context"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"

	"github.com/google/uuid"
)

func (s *accountingServices) CreateJournalLine(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
	journalLineRequest accountingdtos.JournalLineExternalRequestDto,
) (journalLineID uuid.UUID, err error) {
	panic("un")
}
