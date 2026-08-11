package hotel_routes

import (
	"sharpnr-engine/app"
	hotelcontrollers "sharpnr-engine/property_business_ops/hospitality/hotel/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHotelRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	hotelService := modules.Hospitality().HotelService()
	hc := hotelcontrollers.NewHotelController(hotelService)

	hr := v1.Group("/hotel/:hotelID/")

	hr.GET("/layout", hc.GetHotelLayout)
	hr.GET("")
}
