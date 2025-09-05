package base

import (
	"sharpnr-engine/modules/base/accounting"
	"sharpnr-engine/modules/base/hr"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This is the Routes of All Base Module
func RegisterBaseModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	accounting.RegisterAccountingRoutes(v1, db)
	hr.RegisterHRRoutes(v1, db)
}
