package accountingdtos

import "github.com/google/uuid"

type AccountMappingDTO struct {
	ID uuid.UUID `json:"id"`

	Module     string `json:"module"`
	SourceType string `json:"source_type"`
	SourceCode string `json:"source_code"`

	AccountID   uuid.UUID `json:"account_id"`
	AccountName string    `json:"account_name"`
	AccountCode string    `json:"account_code"`
}
