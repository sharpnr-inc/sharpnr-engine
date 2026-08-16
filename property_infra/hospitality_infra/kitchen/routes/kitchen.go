package kitchen_routes

import (
	kitchencontrollers "sharpnr-engine/property_infra/hospitality_infra/kitchen/controllers"

	"github.com/gin-gonic/gin"
)

func kitchengeneralRoutes(v1 *gin.RouterGroup, kc *kitchencontrollers.KitchenController) {

	r := v1.Group("/general")
	r.GET("", kc.GetKitchenDetails)         // Retuns the general details of the kitchen like name, address, contact details, etc.
	r.GET("/configs", kc.GetKitchenConfigs) // fetches all the configs related to the kitchen
	r.PATCH("/update-general-details", kc.UpdateKitchenGenralDetails)

}
