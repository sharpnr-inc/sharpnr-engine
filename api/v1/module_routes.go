package v1

import (
	"sharpnr-engine/app"
	payment_engine_routes "sharpnr-engine/payment_engine/routes"
	accounting_engine_routes "sharpnr-engine/service_engines/accounting_engine/routes"

	hospitality_routes "sharpnr-engine/business_ops/hospitality/routes"
	hr_routes "sharpnr-engine/sharp_hrms_engine/routes"

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
	payment_engine_routes.PaymentGatewaysRoutes(v1, modules)
}
