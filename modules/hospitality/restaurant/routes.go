package restaurant

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestaurantRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	newRestaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	newRestaurantService := restaurantservices.NewRestaurantRepository(newRestaurantRepo, newKitchenRepo)
	newRestaurantController := restaurantcontrollers.NewRestaurantController(newRestaurantService)

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
