package accounting

import (
	"sharpnr-engine/modules/accounting/accounting_controllers"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/modules/accounting/accounting_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAccountingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	accRepo := accounting_repo.NewAccountingRepository(db)
	accService := accounting_service.NewAccountingService(accRepo)
	accController := accounting_controllers.NewAccountingController(accService)

	accRoutes := v1.Group("/accounting")
	{
		// -------------------------
		// Accounting Books
		// -------------------------
		books := accRoutes.Group("/books")
		{
			books.GET("", accController.CreateTest)
			books.GET("/:bookId", accController.CreateTest)

			// -------------------------
			// Account Groups
			// -------------------------
			books.GET("/:bookId/groups", accController.CreateTest)
			books.GET("/:bookId/groups/:groupId", accController.CreateTest)

			// -------------------------
			// Accounts
			// -------------------------
			books.GET("/:bookId/accounts", accController.CreateTest)
			books.GET("/:bookId/accounts/:accountId", accController.CreateTest)
		}

		// -------------------------
		// Account Mapping (Posting Rules)
		// -------------------------
		mappings := accRoutes.Group("/mappings")
		{
			mappings.GET("", accController.CreateTest)
			mappings.POST("", accController.CreateTest)
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
	}
}
