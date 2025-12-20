package kitchen

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	kitchencontrollers "sharpnr-engine/modules/hospitality/kitchen/kitchenControllers"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func KitchenRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	//
	newRestaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	newBarRepo := barrepo.NewBarRepository(db)
	newHotelRepo := hotelrepo.NewHotelRepository(db)
	newMenuRepo := menurepo.NewHospitalityMenuRepository(db)

	newKitchenService := kitchenservices.NewKitchenService(newKitchenRepo, newRestaurantRepo, newBarRepo, newHotelRepo, newMenuRepo)
	newKitchenController := kitchencontrollers.NewKitchenController(newKitchenService)

	kitchenRoutes := v1.Group("/kitchen")

	// 🧾 Regular REST APIs
	kitchenRoutes.GET("/get-orders", newKitchenController.GetOrders)
	kitchenRoutes.POST("/update-orders", newKitchenController.UpdateOrder)

	// 🔥 Real-time WebSocket route
	kitchenRoutes.GET("/ws/:kitchenID", newKitchenController.KitchenWebSocket)
}
