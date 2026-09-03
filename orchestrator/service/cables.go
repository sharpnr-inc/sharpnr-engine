package orchestrator_service

import (
	cablesEngineService "sharpnr-engine/cables_engine/service"
)

type cablesEngineOrchestrator struct {
	cablesEngineService cablesEngineService.CablesEngineServiceInterface
}

func (r *cablesEngineOrchestrator) GetCablesAccountStatus() {

}
