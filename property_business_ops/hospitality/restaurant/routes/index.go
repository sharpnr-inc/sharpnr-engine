package restaurant_routes

import (
	"sharpnr-engine/app"
	restaurantcontrollers "sharpnr-engine/property_business_ops/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider) {
	restaurantService := modules.Hospitality().RestaurantService()
	newRestaurantController := restaurantcontrollers.NewRestaurantController(restaurantService)

	resRoutes := v1.Group("/restaurant/:restaurantID")

	restaurantGenralRoutes(resRoutes, *newRestaurantController)
	restaurantConfigsRoutes(resRoutes, *newRestaurantController)
	restsutantDineInSessionRoutes(resRoutes, *newRestaurantController)
	restaurantOrdersRoutes(resRoutes, *newRestaurantController)
	registerRestaurantInventoryRoutes(resRoutes, *newRestaurantController)

}
