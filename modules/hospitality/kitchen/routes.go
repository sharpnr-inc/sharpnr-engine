package kitchen

import (
	"sharpnr-engine/app"

	"github.com/gin-gonic/gin"
)

func KitchenRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {
	// kitchenService := modules.Hospitality().KitchenService()
	// newKitchenController := kitchencontrollers.NewKitchenController(kitchenService)

	// kitchenRoutes := v1.Group("/kitchen")

	// // 🧾 Regular REST APIs
	// kitchenRoutes.GET("/get-orders")
	// kitchenRoutes.POST("/update-orders")

}
