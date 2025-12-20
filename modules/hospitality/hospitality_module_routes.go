package hospitality

import (
	"sharpnr-engine/middlewares"
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/concierge"
	hospitalityinventory "sharpnr-engine/modules/hospitality/hospitalityInventory"
	"sharpnr-engine/modules/hospitality/hotel"
	"sharpnr-engine/modules/hospitality/kitchen"
	hospitalitymenu "sharpnr-engine/modules/hospitality/menu"
	"sharpnr-engine/modules/hospitality/restaurant"
	"sharpnr-engine/pkg/jwtpack"
	"sharpnr-engine/sere"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHotelModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB, sere *sere.SERE) {

	// Route Group
	sharpHospitality := v1.Group("/sharp-hospitality")

	//Middleware layer
	sharpHospitality.Use(middlewares.RequireTokenType(jwtpack.TokenTypeAuth), middlewares.SetOrgUUID())

	// Hospitality Routes
	hotel.Hospitalityhotelroutes(sharpHospitality, db, sere)
	kitchen.KitchenRoutes(sharpHospitality, db)
	restaurant.RestaurantRoutes(sharpHospitality, db)
	bar.BarRoutes(sharpHospitality, db)
	hospitalitymenu.MenuRoutes(sharpHospitality, db)
	hospitalityinventory.HospitalityRoutes(sharpHospitality, db)
	concierge.ConciergeRoutes(sharpHospitality, db)

}
