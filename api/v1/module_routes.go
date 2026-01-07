package v1

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/base"
	"sharpnr-engine/modules/hospitality"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(
	v1 *gin.RouterGroup,
	db *gorm.DB,
	sere *sere.SERE,
	modules app.ModulesProvider,
) {
	base.RegisterBaseModuleRoutes(v1, db, modules)
	hospitality.RegisterHotelModuleRoutes(v1, db, sere, modules)
}
