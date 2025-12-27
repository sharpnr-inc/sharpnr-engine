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
