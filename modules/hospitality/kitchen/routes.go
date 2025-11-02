package kitchen

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KitchenRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	kitchenRoutes := v1.Group("/kitchen")

	kitchenRoutes.GET("/get-orders")
	kitchenRoutes.POST("/update-orders")
}
