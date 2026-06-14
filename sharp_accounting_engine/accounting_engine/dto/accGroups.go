package accountingdtos

import (
	accountingconstants "sharpnr-engine/sharp_accounting_engine/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type GetAccountingGroupOfBook_Response_dto struct {
	AccountGroups []AccountingGroup_Response_dto `json:"account_groups"`
}

type AccountingGroup_Response_dto struct {
	ID string `json:"id"`

	Name string `json:"name"`

	AccountingBookID string `json:"accounting_book_id"`

	Type string `json:"type"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	Accounts []AccountsResponse_dto `json:"accounts"`
}

type AccountsResponse_dto struct {
	ID   string
	Name string
}

/* ---------------------------------- <***> --------------------------------- */

type AccountResponseDto struct {
	ID               uuid.UUID `json:"id"`
	AccountGroupID   uuid.UUID `json:"account_group_id"`
	AccountingBookID uuid.UUID `json:"accounting_book_id"`

	Name string `json:"name"`
	Code string `json:"code"`

	OpeningBalance float64 `json:"opening_balance"`
	OpeningIsDebit bool    `json:"opening_is_debit"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountGroupBalanceResponseDto struct {
	ID               uuid.UUID `json:"id"`
	AccountingBookID uuid.UUID `json:"accounting_book_id"`

	Name string                               `json:"name"`
	Type accountingconstants.AccountGroupType `json:"type"`

	TotalDebitBalance  float64 `json:"total_debit_balance"`
	TotalCreditBalance float64 `json:"total_credit_balance"`

	NetBalance     float64 `json:"net_balance"`
	NetBalanceType string  `json:"net_balance_type"` // DEBIT / CREDIT

	Accounts []AccountResponseDto `json:"accounts"`
}
