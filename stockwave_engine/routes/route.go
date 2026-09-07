package stockwaveEngineRoutes

import (
	"sharpnr-engine/app"
	stockwaveEngineControllers "sharpnr-engine/stockwave_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterStockwaveEngineRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {

	stockwaveContr := stockwaveEngineControllers.NewStockwaveEngineController(modules.StockwaveEngineService())

	stockwave := v1.Group("/stockwave")
	{
		// Has the org taken stockwave, and is the account live? The client asks
		// this before anything else, so it answers serviceStarted=false rather
		// than 404 for an org that never started the service.
		//
		// Provisioning lives on the orchestrator (POST
		// /orchestrator/stockwave/start) — it is an admin action on the org, not
		// a stockwave product call, which is why it is not mirrored here.
		stockwave.GET("/status", stockwaveContr.StockwaveAccountStatusHandler)
	}

}
