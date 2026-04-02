package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerInventoryRoutes(r *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {
	invAcc_routes := r.Group("/inventory-account")
	{
		invAcc_routes.GET("", c.ListAllInventoryAccountHandler)
		invAcc_routes.POST("/add", c.NewInventoryAccountHandler)

		inv_routes := invAcc_routes.Group("/:inventoryAccountID/inventory/")
		{
			inv_routes.GET("details", c.ListAllInventoriesOfInvetoryAccountHandler)
			inv_routes.POST("add")
		}

	}
}
