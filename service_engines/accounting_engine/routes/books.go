package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccBooksRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	br := ar.Group("/:accountingBookID/books")

	{
		br.GET("", ac.GetAccountingBooks)
	}
}
