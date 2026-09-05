package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCrmOrchestratorRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {
	crm := r.Group("/sharpcrm")
	{
		crm.GET("", co.GetCrmAccountStatusHandler)
	}

}
