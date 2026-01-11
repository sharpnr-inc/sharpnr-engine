package hospitalityinventory

import (
	"sharpnr-engine/app"

	"github.com/gin-gonic/gin"
)

func HospitalityRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	inventoryRoutes := v1.Group("/inventory")

	inventoryRoutes.GET("/get")
	inventoryRoutes.GET("/getItem/:itemId")
}
