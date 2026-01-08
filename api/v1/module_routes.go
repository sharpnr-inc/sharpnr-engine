package v1

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/base"
	"sharpnr-engine/modules/hospitality"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(
	v1 *gin.RouterGroup,
	sere *sere.SERE,
	modules app.ModulesProvider,
) {
	base.RegisterBaseModuleRoutes(v1, modules)
	hospitality.RegisterHotelModuleRoutes(v1, sere, modules)
}
