package accounting_models

import (
	accountingconstants "sharpnr-engine/sharp_services/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type JournalEntry struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	AccountingBookID uuid.UUID      `gorm:"type:uuid;not null;index"`
	AccountingBook   AccountingBook `gorm:"foreignKey:AccountingBookID"`

	FiscalYearID uuid.UUID  `gorm:"type:uuid;not null;index"`
	FiscalYear   FiscalYear `gorm:"foreignKey:FiscalYearID"`

	EntryDate time.Time `gorm:"not null"`

	Reference string // Invoice no, Payment ref, etc.
	Narration string // Human readable explanation

	EntryType accountingconstants.JournalEntryType
	// NORMAL, OPENING_BALANCE, ADJUSTMENT

	Lines []JournalLine `gorm:"foreignKey:JournalEntryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (JournalEntry) TableName() string {
	return "accounting_engine.journal_entry"
}
