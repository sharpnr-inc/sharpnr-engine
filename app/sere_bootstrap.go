package app

import (
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/accounting/accounting_sere_handlers"
	"sharpnr-engine/sere"
)

func RegisterSereEventHandler(
	sere *sere.SERE,
	accountingRepo accounting_repo.AccountingRepositoryInterface,
) {
	accounting_sere_handlers.Register(sere, accountingRepo)
}
