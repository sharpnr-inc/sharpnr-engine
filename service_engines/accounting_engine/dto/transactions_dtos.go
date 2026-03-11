package accountingdtos

import (
	"time"

	"github.com/google/uuid"
)

type JournalEntriesDTO struct {
	ID        uuid.UUID     `json:"id"`
	EntryDate time.Time     `json:"entry_date"`
	Reference string        `json:"reference"`
	Narration string        `json:"narration"`
	Balance   float64       `json:"balance"`
	Lines     []JournalLine `json:"lines"`
}

type JournalLine struct {
	ID          uuid.UUID `json:"id"`
	AccountID   uuid.UUID `json:"account_id"`
	AccountName string    `json:"account_name"`
	Credit      *float64  `json:"credit"`
	Debit       *float64  `json:"debit"`
}

type LedgerPaginationResponse struct {
	Data       []LedgerEntryDTO `json:"data"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalRows  int64            `json:"total_rows"`
	TotalPages int              `json:"total_pages"`
}

type LedgerEntryDTO struct {
	ID        uuid.UUID `json:"id"`
	EntryDate time.Time `json:"entry_date"`
	Reference string    `json:"reference"`
	Narration string    `json:"narration"`
	Debit     float64   `json:"debit"`
	Credit    float64   `json:"credit"`
	Balance   float64   `json:"balance"`
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

type JournalLineDTO struct {
	AccountID   uuid.UUID `json:"account_id"`
	AccountName string    `json:"account_name"`
	AccountCode string    `json:"account_code"`

	Debit  float64 `json:"debit"`
	Credit float64 `json:"credit"`
}
