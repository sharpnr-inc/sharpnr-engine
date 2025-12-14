package hotel

import (
	hotelcontrollers "sharpnr-engine/modules/hospitality/hotel/hotel_controllers"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/****************************************
 * These are the endpoint for the Hotel *
****************************************/

func Hospitalityhotelroutes(v1 *gin.RouterGroup, db *gorm.DB) {
	newHotelRepo := hotelrepo.NewHotelRepository(db)

	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	newHouseKeepingRepo := hotelrepo.NewHotelHouseKeepingRepository(db)
	newHotelService := hotelservices.NewHotelService(newHotelRepo, newHouseKeepingRepo, newKitchenRepo)

	newHotelController := hotelcontrollers.NewHotelController(newHotelService)

	hotel_routes := v1.Group("/hotel")

	hotel_routes.GET("/get-all-rooms")
	hotel_routes.POST("/create-rooms", newHotelController.CreateRoom)
	hotel_routes.POST("/create-floor", newHotelController.CreateFloorInHotel)
	hotel_routes.POST("/book-room/:roomId")
	hotel_routes.GET("/get-hotel", newHotelController.GetHotel)
	hotel_routes.GET("/overview", newHotelController.GetHotelOverview)
	hotel_routes.GET("/bookings", newHotelController.GetHotelsBookings)
	hotel_routes.GET("/request", newHotelController.GetRequestInHotels)
	hotel_routes.POST("/configure", newHotelController.ConfigureAndUpdate)

	// This is the Endpoint for Fetching the Hotel's frontdesk data
	// It include the Data Available, occupied rooms
	hotel_routes.GET("/frontdesk", newHotelController.GetHotelFrontDesk)

	// TODO For Now this is Routes is used to return the Hotel Overview Data
	// In future we will replace this Frontdesk endpoint
	hotel_routes.GET("/reception", newHotelController.GetHotelReception)

	// This is the endpoint for the changing the room status
	hotel_routes.GET("/room-status/:hotelId/:roomId/:status", newHotelController.UpdateRoomStatus)

	hotel_routes.POST("/create-booking", newHotelController.CreateBooking)
	hotel_routes.GET("/get-booking/:roomId", newHotelController.GetBookingDetails)

	houseKeeping := hotel_routes.Group("/houseKeeping")
	{
		houseKeeping.GET("/requests")
		houseKeeping.POST("")
	}

}
