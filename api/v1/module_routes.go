package v1

import (
	"sharpnr-engine/app"
	cablesEngineRoutes "sharpnr-engine/cables_engine/routes"
	payment_engine_routes "sharpnr-engine/payment_engine/routes"
	sharpCrmEngineRoutes "sharpnr-engine/sharpCRM_engine/routes"
	accounting_engine_routes "sharpnr-engine/sharp_accounting_engine/accounting_engine/routes"

	hr_routes "sharpnr-engine/hrms_engine/routes"
	hospitality_routes "sharpnr-engine/property_business_ops/hospitality/routes"

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
	cablesEngineRoutes.ResgisterCablesEngineRoutes(v1, modules)
	sharpCrmEngineRoutes.RegisterSharpCrmEngine(v1, modules)

}
