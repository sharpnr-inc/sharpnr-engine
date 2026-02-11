package hotel_routes

import (
	"sharpnr-engine/app"
	hotelcontrollers "sharpnr-engine/modules/hospitality/hotel/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHotelRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {
	hotelService := modules.Hospitality().HotelService()
	hc := hotelcontrollers.NewHotelController(hotelService)

	hr := v1.Group("/hotel")

	hr.GET("/", hc.GetOverview)
}
