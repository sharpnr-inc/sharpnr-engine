package bar

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BarRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	barRoute := v1.Group("/bar")

	barRoute.GET("/create-tables")
}
