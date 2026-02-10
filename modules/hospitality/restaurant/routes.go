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

}
