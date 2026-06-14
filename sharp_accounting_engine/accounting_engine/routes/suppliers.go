package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/sharp_accounting_engine/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerAccountingSuppliersRoutes(r *gin.RouterGroup, c *accounting_controllers.AccountingController) {
	accSupp := r.Group("/legal-entity-suppliers")
	{
		accSupp.GET("", c.ListAllLegalEntitySuppliersHandlers)
	}
}
