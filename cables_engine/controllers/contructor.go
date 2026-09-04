package cablesEngineControllers

import cablesEngineService "sharpnr-engine/cables_engine/service"

type cablesEngineController struct {
	cablesService cablesEngineService.CablesEngineServiceInterface
}

func NewCablesEngineController(cablesService cablesEngineService.CablesEngineServiceInterface) *cablesEngineController {
	return &cablesEngineController{
		cablesService: cablesService,
	}
}
