package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerUtilsRoutes(v1 *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	ur := v1.Group("/utils")
	{
		ur.GET("/buinsess_modules", c.BusinessOperationModulesHandler)
		ur.GET("/Infra_modules", c.InfraModulesHandler)
	}
}
