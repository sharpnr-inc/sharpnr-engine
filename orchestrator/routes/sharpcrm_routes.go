package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCrmOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	crm := r.Group("/sharpcrm")
	{
		// Is SharpCRM on for this org, and is the account live? The admin shell
		// asks this before offering the start action, so it answers
		// serviceStarted=false rather than 404 for an org that never started it.
		crm.GET("", co.GetCrmAccountStatusHandler)

		// Provisioning lives here rather than on the CRM module routes: it is an
		// admin action on the organization, not a CRM product call.
		crm.POST("/start", co.StartCrmServiceHandler)
	}

}
