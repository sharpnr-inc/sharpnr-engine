package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCortexOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	cortex := r.Group("/cortex")
	{
		// Is Cortex on for this org, and is the account live? The admin shell
		// asks this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		cortex.GET("", co.GetCortexAccountStatusHandler)

		// Provisioning lives here rather than on the Cortex module routes: it is
		// an admin action on the organization, not a Cortex product call.
		cortex.POST("/start", co.StartCortexServiceHandler)
	}

}
