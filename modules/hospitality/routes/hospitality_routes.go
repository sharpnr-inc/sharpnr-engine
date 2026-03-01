package hospitality_routes

import (
	"sharpnr-engine/app"
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/concierge"
	hotel_routes "sharpnr-engine/modules/hospitality/hotel/routes"
	kitchen_routes "sharpnr-engine/modules/hospitality/kitchen/routes"
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

	hotel_routes.RegisterHotelRoutes(sharpHospitality, modules)

	kitchen_routes.KitchenRoutes(sharpHospitality, modules)
	menu_routes.RegisterMenuRoutes(sharpHospitality, modules)

	restaurant.RestaurantRoutes(sharpHospitality, modules)

	bar.BarRoutes(sharpHospitality, modules)

	concierge.ConciergeRoutes(sharpHospitality, modules)

	hospitality_utilities.RegisterHospitalityUtilityRoutes(v1)

}
