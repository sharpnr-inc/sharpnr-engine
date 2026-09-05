package cortexEngineControllers

import cortexEngineService "sharpnr-engine/cortex_engine/service"

type cortexController struct {
	cortexService cortexEngineService.CortexEngineServiceInterface
}

// NewCortexEngineController builds the module's controllers. (It returns a
// controller, not a service — the constructor was named NewCortexEngineService
// by copy-paste before anything called it.)
func NewCortexEngineController(cortexService cortexEngineService.CortexEngineServiceInterface) *cortexController {
	return &cortexController{
		cortexService: cortexService,
	}
}
