package restaurantcontrollers

import restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"

type restaurantController struct {
	resSvc restaurantservices.RestaurantService
}

func NewRestaurantController(resSvc restaurantservices.RestaurantService) *restaurantController {
	return &restaurantController{
		resSvc: resSvc,
	}
}
