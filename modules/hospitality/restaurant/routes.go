package restaurant

import (
	"sharpnr-engine/app"
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {

	restaurantService := modules.Hospitality().RestaurantService()
	newRestaurantController := restaurantcontrollers.NewRestaurantController(restaurantService)

	resRoutes := v1.Group("/restaurant/:restaurantID")
	resRoutes.GET("/layout", newRestaurantController.GetRestaurantLayout)

	resRoutes.GET("/:tableID/dineIn-session", newRestaurantController.CreateDineSession)
	//dine In session Related
	resRoutes.GET("/dining-session/:dineSessionID/orders", newRestaurantController.GetDineInSessionOrders)
	resRoutes.POST("/dining-session/:dineSessionID/add-order", newRestaurantController.AddOrderInDiningSession)

}
