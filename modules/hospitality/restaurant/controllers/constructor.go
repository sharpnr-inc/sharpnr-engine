package restaurantcontrollers

import restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"

type RestaurantController struct {
	resSvc restaurantservices.RestaurantServiceInterface
}

func NewRestaurantController(resSvc restaurantservices.RestaurantServiceInterface) *RestaurantController {
	return &RestaurantController{
		resSvc: resSvc,
	}
}
