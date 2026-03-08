package restaurant_routes

import (
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                             Restaurant Genrals                             */
/* -------------------------------------------------------------------------- */

func restaurantGenralRoutes(v1 *gin.RouterGroup, rc restaurantcontrollers.RestaurantController) {

	v1.GET("/general-details", rc.GetGeneralDetails)
	v1.POST("/update-general-details", rc.UpdateGeneralDetails)
	v1.GET("/layout", rc.GetRestaurantLayout)
	v1.GET("/preferred-kitchen", rc.GetRestaurantPreferredKitchen)
}

/* -------------------------------------------------------------------------- */
/*                             Restaurant Configs                             */
/* -------------------------------------------------------------------------- */

func restaurantConfigsRoutes(
	resRoutes *gin.RouterGroup,
	restaurantController restaurantcontrollers.RestaurantController) {

	rc := resRoutes.Group("/configs")

	rc.POST("/add-layout", restaurantController.AddLayout)
	rc.POST("/add-preferred-kitchen/:pfKitchenID", restaurantController.AddRestaurantPreferredKitchen)
	rc.PATCH("/activate-deactivate/:switch", restaurantController.ActivateDeactivateKitchen)

}
