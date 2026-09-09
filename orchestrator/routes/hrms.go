package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerHrmsOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	// The group is /sharp-hr rather than /hrms so the orchestrator path matches
	// the module route and the client's own /sharp-hr URL.
	hr := r.Group("/sharp-hr")
	{
		// Is HRMS on for this org, and is the account live? The admin shell asks
		// this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		hr.GET("", co.GetHrmsAccountStatusHandler)

		// Provisioning lives here rather than on the HRMS module routes: it is an
		// admin action on the organization, not an HR product call.
		hr.POST("/start", co.StartHrmsServiceHandler)
	}

}
