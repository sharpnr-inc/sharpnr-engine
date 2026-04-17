package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerReportsRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	repr := ar.Group("/reports")
	{
		repr.GET("/profit-loss", ac.GetProfitAndLossReport)
	}
}
