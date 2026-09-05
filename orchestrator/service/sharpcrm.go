package orchestrator_service

import (
	"context"
	sharpCrmEngineService "sharpnr-engine/sharpCRM_engine/service"
)

type sharpCrmOrchestratorService struct {
	crmService sharpCrmEngineService.SharpCrmEngineServiceInterface
}

func (s *sharpCrmOrchestratorService) GetCrmAccountStatus(ctx context.Context) error {
	panic("")
}
