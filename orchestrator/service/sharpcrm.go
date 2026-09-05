package orchestrator_service

import (
	"context"

	sharpCrmEngineDtos "sharpnr-engine/sharpCRM_engine/dtos"
	sharpCrmEngineService "sharpnr-engine/sharpCRM_engine/service"

	"github.com/google/uuid"
)

// sharpCrmOrchestratorService implements SharpCrmOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the CRM service owns its own consistency, and it already
// commits the account, its settings row and the audit entry describing them in
// one transaction.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): SharpCRM has no Sentinel
// resource registered, so there is no action id to check against — guarding
// against an unregistered type would deny every request rather than protect it.
// Starting the org's CRM service is currently open to any authenticated
// employee of the org; wire the guard here the day that registration lands, the
// way departmentOrchestrator does.
type sharpCrmOrchestratorService struct {
	crmService sharpCrmEngineService.SharpCrmEngineServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sharpCrmOrchestratorService) GetCrmAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (sharpCrmEngineDtos.SharpCrmAccountStatusResponse, error) {
	return s.crmService.GetCrmAccountStatus(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sharpCrmOrchestratorService) StartCrmService(ctx context.Context, orgID, emplID uuid.UUID, request sharpCrmEngineDtos.StartCRMServiceRequestDto) error {
	return s.crmService.StartCrmServiceForOrg(ctx, orgID, emplID, request)
}
