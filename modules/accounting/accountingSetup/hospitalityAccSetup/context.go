package hospitality_acc_setup

import (
	"sharpnr-engine/modules/accounting/accounting_repo"

	"github.com/google/uuid"
)

type AccSetupContext struct {
	OrganizationID   uuid.UUID
	OrganizationName string
	OrganizationType string
	Modules          []string
	AccountingRepo   accounting_repo.AccountingRepositoryInterface
}
