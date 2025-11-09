package kitchen

import (
	kitchencontrollers "sharpnr-engine/modules/hospitality/kitchen/kitchenControllers"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KitchenRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	newKitchenService := kitchenservices.NewKitchenService(newKitchenRepo)
	newKitchenController := kitchencontrollers.NewKitchenController(newKitchenService)

	kitchenRoutes := v1.Group("/kitchen")

	// 🧾 Regular REST APIs
	kitchenRoutes.GET("/get-orders", newKitchenController.GetOrders)
	kitchenRoutes.POST("/update-orders", newKitchenController.UpdateOrder)

	// 🔥 Real-time WebSocket route
	kitchenRoutes.GET("/ws/:kitchenID", newKitchenController.KitchenWebSocket)
}
