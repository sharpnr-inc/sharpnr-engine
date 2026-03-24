package accountingdtos

import (
	"errors"
	accountingconstants "sharpnr-engine/service_engines/accounting_engine/constants"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Dto for the Creating Journal Entry
// will be used in the Service
type JournalEntryRequestDto struct {
	AccountingBookID uuid.UUID
	FiscalYearID     uuid.UUID
	EntryDate        time.Time
	Reference        string
	Narration        string
	EntryType        accountingconstants.JournalEntryType
	JournalLines     []JournalLineRequestDto
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// Validation Function for the Journal Request Dto
func (j *JournalEntryRequestDto) ValidateRequest() error {
	if j.AccountingBookID == uuid.Nil {
		return errors.New("accounting_book_id is required")
	}

	if j.FiscalYearID == uuid.Nil {
		return errors.New("fiscal_year_id is required")
	}

	if j.EntryDate.IsZero() {
		return errors.New("entry_date is required")
	}

	if strings.TrimSpace(j.Reference) == "" {
		return errors.New("reference is required")
	}

	if strings.TrimSpace(j.Narration) == "" {
		return errors.New("narration is required")
	}

	if !j.EntryType.IsValid() {
		return errors.New("invalid entry_type")
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

type JournalLineRequestDto struct {
	JournalEntryID uuid.UUID
	AccountID      uuid.UUID
	Debit          float64
	Credit         float64
	PartyID        *uuid.UUID
	CreatedAt      time.Time
}

func (jl *JournalLineRequestDto) ValidateDto() error {
	if jl.JournalEntryID == uuid.Nil {
		return errors.New("journal_entry_id is required")
	}

	if jl.AccountID == uuid.Nil {
		return errors.New("account_id is required")
	}

	if jl.Debit < 0 {
		return errors.New("debit cannot be negative")
	}

	if jl.Credit < 0 {
		return errors.New("credit cannot be negative")
	}

	if jl.Debit == 0 && jl.Credit == 0 {
		return errors.New("either debit or credit must be greater than zero")
	}

	if jl.Debit > 0 && jl.Credit > 0 {
		return errors.New("only one of debit or credit can be greater than zero")
	}

	return nil
}
