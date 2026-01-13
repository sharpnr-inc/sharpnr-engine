package app

import (
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/kitchenServices"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
)

type HospitalityDomain struct {
	hotelRepo    hotelrepo.HotelRepositoryInterface
	hotelService hotelservices.HotelService

	barRepo    barrepo.BarRepository
	barService barservices.BarServices

	restaurantRepo    restaurantrepo.RestaurantRepository
	restaurantService restaurantservices.RestaurantService

	kitchenRepo    kitchenrepo.KitchenRepositoryInterface
	kitchenService kitchenservices.KitchenServiceInterface

	menuRepo    menurepo.MenuRepositoryInterface
	menuService menuservice.MenuService
}

func NewHospitalityDomain(
	hotelRepo hotelrepo.HotelRepositoryInterface,
	hotelService hotelservices.HotelService,

	barRepo barrepo.BarRepository,
	barService barservices.BarServices,

	restaurantRepo restaurantrepo.RestaurantRepository,
	restaurantService restaurantservices.RestaurantService,

	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
	kitchenService kitchenservices.KitchenServiceInterface,

	menuRepo menurepo.MenuRepositoryInterface,
	menuService menuservice.MenuService,
) *HospitalityDomain {
	return &HospitalityDomain{
		hotelRepo:    hotelRepo,
		hotelService: hotelService,

		barRepo:    barRepo,
		barService: barService,

		restaurantRepo:    restaurantRepo,
		restaurantService: restaurantService,

		kitchenRepo:    kitchenRepo,
		kitchenService: kitchenService,
		menuRepo:       menuRepo,
		menuService:    menuService,
	}
}

/* ---- getters (ONLY what you want to expose) ---- */

func (h *HospitalityDomain) HotelRepo() hotelrepo.HotelRepositoryInterface {
	return h.hotelRepo
}

func (h *HospitalityDomain) HotelService() hotelservices.HotelService {
	return h.hotelService
}

func (h *HospitalityDomain) BarRepo() barrepo.BarRepository {
	return h.barRepo
}

func (h *HospitalityDomain) BarService() barservices.BarServices {
	return h.barService
}

func (h *HospitalityDomain) RestaurantRepo() restaurantrepo.RestaurantRepository {
	return h.restaurantRepo
}

func (h *HospitalityDomain) RestaurantService() restaurantservices.RestaurantService {
	return h.restaurantService
}

func (h *HospitalityDomain) KitchenRepo() kitchenrepo.KitchenRepositoryInterface {
	return h.kitchenRepo
}

func (h *HospitalityDomain) KitchenService() kitchenservices.KitchenServiceInterface {
	return h.kitchenService
}

func (h *HospitalityDomain) MenuRepo() menurepo.MenuRepositoryInterface {
	return h.menuRepo
}

func (h *HospitalityDomain) MenuService() menuservice.MenuService {
	return h.menuService
}
