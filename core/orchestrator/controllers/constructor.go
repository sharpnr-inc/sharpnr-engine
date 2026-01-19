package orchestrator_controllers

import orchestrator_service "sharpnr-engine/core/orchestrator/service"

type OrchestratorController struct {
	orchestService orchestrator_service.OrchestratorServiceInterface
}

func NewOrchestratorController(
	orchestService orchestrator_service.OrchestratorServiceInterface,
) *OrchestratorController {
	return &OrchestratorController{
		orchestService: orchestService,
	}
}
