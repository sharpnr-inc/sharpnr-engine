package base

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/accounting"
	"sharpnr-engine/modules/hr"

	"github.com/gin-gonic/gin"
)

// This is the Routes of All Base Module
func RegisterBaseModuleRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {
	accounting.RegisterAccountingRoutes(v1, modules)
	hr.RegisterHRRoutes(v1, modules)
}
