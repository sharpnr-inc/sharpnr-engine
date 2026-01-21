package accounting_engine_routes

import (
	"sharpnr-engine/app"
	accounting_controllers "sharpnr-engine/sharp_services/accounting_engine/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAccountingRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {

	accountingService := modules.AccountingService()
	accController := accounting_controllers.NewAccountingController(accountingService)

	accRoutes := v1.Group("/accounting")
	{
		// -------------------------
		// Accounting Books
		// -------------------------
		books := accRoutes.Group("/books")
		{
			books.GET("", accController.GetAccountingBooks)
			books.GET("/:bookId", accController.CreateTest)

			// -------------------------
			// Account Groups
			// -------------------------
			books.GET("/:bookId/groups", accController.GetAccountBookGroups)
			books.GET("/:bookId/groups/:groupId", accController.CreateTest)

			// -------------------------
			// Accounts
			// -------------------------
			books.GET("/:bookId/accounts", accController.GetAccounts)
			books.GET("/:bookId/accounts/:accountId", accController.CreateTest)
			books.POST("/:bookId/create-account", accController.CreateAccount)
		}

		// -------------------------
		// Account Mapping (Posting Rules)
		// -------------------------
		mappings := accRoutes.Group("/mappings")
		{
			mappings.GET("", accController.GetAllAccountMapping)
			mappings.POST("", accController.CreateAccountMapping)
		}

		// -------------------------
		// Journal Entries
		// -------------------------
		journal := accRoutes.Group("/journal")
		{
			journal.GET("", accController.CreateTest)
			journal.GET("/:entryId", accController.CreateTest)
		}

		// -------------------------
		// Reports (future-ready)
		// -------------------------
		reports := accRoutes.Group("/reports")
		{
			reports.GET("/trial-balance", accController.CreateTest)
			reports.GET("/profit-loss", accController.CreateTest)
			reports.GET("/balance-sheet", accController.CreateTest)
		}

		// Transactions

		tranx := accRoutes.Group("/transactions")
		{
			tranx.GET("", accController.GetTransactions)
			tranx.GET("/entry/:entryId", accController.GetTransactionDetails)
		}
	}
}
