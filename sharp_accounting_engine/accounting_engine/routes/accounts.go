package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/sharp_accounting_engine/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerAccAccountsRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	accr := ar.Group("/accounts")
	{
		accr.GET("/")
		accr.GET("/:accountID/ledger", ac.GetAccountLedger)
		accr.POST("/create", ac.CreateAccount)
	}
}
