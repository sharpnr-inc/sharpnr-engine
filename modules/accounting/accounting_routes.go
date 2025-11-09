package accounting

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAccountingRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	accRoutes := v1.Group("/accounting")

	// This will fetch the transaction
	// it will have five Filters 1). All 2).Receipts 3).Payments 4).Transfers 5).Journal Entries
	accRoutes.GET("/get/transactions")
	accRoutes.GET("/get/sales")

	accRoutes.GET("/get/purchases")
	accRoutes.GET("/get/ledger")

	accRoutes.GET("/get/pnl")
	accRoutes.GET("/get/balanceSheet")
	accRoutes.GET("/get/cashFlows")
	accRoutes.GET("/get/taxSummary")
	accRoutes.GET("/get/reports/ledger")

}
