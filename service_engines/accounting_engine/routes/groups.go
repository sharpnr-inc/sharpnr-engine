package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccGroupsRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	br := ar.Group("/:accountingBookID/group")
	{
		br.POST("/add", ac.GetAccountingBooks) // Add Accounting Group
		br.PATCH("/update") // Update Accounting Book

	}
}
