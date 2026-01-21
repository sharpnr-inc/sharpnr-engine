package kitchen

import (
	"sharpnr-engine/app"
	kitchencontrollers "sharpnr-engine/modules/hospitality/kitchen/controllers"

	"github.com/gin-gonic/gin"
)

func KitchenRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {
	kitchenService := modules.Hospitality().KitchenService()
	newKitchenController := kitchencontrollers.NewKitchenController(kitchenService)

	kitchenRoutes := v1.Group("/kitchen")

	// 🧾 Regular REST APIs
	kitchenRoutes.GET("/get-orders", newKitchenController.GetOrders)
	kitchenRoutes.POST("/update-orders", newKitchenController.UpdateOrder)

	// 🔥 Real-time WebSocket route
	kitchenRoutes.GET("/ws/:kitchenID", newKitchenController.KitchenWebSocket)
}
