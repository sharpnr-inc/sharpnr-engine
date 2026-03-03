package restaurant_routes

import (
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func restaurantOrdersRoutes(
	resRoutes *gin.RouterGroup,
	restaurantController restaurantcontrollers.RestaurantController) {

	// ror := resRoutes.Group("/orders")

}
