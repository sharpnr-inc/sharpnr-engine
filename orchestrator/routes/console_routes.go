package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerConsoleRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	cr := r.Group("/console")
	registerPropertyRoutes(cr, c)
	registerLegalEntityRoutes(cr, c)
}
