package accounting_engine_routes

import (
	"sharpnr-engine/app"
	accounting_controllers "sharpnr-engine/service_engines/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountingRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {

	accountingService := modules.AccountingService()
	accController := accounting_controllers.NewAccountingController(accountingService)

	accRoutes := v1.Group("/accounting/:legalEntityID/")

	RegisterAccBooksRoutes(accRoutes, accController)
	RegisterAccGroupsRoutes(accRoutes, accController)
	registerJournalEntryRoutes(accRoutes, accController)
	registerAccountingPaymentsRoutes(accRoutes, accController)
	registerAccAccountsRoutes(accRoutes, accController)
	registerReportsRoutes(accRoutes, accController)
	registerAccountingSuppliersRoutes(accRoutes, accController)

}
