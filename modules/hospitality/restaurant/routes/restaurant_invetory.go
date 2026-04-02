package restaurant_routes

import (
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func registerRestaurantInventoryRoutes(
	resRoutes *gin.RouterGroup,
	restaurantController restaurantcontrollers.RestaurantController,
) {

	resInv := resRoutes.Group("/inventory")
	{
		resInv.GET("", restaurantController.GetAllInventoryAccounts)
	}

}
