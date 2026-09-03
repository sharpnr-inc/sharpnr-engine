package orchestrator_service

import (
	"context"

	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	cablesEngineService "sharpnr-engine/cables_engine/service"

	"github.com/google/uuid"
)

// cablesEngineOrchestrator implements CablesOrchestratorServiceInterface.
//
// Every method is a single delegated call, so there is no transaction here
// (rule 2 in doc.go) — the cables service owns its own consistency, and it
// already commits each write together with the audit entry describing it.
//
// ⚠ No guard.Authorize calls yet (doc.go rule 1): cables has no Sentinel
// resource registered, so there is no action id to check against — guarding
// against an unregistered type would deny every request rather than protect
// it. Starting the org's cables service and creating channels is currently
// open to any authenticated employee of the org; wire the guard here the day
// the resource lands, the way departmentOrchestrator does.
type cablesEngineOrchestrator struct {
	cablesEngineService cablesEngineService.CablesEngineServiceInterface
}

/* ---------------------------------- <***> --------------------------------- */

func (r *cablesEngineOrchestrator) GetCablesAccountStatus(ctx context.Context, orgID, emplID uuid.UUID) (cablesEngineDtos.CablesServiceResponseDto, error) {
	return r.cablesEngineService.GetCablesServiceDetails(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

func (r *cablesEngineOrchestrator) StartCablesService(ctx context.Context, orgID, emplID uuid.UUID, request cablesEngineDtos.StartCablesServiceRequestsDto) error {
	return r.cablesEngineService.StartCablesService(ctx, orgID, emplID, request)
}

/* ---------------------------------- <***> --------------------------------- */

func (r *cablesEngineOrchestrator) CreateCablesChannel(ctx context.Context, orgID, emplID uuid.UUID, request cablesEngineDtos.CreateChannelRequestDto) error {
	return r.cablesEngineService.AddNewChannel(ctx, orgID, emplID, request)
}

/* ---------------------------------- <***> --------------------------------- */

func (r *cablesEngineOrchestrator) ListCablesChannels(ctx context.Context, orgID, emplID uuid.UUID) ([]cablesEngineDtos.ChannelResponseDto, error) {
	return r.cablesEngineService.ListChannels(ctx, orgID, emplID)
}

/* ---------------------------------- <***> --------------------------------- */

// GetCablesChannelDetails takes the id from the path rather than the body, the
// same way the department orchestrator does: a payload naming a different
// channel must not be able to retarget the read.
func (r *cablesEngineOrchestrator) GetCablesChannelDetails(ctx context.Context, orgID, emplID, channelID uuid.UUID) (cablesEngineDtos.ChannelResponseDto, error) {
	return r.cablesEngineService.GetChannelDetails(ctx, orgID, emplID, channelID)
}

/* ---------------------------------- <***> --------------------------------- */
