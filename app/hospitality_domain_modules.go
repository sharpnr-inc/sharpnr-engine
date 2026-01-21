package app

import (
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	menuservice "sharpnr-engine/modules/hospitality/menu/service"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"
)

type HospitalityDomain struct {
	hotelService      hotelservices.HotelServiceInterface
	barService        barservices.BarServicesInterface
	restaurantService restaurantservices.RestaurantServiceInterface
	kitchenService    kitchenservices.KitchenServiceInterface
	menuService       menuservice.MenuService
}

func NewHospitalityDomain(
	hotelService hotelservices.HotelServiceInterface,
	barService barservices.BarServicesInterface,
	restaurantService restaurantservices.RestaurantServiceInterface,
	kitchenService kitchenservices.KitchenServiceInterface,
	menuService menuservice.MenuService,
) *HospitalityDomain {
	return &HospitalityDomain{
		hotelService:      hotelService,
		barService:        barService,
		restaurantService: restaurantService,
		kitchenService:    kitchenService,
		menuService:       menuService,
	}
}

/* ---- getters (ONLY what you want to expose) ---- */

func (h *HospitalityDomain) HotelService() hotelservices.HotelServiceInterface {
	return h.hotelService
}

func (h *HospitalityDomain) BarService() barservices.BarServicesInterface {
	return h.barService
}

func (h *HospitalityDomain) RestaurantService() restaurantservices.RestaurantServiceInterface {
	return h.restaurantService
}

func (h *HospitalityDomain) KitchenService() kitchenservices.KitchenServiceInterface {
	return h.kitchenService
}

func (h *HospitalityDomain) MenuService() menuservice.MenuService {
	return h.menuService
}
