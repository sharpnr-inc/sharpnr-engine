package accountingdtos

import (
	"time"

	"github.com/google/uuid"
)

type JournalEntryListDTO struct {
	ID uuid.UUID `json:"id"`

	EntryDate time.Time `json:"entry_date"`
	Reference string    `json:"reference"`
	Narration string    `json:"narration"`

	TotalDebit  float64 `json:"total_debit"`
	TotalCredit float64 `json:"total_credit"`

	CreatedAt time.Time `json:"created_at"`
}

type JournalLineDTO struct {
	AccountID   uuid.UUID `json:"account_id"`
	AccountName string    `json:"account_name"`
	AccountCode string    `json:"account_code"`

	Debit  float64 `json:"debit"`
	Credit float64 `json:"credit"`
}

type JournalEntryDetailDTO struct {
	ID uuid.UUID `json:"id"`

	EntryDate time.Time `json:"entry_date"`
	Reference string    `json:"reference"`
	Narration string    `json:"narration"`

	Lines []JournalLineDTO `json:"lines"`

	TotalDebit  float64 `json:"total_debit"`
	TotalCredit float64 `json:"total_credit"`
}
