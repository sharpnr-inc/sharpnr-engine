package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerInventoryRoutes(r *gin.RouterGroup, c *accounting_controllers.AccountingController) {
	invr := r.Group("/inventories")
	{
		invr.GET("", c.GetAllInventory)

	}
}
