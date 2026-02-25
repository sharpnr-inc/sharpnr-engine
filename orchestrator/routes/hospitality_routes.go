package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerHospitalityRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	pr := r.Group("/hospitality")
	pr.GET("/restaurant", c.GetPropertiesOfRestaurant)
}
