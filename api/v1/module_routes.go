package v1

import (
	"sharpnr-engine/app"
	accounting_engine_routes "sharpnr-engine/service_engines/accounting_engine/routes"

	hospitality_routes "sharpnr-engine/modules/hospitality/routes"
	hr_routes "sharpnr-engine/modules/hr/routes"

	"github.com/gin-gonic/gin"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {
	// base.RegisterBaseModuleRoutes(v1, modules)
	hospitality_routes.RegisterHotelModuleRoutes(v1, modules)
	accounting_engine_routes.RegisterAccountingRoutes(v1, modules)
	hr_routes.RegisterHRRoutes(v1, modules)
}
