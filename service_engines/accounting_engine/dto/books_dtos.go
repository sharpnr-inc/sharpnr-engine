package accountingdtos

import (
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
