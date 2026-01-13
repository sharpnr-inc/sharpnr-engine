package app

import (
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
)

type HospitalityDomain struct {
	hotelRepo         hotelrepo.HotelRepositoryInterface
	hotelService      hotelservices.HotelService
	barService        barservices.BarServices
	restaurantService restaurantservices.RestaurantService
	kitchenService    kitchenservices.KitchenServiceInterface
	menuService       menuservice.MenuService
}

func NewHospitalityDomain(
	hotelRepo hotelrepo.HotelRepositoryInterface,
	hotelService hotelservices.HotelService,
	barService barservices.BarServices,
	restaurantService restaurantservices.RestaurantService,
	kitchenService kitchenservices.KitchenServiceInterface,
	menuService menuservice.MenuService,
) *HospitalityDomain {
	return &HospitalityDomain{
		hotelRepo:         hotelRepo,
		hotelService:      hotelService,
		barService:        barService,
		restaurantService: restaurantService,
		kitchenService:    kitchenService,
		menuService:       menuService,
	}
}

/* ---- getters (ONLY what you want to expose) ---- */

func (h *HospitalityDomain) HotelRepo() hotelrepo.HotelRepositoryInterface {
	return h.hotelRepo
}

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
