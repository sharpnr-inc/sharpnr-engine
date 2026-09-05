package siloEngineControllers

import siloEngineService "sharpnr-engine/silo_engine/service"

type siloController struct {
	siloService siloEngineService.SiloEngineServiceInterface
}

func NewSiloEngineController(siloService siloEngineService.SiloEngineServiceInterface) *siloController {
	return &siloController{
		siloService: siloService,
	}
}
