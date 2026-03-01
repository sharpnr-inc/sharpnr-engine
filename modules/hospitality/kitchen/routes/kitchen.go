package kitchen_routes

import (
	kitchencontrollers "sharpnr-engine/modules/hospitality/kitchen/controllers"

	"github.com/gin-gonic/gin"
)

func kitchengeneralRoutes(v1 *gin.RouterGroup, kc *kitchencontrollers.KitchenController) {

	kitchenGeneralRoutes := v1.Group("/general")
	kitchenGeneralRoutes.GET("", kc.GetKitchenDetails)
}
