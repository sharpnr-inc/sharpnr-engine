package v1

import (
	"sharpnr-engine/app"
	accounting_engine_routes "sharpnr-engine/sharp_services/accounting_engine/routes"

	"sharpnr-engine/modules/hospitality"
	"sharpnr-engine/modules/hr"

	"github.com/gin-gonic/gin"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {
	// base.RegisterBaseModuleRoutes(v1, modules)
	hospitality.RegisterHotelModuleRoutes(v1, modules)
	accounting_engine_routes.RegisterAccountingRoutes(v1, modules)
	hr.RegisterHRRoutes(v1, modules)
}
