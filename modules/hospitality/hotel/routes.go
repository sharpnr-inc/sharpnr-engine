package hotel

import (
	hotelcontrollers "sharpnr-engine/modules/hospitality/hotel/hotel_controllers"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/****************************************
 * These are the endpoint for the Hotel *
****************************************/

func Hospitalityhotelroutes(v1 *gin.RouterGroup, db *gorm.DB, sere *sere.SERE) {
	newHotelRepo := hotelrepo.NewHotelRepository(db)

	newKitchenRepo := kitchenrepo.NewKitchenRepository(db)
	newHouseKeepingRepo := hotelrepo.NewHotelHouseKeepingRepository(db)
	newHotelService := hotelservices.NewHotelService(sere, newHotelRepo, newHouseKeepingRepo, newKitchenRepo)
	newHotelController := hotelcontrollers.NewHotelController(newHotelService)

	/************************
	 * Main Routes of Hotel *
	 ************************/
	hotel_routes := v1.Group("/hotel")

	hotel_routes.POST("/configure", newHotelController.ConfigureAndUpdate)
	hotel_routes.GET("/frontdesk", newHotelController.GetHotelFrontDesk)
	hotel_routes.GET("/reception", newHotelController.GetHotelReception)
	hotel_routes.GET("/get-hotel", newHotelController.GetHotel)
	hotel_routes.GET("/overview", newHotelController.GetHotelOverview)

	/*********************************
	 * Sub Routes for Floor in Hotel *
	 *********************************/
	hotelFloor := hotel_routes.Group("/floor")
	{
		hotelFloor.POST("/create", newHotelController.CreateFloorInHotel)
		hotelFloor.GET("/all-rooms/:floorId")
	}

	/********************************
	 * Sub Routes of Hotel for Room *
	 ********************************/
	hotelRoom := hotel_routes.Group("/room")
	{
		hotelRoom.GET("")
		hotelRoom.POST("/create", newHotelController.CreateRoom)
		hotelRoom.POST("/update/:roomId")
		hotelRoom.GET("/delete/:roomId")
		hotelRoom.GET("/room-status/:hotelId/:roomId/:status", newHotelController.UpdateRoomStatus)
	}

	/********************************
	 * Sub Routes of Hotel for Booking *
	 ********************************/
	hotelBooking := hotel_routes.Group("/booking")
	{
		hotelBooking.POST("/create", newHotelController.CreateBooking)
		hotelBooking.GET("/:bookingId")
		hotelBooking.GET("/details/:roomId", newHotelController.GetBookingDetails)

	}

	/*************************************************
	 * Sub Routes For the HouseKeeping for the Hotel *
	 *************************************************/
	houseKeeping := hotel_routes.Group("/houseKeeping")
	{
		houseKeeping.GET("/requests")
		houseKeeping.POST("")
	}

	/*************************************
	 * Sub Routes for the Guest of Hotel *
	 *************************************/
	hotelGuest := hotel_routes.Group("/guest")
	{
		hotelGuest.GET("")
		hotelGuest.GET("/:hotelGuestId")
		hotelGuest.GET("/search", newHotelController.SearchGuests)
		hotelGuest.POST("/create", newHotelController.CreateHotelGuest)
	}

	hotelTax := hotel_routes.Group("/tax")
	{
		hotelTax.POST("/create-rule", newHotelController.CreateTaxRule)
	}

}
