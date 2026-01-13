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

	resRoutes := v1.Group("/restaurant")

	resRoutes.POST("/add-table", newRestaurantController.CreateOrAddTables)
	resRoutes.POST("/update-tables/:id", newRestaurantController.EditTablesDetailsOfRestaurant)

	resRoutes.POST("/add-floor", newRestaurantController.CreateOrAddFloor)
	resRoutes.GET("/get-restaurant-layout", newRestaurantController.GetRestaurantLayout)
	resRoutes.GET("/getRestaurant-Gen-details", newRestaurantController.GetRestaurantDetails)
	resRoutes.POST("/update", newRestaurantController.UpdateRestaurant)

	/******************************
	 * Order Routes of Restaurant *
	 ******************************/

	sessionRoutes := resRoutes.Group("/session")
	{
		sessionRoutes.POST("/create", newRestaurantController.CreateSession)
		sessionRoutes.GET("/getbyTable/:id", newRestaurantController.GetBookingSessionById)
		// id is sessionId
		sessionRoutes.POST("/addItem/:id", newRestaurantController.AddOrdersInSession)
		// id is KitchenTicketId
		sessionRoutes.POST("/getOrder/:id", newRestaurantController.GetOrderStatus)

	}
}
