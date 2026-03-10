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
    ID          uuid.UUID `json:"id"`
    GroupName   string    `json:"group_name"`
    Type        string    `json:"type"`
    LastUpdated time.Time `json:"last_updated"`
}