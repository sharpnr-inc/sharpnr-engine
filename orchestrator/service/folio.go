package orchestrator_service

import (
	"context"

	folioEngineDtos "sharpnr-engine/folio_engine/dtos"
	folioEngineService "sharpnr-engine/folio_engine/service"

	"github.com/google/uuid"
)

// folioOrchestratorService implements FolioOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the Folio service owns its own consistency, and it
// already commits the account, its settings row and the audit entry describing
// them in one transaction.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): Folio has no Sentinel
// resource registered, so there is no action id to check against — guarding
// against an unregistered type would deny every request rather than protect it.
// Starting the org's Folio service is currently open to any authenticated
// employee of the org; wire the guard here the day that registration lands, the
// way departmentOrchestrator does.
type folioOrchestratorService struct {
	folioService folioEngineService.FolioEngineServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (s *folioOrchestratorService) GetFolioAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (folioEngineDtos.FolioAccountStatusResponse, error) {
	return s.folioService.GetFolioAccountStatus(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (s *folioOrchestratorService) StartFolioService(ctx context.Context, orgID, emplID uuid.UUID, request folioEngineDtos.StartFolioServiceRequestDto) error {
	return s.folioService.StartFolioServiceForOrg(ctx, orgID, emplID, request)
}
