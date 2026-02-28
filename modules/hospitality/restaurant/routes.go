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

	resRoutes.GET("/general-details", newRestaurantController.GetGeneralDetails)
	resRoutes.POST("/update-general-details", newRestaurantController.UpdateGeneralDetails)
	resRoutes.GET("/layout", newRestaurantController.GetRestaurantLayout)

	/* ---------------------------------- <***> --------------------------------- */
	resRoutes.POST("/configs/add-layout", newRestaurantController.AddLayout)

	/* ---------------------------- <Tables Related> ---------------------------- */
	resRoutes.GET("/:tableID/dineIn-session", newRestaurantController.CreateDineSession)
	/* ---------------------------- <Dining Session> ---------------------------- */
	resRoutes.GET("/dining-session/:dineSessionID/orders", newRestaurantController.GetDineInSessionOrders)
	resRoutes.POST("/dining-session/:dineSessionID/add-order", newRestaurantController.AddOrderInDiningSession)

}
