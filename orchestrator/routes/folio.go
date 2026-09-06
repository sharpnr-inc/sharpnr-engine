package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerFolioOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	folio := r.Group("/folio")
	{
		// Is Folio on for this org, and is the account live? The admin shell
		// asks this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		folio.GET("", co.GetFolioAccountStatusHandler)

		// Provisioning lives here rather than on the Folio module routes: it is
		// an admin action on the organization, not a Folio product call.
		folio.POST("/start", co.StartFolioServiceHandler)
	}

}
