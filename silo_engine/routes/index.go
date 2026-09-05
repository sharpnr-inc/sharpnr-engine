package siloEngineRoutes

import (
	"sharpnr-engine/app"
	siloEngineControllers "sharpnr-engine/silo_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterSiloEngineRoutes(r *gin.RouterGroup, modules app.ModulesProvider) {

	siloContr := siloEngineControllers.NewSiloEngineController(modules.SiloEngineService())

	silo := r.Group("/silo")
	{
		// Is Silo on for this org, and is the account live? The client asks this
		// before anything else, so it answers serviceStarted=false rather than
		// 404 for an org that never started the service.
		//
		// Provisioning lives on the orchestrator (POST /orchestrator/silo/start)
		// — it is an admin action on the org, not a Silo product call, which is
		// why it is not mirrored here.
		silo.GET("/status", siloContr.SiloAccountStatusHandler)
	}

}
