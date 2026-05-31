package v1mobile

import (
	"sharpnr-engine/app"
	hospitality_routes "sharpnr-engine/business_ops/hospitality/routes"

	"github.com/gin-gonic/gin"
)

func RegisterMobilesModulesRoutes(mobileV1 *gin.RouterGroup, modules app.ModulesProvider) {
	hospitality_routes.RegisterHotelModuleRoutes(mobileV1, modules)
}
