package v1

import (
	"sharpnr-engine/modules/base"
	"sharpnr-engine/modules/hospitality"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB, sere *sere.SERE) {
	base.RegisterBaseModuleRoutes(v1, db)
	hospitality.RegisterHotelModuleRoutes(v1, db, sere)
}
