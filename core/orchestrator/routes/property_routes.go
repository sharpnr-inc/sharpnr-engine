package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/core/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerPropertyRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	r.GET("/property", c.ListAllOrgProperty)
	r.GET("/property/:propertyID", c.GetPropertyDetails)
}
