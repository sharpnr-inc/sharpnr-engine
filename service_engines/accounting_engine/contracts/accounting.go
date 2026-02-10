package accounting_engine_contracts

import "github.com/google/uuid"

type DraftAccountingSetupInput struct {
	OrgID               uuid.UUID
	LegalEntityID       uuid.UUID
	Name_of_legalEntity string
}
