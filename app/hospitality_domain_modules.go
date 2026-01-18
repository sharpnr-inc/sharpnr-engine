package app

import (
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
)

type HospitalityDomain struct {
	hotelService      hotelservices.HotelService
	barService        barservices.BarServices
	restaurantService restaurantservices.RestaurantService
	kitchenService    kitchenservices.KitchenServiceInterface
	menuService       menuservice.MenuService
}

func NewHospitalityDomain(
	hotelService hotelservices.HotelService,
	barService barservices.BarServices,
	restaurantService restaurantservices.RestaurantService,
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

func (h *HospitalityDomain) HotelService() hotelservices.HotelService {
	return h.hotelService
}

func (h *HospitalityDomain) BarService() barservices.BarServices {
	return h.barService
}

func (h *HospitalityDomain) RestaurantService() restaurantservices.RestaurantService {
	return h.restaurantService
}

func (h *HospitalityDomain) KitchenService() kitchenservices.KitchenServiceInterface {
	return h.kitchenService
}

func (h *HospitalityDomain) MenuService() menuservice.MenuService {
	return h.menuService
}
