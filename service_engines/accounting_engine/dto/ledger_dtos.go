package accountingdtos

import (
	"time"

	"github.com/google/uuid"
)

type LedgerLineDto struct {
	EntryDate      time.Time
	Reference      string
	Narration      string
	EntryType      string
	Debit          float64
	Credit         float64
	RunningBalance float64
}

type AccountLedgerDto struct {
	AccountID      uuid.UUID
	AccountName    string
	AccountCode    string
	OpeningBalance float64 // from Account.OpeningBalance
	OpeningIsDebit bool    // from Account.OpeningIsDebit
	Lines          []LedgerLineDto
	ClosingBalance float64 // computed
}
