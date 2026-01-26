package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/core/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerConsoleRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	cr := r.Group("/console")

	cr.GET("/org-status", c.GetOrgStatus)
	cr.GET("/property", c.ListAllOrgProperty)
	cr.GET("/legal-entity")

}
