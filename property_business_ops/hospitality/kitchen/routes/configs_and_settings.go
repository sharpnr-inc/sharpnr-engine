package kitchen_routes

import (
	kitchencontrollers "sharpnr-engine/property_business_ops/hospitality/kitchen/controllers"

	"github.com/gin-gonic/gin"
)

func kitchenConfigsAndSettingsRoutes(v1 *gin.RouterGroup, kc *kitchencontrollers.KitchenController) {

	kitchenconfigRoutes := v1.Group("/configs")
	kitchenconfigRoutes.GET("") // Get all configs for a kitchen
	kitchenconfigRoutes.PATCH("/activate-dactivate/:kitchen_switch", kc.ActivateDeactivateKitchen)
}
