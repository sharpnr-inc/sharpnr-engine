package orchestrator_service

import (
	"context"

	hrmsEngineDtos "sharpnr-engine/hrms_engine/dtos"
	hrmsEngineService "sharpnr-engine/hrms_engine/service"

	"github.com/google/uuid"
)

// hrmsOrchestratorService implements HrmsOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the HRMS service owns its own consistency, and it already
// commits the account, its settings row and the audit entry describing them in
// one transaction.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): HRMS has no Sentinel resource
// registered, so there is no action id to check against — guarding against an
// unregistered type would deny every request rather than protect it. Starting
// the org's HRMS service is currently open to any authenticated employee of the
// org; wire the guard here the day that registration lands, the way
// departmentOrchestrator does.
type hrmsOrchestratorService struct {
	hrmsService hrmsEngineService.HrmsEngineServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *hrmsOrchestratorService) GetHrmsAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (hrmsEngineDtos.HrmsAccountStatusResponse, error) {
	return s.hrmsService.GetHrmsAccountStatus(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *hrmsOrchestratorService) StartHrmsService(ctx context.Context, orgID, emplID uuid.UUID, request hrmsEngineDtos.StartHrmsServiceRequestDto) error {
	return s.hrmsService.StartHrmsServiceForOrg(ctx, orgID, emplID, request)
}
