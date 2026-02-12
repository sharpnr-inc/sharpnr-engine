package accountingdtos

import "github.com/google/uuid"

type CreateAccountRequest struct {
	AccountingBookID uuid.UUID `json:"accounting_book_id"`
	AccountGroupID   uuid.UUID `json:"account_group_id"`

	Name string `json:"name"`
	Code string `json:"code"`

	OpeningBalance float64 `json:"opening_balance"`
	IsDebit        bool    `json:"is_debit"` // true = Dr, false = Cr

	IsDefault bool `json:"is_default"`
}
