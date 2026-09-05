package cortexEngineRoutes

import (
	"sharpnr-engine/app"
	cortexEngineControllers "sharpnr-engine/cortex_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCortexEngineRoutes(r *gin.RouterGroup, modules app.ModulesProvider) {

	cortexContr := cortexEngineControllers.NewCortexEngineController(modules.CortexEngineService())

	crtx := r.Group("/cortex")
	{
		// Is Cortex on for this org, and is the account live? The client asks
		// this before anything else, so it answers serviceStarted=false rather
		// than 404 for an org that never started the service.
		//
		// Provisioning lives on the orchestrator (POST
		// /orchestrator/cortex/start) — it is an admin action on the org, not a
		// Cortex product call, which is why it is not mirrored here.
		crtx.GET("/status", cortexContr.CortexAccountStatusHandler)
	}

}
