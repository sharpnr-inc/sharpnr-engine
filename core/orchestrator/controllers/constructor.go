package orchestrator_controllers

import orchestrator_service "sharpnr-engine/core/orchestrator/service"

type orchestratorController struct {
	orchestService orchestrator_service.OrchestratorServiceInterface
}

func NewOrchestratorController(
	orchestService orchestrator_service.OrchestratorServiceInterface,
) *orchestratorController {
	return &orchestratorController{
		orchestService: orchestService,
	}
}
