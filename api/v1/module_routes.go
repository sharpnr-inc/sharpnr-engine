package v1

import (
	"sharpnr-engine/modules/base"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterModuleRoutes sets up all routes related to the specific business modules.
func RegisterModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	base.RegisterBaseModuleRoutes(v1, db)
}
