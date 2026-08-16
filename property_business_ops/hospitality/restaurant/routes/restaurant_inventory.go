package restaurant_routes

import (
	restaurantcontrollers "sharpnr-engine/property_business_ops/hospitality/restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func registerRestaurantInventoryRoutes(
	resRoutes *gin.RouterGroup,
	restaurantController restaurantcontrollers.RestaurantController,
) {

	resInv := resRoutes.Group("/inventory")
	{
		resInv.GET("", restaurantController.GetBusinessUnitInventories) // uses filter of the InventoryAccountID
		resInv.GET("/:inventoryID/items", restaurantController.GetInventoryItems)
		resInv.GET("/accounts", restaurantController.GetAllInventoryAccounts)
		resInv.POST("/add-inventory", restaurantController.AddInventoryInInventoryAccount)
	}

}
