package app

import (
	orchestrator_service "sharpnr-engine/core/orchestrator/service"
)

func bootstrapSharpnrOrchestrator() orchestrator_service.OrchestratorServiceInterface {
	return orchestrator_service.NewSharpnrOrchestratorService()
}
