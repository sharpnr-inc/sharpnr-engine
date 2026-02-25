package accountingdtos

import (
	"time"

	"github.com/google/uuid"
)

type LedgerLineDTO struct {
	Date      time.Time `json:"date"`
	Reference string    `json:"reference"`
	Narration string    `json:"narration"`

	Debit  float64 `json:"debit"`
	Credit float64 `json:"credit"`

	Balance float64 `json:"balance"`
}

type LedgerResponseDTO struct {
	AccountID   uuid.UUID `json:"account_id"`
	AccountName string    `json:"account_name"`
	AccountCode string    `json:"account_code"`

	OpeningBalance float64 `json:"opening_balance"`

	Lines []LedgerLineDTO `json:"lines"`

	ClosingBalance float64 `json:"closing_balance"`
}
