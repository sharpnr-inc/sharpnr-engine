package accountingdtos

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	AccountingBookID uuid.UUID `json:"accounting_book_id"` // must me uuid
	AccountGroupID   uuid.UUID `json:"account_group_id"`
	Name             string    `json:"name"`
	Code             string    `json:"code"`
	OpeningBalance   float64   `json:"opening_balance"`
	IsDebit          bool      `json:"is_debit"` // true = Dr, false = Cr
	IsDefault        bool      `json:"is_default"`
}

func (r *CreateAccountRequest) ValidateRequest() error {

	if r.AccountingBookID == uuid.Nil {
		return errors.New("accounting_book_id is required")
	}

	if r.AccountGroupID == uuid.Nil {
		return errors.New("account_group_id is required")
	}

	if r.Name == "" {
		return errors.New("account name is required")
	}

	if r.Code == "" {
		return errors.New("account code is required")
	}

	if r.OpeningBalance < 0 {
		return errors.New("opening balance cannot be negative")
	}

	return nil
}

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
