package v1mobile

import (
	"sharpnr-engine/app"
	cablesEngineRoutes "sharpnr-engine/cables_engine/routes"
	hospitality_routes "sharpnr-engine/property_business_ops/hospitality/routes"

	"github.com/gin-gonic/gin"
)

func RegisterMobilesModulesRoutes(mobileV1 *gin.RouterGroup, modules app.ModulesProvider) {
	hospitality_routes.RegisterHotelModuleRoutes(mobileV1, modules)
	cablesEngineRoutes.ResgisterCablesEngineRoutes(mobileV1, modules)
}
