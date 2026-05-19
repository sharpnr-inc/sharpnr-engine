package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerOrgRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	orgroutes := r.Group("/org")
	{
		orgroutes.GET("", c.OrgDetailsHandler)
	}

}
