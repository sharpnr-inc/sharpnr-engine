package hospitality

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/concierge"
	"sharpnr-engine/modules/hospitality/hotel"
	"sharpnr-engine/modules/hospitality/kitchen"
	hospitalitymenu "sharpnr-engine/modules/hospitality/menu"
	"sharpnr-engine/modules/hospitality/restaurant"

	"github.com/gin-gonic/gin"
)

func RegisterHotelModuleRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {

	// Route Group
	sharpHospitality := v1.Group("/sharp-hospitality")

	hotel.Hospitalityhotelroutes(sharpHospitality, modules)

	kitchen.KitchenRoutes(sharpHospitality, modules)

	restaurant.RestaurantRoutes(sharpHospitality, modules)

	bar.BarRoutes(sharpHospitality, modules)

	hospitalitymenu.MenuRoutes(sharpHospitality, modules)

	concierge.ConciergeRoutes(sharpHospitality, modules)

}
