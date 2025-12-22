package accountingdtos

import "time"

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
}
