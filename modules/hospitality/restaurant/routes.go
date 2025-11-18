package restaurant

import (
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestaurantRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	newRestaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	newRestaurantService := restaurantservices.NewRestaurantRepository(newRestaurantRepo)
	newRestaurantController := restaurantcontrollers.NewRestaurantController(newRestaurantService)

	resRoutes := v1.Group("/restaurant")

	resRoutes.GET("/add-table")
	resRoutes.GET("/get-details", newRestaurantController.GetRestaurantDetails)
	resRoutes.POST("/update", newRestaurantController.UpdateRestaurant)
}
