package accounting_models

import (
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

/*
Account Type    Debit (Left)    Credit (Right)
Assets          Increase (+)    Decrease (-)
Liabilities     Decrease (-)    Increase (+)
Equity          Decrease (-)    Increase (+)
Revenue         Decrease (-)    Increase (+)
Expenses        Increase (+)    Decrease (-)
*/

type JournalEntry struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccountingBookID uuid.UUID `gorm:"type:uuid;not null;index"`
	FiscalYearID     uuid.UUID `gorm:"type:uuid;not null;index"`

	EntryDate time.Time `gorm:"not null"`

	Reference string // Invoice no, Payment ref, etc.
	Narration string // Human readable explanation

	EntryType accountingconstants.JournalEntryType

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (JournalEntry) TableName() string {
	return "accounting_engine.journal_entry"
}
