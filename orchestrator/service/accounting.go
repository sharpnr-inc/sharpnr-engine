package orchestrator_service

import (
	"context"

	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accounting_service "sharpnr-engine/sharp_accounting_engine/accounting_engine/services"

	"github.com/google/uuid"
)

// accountingOrchestratorService implements AccountingOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the accounting service owns its own consistency, and it
// already commits the account, its settings row and the audit entry describing
// them in one transaction.
//
// This is the org-level switch only. The per-legal-entity accounting reads —
// books and account balances — live on legalEntityOrchestrator, because they are
// scoped to an entity rather than to the organization.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): the accounting engine has no
// Sentinel resource registered, so there is no action id to check against —
// guarding against an unregistered type would deny every request rather than
// protect it. Starting the org's accounting service is currently open to any
// authenticated employee of the org; wire the guard here the day that
// registration lands, the way departmentOrchestrator does.
type accountingOrchestratorService struct {
	accountingService accounting_service.AccountingServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *accountingOrchestratorService) GetAccountingAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (accountingdtos.AccountingAccountStatusResponse, error) {
	return s.accountingService.GetAccountingAccountStatus(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *accountingOrchestratorService) StartAccountingService(ctx context.Context, orgID, emplID uuid.UUID, request accountingdtos.StartAccountingServiceRequestDto) error {
	return s.accountingService.StartAccountingServiceForOrg(ctx, orgID, emplID, request)
}
