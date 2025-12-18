package hospitality_acc_setup

import (
	"sharpnr-engine/core/models"
	"sharpnr-engine/modules/accounting/accounting_repo"
)

type AccSetupContext struct {
	Org            *models.Organization
	AccountingRepo accounting_repo.AccountingRepositoryInterface
}
