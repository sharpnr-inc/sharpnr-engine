package hospitality

import (
	"sharpnr-engine/middlewares"
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/hotel"
	"sharpnr-engine/modules/hospitality/kitchen"
	"sharpnr-engine/modules/hospitality/restaurant"
	"sharpnr-engine/pkg/jwtpack"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHotelModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	sharpHospitality := v1.Group("/sharp-hospitality")
	sharpHospitality.Use(middlewares.RequireTokenType(jwtpack.TokenTypeAuth), middlewares.SetOrgUUID())
	hotel.Hospitalityhotelroutes(sharpHospitality, db)
	kitchen.KitchenRoutes(sharpHospitality, db)
	restaurant.RestaurantRoutes(sharpHospitality, db)
	bar.BarRoutes(sharpHospitality, db)

}
