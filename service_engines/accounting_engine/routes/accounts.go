package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerAccAccountsRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	accr := ar.Group("/accounts")
	{
		accr.GET("/")
	}
}
