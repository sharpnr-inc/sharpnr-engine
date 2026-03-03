package restaurant_routes

import (
	restaurantcontrollers "sharpnr-engine/modules/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func restaurantGenralRoutes(v1 *gin.RouterGroup, rc restaurantcontrollers.RestaurantController) {

	v1.GET("/general-details", rc.GetGeneralDetails)
	v1.POST("/update-general-details", rc.UpdateGeneralDetails)
	v1.GET("/layout", rc.GetRestaurantLayout)
}

func restaurantConfigsRoutes(
	resRoutes *gin.RouterGroup,
	restaurantController restaurantcontrollers.RestaurantController) {

	resRoutes.POST("/configs/add-layout", restaurantController.AddLayout)

}
