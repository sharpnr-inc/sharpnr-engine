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

	kitchenRoutes := v1.Group("/kitchen/:kitchenID")
	kitchenRoutes.GET("/get-orders", newKitchenController.GetKots)

}
