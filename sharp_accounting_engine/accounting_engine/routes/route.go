package accounting_engine_routes

import (
	"sharpnr-engine/app"
	accounting_controllers "sharpnr-engine/sharp_accounting_engine/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountingRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {

	accountingService := modules.AccountingService()
	accController := accounting_controllers.NewAccountingController(accountingService)

	// GET /accounting/status — has the org taken the accounting service, and is
	// that account live? The client asks this before anything else, so it
	// answers serviceStarted=false rather than 404 for an org that never
	// started it. Static sibling of /accounting/:legalEntityID below, the same
	// way /accounting/e already is.
	//
	// Provisioning lives on the orchestrator (POST
	// /orchestrator/sharp-accounting/start) — it is an admin action on the org,
	// not an accounting product call, which is why it is not mirrored here.
	v1.GET("/accounting/status", accController.GetAccountingAccountStatusHandler)

	// GET /accounting/e — fetch all legal entities of the organization
	v1.GET("/accounting/e", accController.GetAllLegalEntities)

	// All entity-scoped routes live under /accounting/:legalEntityID
	accRoutes := v1.Group("/accounting/:legalEntityID")
	RegisterAccBooksRoutes(accRoutes, accController)
	RegisterAccGroupsRoutes(accRoutes, accController)
	registerJournalEntryRoutes(accRoutes, accController)
	registerAccountingPaymentsRoutes(accRoutes, accController)
	registerAccAccountsRoutes(accRoutes, accController)
	registerReportsRoutes(accRoutes, accController)
	registerAccountingSuppliersRoutes(accRoutes, accController)
	registerInventoryRoutes(accRoutes, accController)
}
