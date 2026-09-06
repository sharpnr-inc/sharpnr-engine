package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerSiloOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	silo := r.Group("/silo")
	{
		// Is Silo on for this org, and is the account live? The admin shell asks
		// this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		silo.GET("", co.GetSiloAccountStatusHandler)

		// Provisioning lives here rather than on the Silo module routes: it is an
		// admin action on the organization, not a Silo product call.
		silo.POST("/start", co.StartSiloServiceHandler)
	}

}
