package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/sharp_accounting_engine/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerJournalEntryRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	jer := ar.Group("/je")
	{
		jer.GET("/entries", ac.GetJournalEntries)
		jer.POST("/entry/create", ac.CreateJournalEntry) // used to create journal entry from external source (like payment service)

	}
}
