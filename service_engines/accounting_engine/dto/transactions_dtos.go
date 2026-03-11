package accountingdtos

import (
	"time"

	"github.com/google/uuid"
)

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

type JournalEntryDetailsDTO struct {
	ID        uuid.UUID           `json:"id"`
	EntryDate time.Time           `json:"entry_date"`
	Reference string              `json:"reference"`
	Narration string              `json:"narration"`
	Lines     []JournalLineDetail `json:"lines"`
}

type JournalLineDetail struct {
	ID          uuid.UUID `json:"id"`
	AccountID   uuid.UUID `json:"account_id"`
	AccountName string    `json:"account_name"`
	Debit       float64   `json:"debit"`
	Credit      float64   `json:"credit"`
}
