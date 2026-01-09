package hospitality

import (
	"sharpnr-engine/app"
	"sharpnr-engine/middlewares"
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/concierge"
	hospitalityinventory "sharpnr-engine/modules/hospitality/hospitalityInventory"
	"sharpnr-engine/modules/hospitality/hotel"
	"sharpnr-engine/modules/hospitality/kitchen"
	hospitalitymenu "sharpnr-engine/modules/hospitality/menu"
	"sharpnr-engine/modules/hospitality/restaurant"
	"sharpnr-engine/pkg/jwtpack"

	"github.com/gin-gonic/gin"
)

func RegisterHotelModuleRoutes(
	v1 *gin.RouterGroup,
	modules app.ModulesProvider,
) {

	// Route Group
	sharpHospitality := v1.Group("/sharp-hospitality")

	//Middleware layer
	sharpHospitality.Use(middlewares.RequireTokenType(jwtpack.TokenTypeAuth), middlewares.SetOrgUUID())

	// Hospitality Routes
	hotel.Hospitalityhotelroutes(sharpHospitality, modules)

	kitchen.KitchenRoutes(sharpHospitality, modules)

	restaurant.RestaurantRoutes(sharpHospitality, modules)

	bar.BarRoutes(sharpHospitality, modules)

	hospitalitymenu.MenuRoutes(sharpHospitality, modules)

	hospitalityinventory.HospitalityRoutes(sharpHospitality, modules)

	concierge.ConciergeRoutes(sharpHospitality, modules)

}
