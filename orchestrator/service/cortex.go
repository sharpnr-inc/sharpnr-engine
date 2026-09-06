package orchestrator_service

import (
	"context"

	cortexEngineDtos "sharpnr-engine/cortex_engine/dtos"
	cortexEngineService "sharpnr-engine/cortex_engine/service"

	"github.com/google/uuid"
)

// cortexOrchestratorService implements CortexOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the Cortex service owns its own consistency, and it
// already commits the account, its settings row and the audit entry describing
// them in one transaction.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): Cortex has no Sentinel
// resource registered, so there is no action id to check against — guarding
// against an unregistered type would deny every request rather than protect it.
// Starting the org's Cortex service is currently open to any authenticated
// employee of the org; wire the guard here the day that registration lands, the
// way departmentOrchestrator does.
type cortexOrchestratorService struct {
	cortexService cortexEngineService.CortexEngineServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *cortexOrchestratorService) GetCortexAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (cortexEngineDtos.CortexAccountStatusResponse, error) {
	return s.cortexService.GetCortexAccountStatus(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *cortexOrchestratorService) StartCortexService(ctx context.Context, orgID, emplID uuid.UUID, request cortexEngineDtos.StartCortexServiceRequestDto) error {
	return s.cortexService.StartCortexServiceForOrg(ctx, orgID, emplID, request)
}
