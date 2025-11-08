package hospitalitymenu

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MenuRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	menuRoutes := v1.Group("/hospitality-menu")

	menuRoutes.POST("/add/item")
	menuRoutes.POST("/remove/item")
}
