package hospitality_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/concierge"
	"sharpnr-engine/modules/hospitality/hotel"
	"sharpnr-engine/modules/hospitality/kitchen"
	menu_routes "sharpnr-engine/modules/hospitality/menu/routes"
	"sharpnr-engine/modules/hospitality/restaurant"
	hospitality_utilities "sharpnr-engine/modules/hospitality/utilities_api"

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
	menu_routes.RegisterMenuRoutes(sharpHospitality, modules)

	restaurant.RestaurantRoutes(sharpHospitality, modules)

	bar.BarRoutes(sharpHospitality, modules)

	concierge.ConciergeRoutes(sharpHospitality, modules)

	hospitality_utilities.RegisterHospitalityUtilityRoutes(v1)

}
