package kitchen_routes

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
	sse := modules.SseEngine()
	controller := kitchencontrollers.NewKitchenController(kitchenService, sse)

	kitchenRoutes := v1.Group("/kitchen/:kitchenID")

	kitchengeneralRoutes(kitchenRoutes, controller)
	kitchenKOTRoutes(kitchenRoutes, controller)
	kitchenConfigsAndSettingsRoutes(kitchenRoutes, controller)

}
