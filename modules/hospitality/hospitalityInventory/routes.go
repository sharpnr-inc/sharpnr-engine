package hospitalityinventory

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HospitalityRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	inventoryRoutes := v1.Group("/inventory")

	inventoryRoutes.GET("/get")
	inventoryRoutes.GET("/getItem/:itemId")
}
