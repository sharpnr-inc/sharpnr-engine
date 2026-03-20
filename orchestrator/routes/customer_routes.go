package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerCustomerRoutes(rg *gin.RouterGroup, orc *orchestrator_controllers.OrchestratorController) {

	orgCust := rg.Group("/org/customers")
	{
		orgCust.POST("/add", orc.CreateCustomerHandler)
		orgCust.GET("/search", orc.SearchCustomerHandler)
	}
}
