package stockwaveEngineControllers

import stockwaveEngineServices "sharpnr-engine/stockwave_engine/services"

type stockwaveController struct {
	stockwaveService stockwaveEngineServices.StockwaveEngineServicesInterface
}

func NewStockwaveEngineController(stockwaveService stockwaveEngineServices.StockwaveEngineServicesInterface) *stockwaveController {
	return &stockwaveController{
		stockwaveService: stockwaveService,
	}
}
