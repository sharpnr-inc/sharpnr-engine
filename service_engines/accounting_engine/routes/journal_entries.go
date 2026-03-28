package accounting_engine_routes

import (
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func registerJournalEntryRoutes(ar *gin.RouterGroup, ac *accounting_controllers.AccountingController) {

	jer := ar.Group("/je")
	{
		jer.GET("/entries", ac.GetJournalEntries)
		jer.POST("/entry/create", ac.CreateJournalEntry)
		jer.GET("/entry/:entryID", ac.GetJornalDetails)
	}
}
