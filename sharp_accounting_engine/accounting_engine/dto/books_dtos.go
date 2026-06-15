package accountingdtos

import (
	accountingconstants "sharpnr-engine/sharp_accounting_engine/accounting_engine/constants"
	"time"

	"github.com/google/uuid"
)

type AccountingBook_Response_dto struct {
	ID     string                       `json:"id"`
	Name   string                       `json:"name"`
	Status string                       `json:"status"`
	Groups []AccountingGroupResponseDto `json:"groups"`
}

type AccountingGroupResponseDto struct {
	ID          uuid.UUID             `json:"id"`
	GroupName   string                `json:"group_name"`
	Type        string                `json:"type"`
	LastUpdated time.Time             `json:"last_updated"`
	Accounts    []AccountsResponseDto `json:"accounts"`
}

type AccountsResponseDto struct {
	ID          uuid.UUID `json:"id"`
	AccountName string    `json:"account_name"`
	Code        string    `json:"code"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
}

type FiscalYearResponseDto struct {
	ID        uuid.UUID                          `json:"id"`
	Name      string                             `json:"name"`
	StartDate time.Time                          `json:"start_date"`
	EndDate   time.Time                          `json:"end_date"`
	Status    accountingconstants.FiscalYearType `json:"status"`
	UpdatedAt time.Time                          `json:"updated_at"`
}

type AccountingBookOfLegalEntity_Response_dto struct {
	Name        string                                   `json:"name"`
	Status      accountingconstants.AccountingBookStatus `json:"status"`
	FiscalYears []FiscalYearResponseDto                  `json:"fiscal_years"`
}
