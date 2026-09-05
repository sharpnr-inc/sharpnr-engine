package folioEngineControllers

import folioEngineService "sharpnr-engine/folio_engine/service"

type folioController struct {
	folioService folioEngineService.FolioEngineServiceInterface
}

func NewFolioEngineController(folioService folioEngineService.FolioEngineServiceInterface) *folioController {
	return &folioController{
		folioService: folioService,
	}
}
