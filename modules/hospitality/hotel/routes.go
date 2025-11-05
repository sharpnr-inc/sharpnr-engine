package hotel

import (
	hotelcontrollers "sharpnr-engine/modules/hospitality/hotel/controllers"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Hospitalityhotelroutes(v1 *gin.RouterGroup, db *gorm.DB) {
	newHotelRepo := hotelrepo.NewHotelRepository(db)
	newHotelService := hotelservices.NewHotelService(newHotelRepo)
	newHotelController := hotelcontrollers.NewHotelController(newHotelService)

	hotel_routes := v1.Group("/hotel")

	hotel_routes.GET("/get-all-rooms")
	hotel_routes.POST("/create-rooms", newHotelController.CreateRoom)
	hotel_routes.POST("/book-room")
	hotel_routes.GET("/get-hotel")

}
