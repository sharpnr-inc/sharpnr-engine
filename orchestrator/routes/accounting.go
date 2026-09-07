package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerAccountingOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	// The group is /sharp-accounting rather than /accounting so the orchestrator
	// path matches the client's own /sharp-accounting route. The engine's own
	// product API stays on /accounting.
	acc := r.Group("/sharp-accounting")
	{
		// Has the org taken accounting, and is the account live? The admin shell
		// asks this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		acc.GET("", co.GetAccountingAccountStatusHandler)

		// Provisioning lives here rather than on the accounting module routes: it
		// is an admin action on the organization, not an accounting product call.
		acc.POST("/start", co.StartAccountingServiceHandler)
	}

}
