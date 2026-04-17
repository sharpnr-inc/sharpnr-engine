package restaurant_routes

import (
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func restsutantDineInSessionRoutes(
	resRoutes *gin.RouterGroup,
	restaurantController restaurantcontrollers.RestaurantController) {

	/* ---------------------------- <Tables Related> ---------------------------- */
	resRoutes.GET("/:tableID/dineIn-session", restaurantController.CreateDineSession)
	/* ---------------------------- <Dining Session> ---------------------------- */
	resRoutes.GET("/dining-session/:dineSessionID/orders", restaurantController.GetDineInSessionOrders)
	resRoutes.POST("/dining-session/:dineSessionID/add-order", restaurantController.AddOrderInDiningSession)
	resRoutes.POST("/dining-session/:dineSessionID/close-session", restaurantController.CloseDineInSession)
}
