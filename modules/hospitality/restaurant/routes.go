package restaurant

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestaurantRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	resRoutes := v1.Group("/restaurant")

	resRoutes.GET("/add-table")
}
