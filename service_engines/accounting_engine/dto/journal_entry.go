package accountingdtos

import (
	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type JournalEntryRequestDto struct {
	AccountingBookID uuid.UUID
	FiscalYearID     uuid.UUID
	EntryDate        time.Time
	Reference        string
	Narration        string
	EntryType        accountingconstants.JournalEntryType
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type JournalLineRequestDto struct {
	JournalEntryID uuid.UUID
	AccountID uuid.UUID
	Debit  float64
	Credit float64
	PartyID *uuid.UUID
	CreatedAt time.Time
}
