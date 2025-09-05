package hotel

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHotelModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	v1.GET("/hotel/getdashboard")
}
