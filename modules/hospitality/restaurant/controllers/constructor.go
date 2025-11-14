package restaurantcontrollers

import "sharpnr-engine/modules/hospitality/restaurant/restaurantservices"

type restaurantController struct {
	resSvc restaurantservices.RestaurantService
}

func NewRestaurantController(resSvc restaurantservices.RestaurantService) *restaurantController {
	return &restaurantController{
		resSvc: resSvc,
	}
}
