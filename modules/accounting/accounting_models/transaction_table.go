package accountingmodels

import "github.com/google/uuid"

type AccountingTransaction struct {
	ID             uuid.UUID
	Description    string
	Credit         bool
	Debit          bool
	OrganizationID uuid.UUID
	Date           string
}
