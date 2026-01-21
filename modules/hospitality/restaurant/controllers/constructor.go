package restaurantcontrollers

import restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"

type restaurantController struct {
	resSvc restaurantservices.RestaurantServiceInterface
}

func NewRestaurantController(resSvc restaurantservices.RestaurantServiceInterface) *restaurantController {
	return &restaurantController{
		resSvc: resSvc,
	}
}
