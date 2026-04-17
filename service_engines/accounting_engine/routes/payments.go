package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerAccountingPaymentsRoutes(r *gin.RouterGroup, c *accounting_controllers.AccountingController) {

	ap := r.Group("/payments")
	{
		ap.GET("", c.GetAllPaymentOrders)
	}

}
