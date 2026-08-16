package kitchen_routes

import (
	kitchencontrollers "sharpnr-engine/property_infra/hospitality_infra/kitchen/controllers"

	"github.com/gin-gonic/gin"
)

func kitchenKOTRoutes(v1 *gin.RouterGroup, kc *kitchencontrollers.KitchenController) {

	kotRoutes := v1.Group("/kot/")
	kotRoutes.GET("orders", kc.GetKots)

	kotRoutes.PATCH("/:kotID/status", kc.UpdateKotsStatus)

}
