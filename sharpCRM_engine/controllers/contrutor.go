package sharpCrmControllers

import sharpCrmEngineService "sharpnr-engine/sharpCRM_engine/service"

type controllers struct {
	sharpCrmService sharpCrmEngineService.SharpCrmEngineServiceInterface
}

func NewSharpCrmControllers(sharpCrmService sharpCrmEngineService.SharpCrmEngineServiceInterface) *controllers {
	return &controllers{
		sharpCrmService: sharpCrmService,
	}
}
