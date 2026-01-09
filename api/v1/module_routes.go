package v1

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/base"
	"sharpnr-engine/modules/hospitality"

	"github.com/gin-gonic/gin"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {
	base.RegisterBaseModuleRoutes(v1, modules)
	hospitality.RegisterHotelModuleRoutes(v1, modules)
}
