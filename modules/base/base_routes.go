package base

import (
	"sharpnr-engine/middlewares"
	"sharpnr-engine/modules/accounting"
	"sharpnr-engine/modules/hr"
	"sharpnr-engine/pkg/jwtpack"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This is the Routes of All Base Module
func RegisterBaseModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	// 🔐 Protected base modules (Org scoped)
	protected := v1.Group("/")
	protected.Use(
		middlewares.RequireTokenType(jwtpack.TokenTypeAuth),
		middlewares.SetOrgUUID(),
	)
	accounting.RegisterAccountingRoutes(protected, db)
	hr.RegisterHRRoutes(protected, db)
}
