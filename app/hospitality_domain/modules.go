package hospitalitydomain

import (
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	menuservice "sharpnr-engine/modules/hospitality/menu/service"
	hospitality_order_engine_service "sharpnr-engine/modules/hospitality/order_engine/service"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"
)

type HospitalityDomain struct {
	hotelService      hotelservices.HotelServiceInterface
	barService        barservices.BarServicesInterface
	restaurantService restaurantservices.RestaurantServiceInterface
	kitchenService    kitchenservices.KitchenServiceInterface
	menuService       menuservice.MenuService
	orderEngine       hospitality_order_engine_service.HospitalityOrderEngineServiceInterface
}

func NewHospitalityDomain(
	hotelService hotelservices.HotelServiceInterface,
	barService barservices.BarServicesInterface,
	restaurantService restaurantservices.RestaurantServiceInterface,
	kitchenService kitchenservices.KitchenServiceInterface,
	menuService menuservice.MenuService,
	orderEngine hospitality_order_engine_service.HospitalityOrderEngineServiceInterface,
) *HospitalityDomain {
	return &HospitalityDomain{
		hotelService:      hotelService,
		barService:        barService,
		restaurantService: restaurantService,
		kitchenService:    kitchenService,
		menuService:       menuService,
		orderEngine:       orderEngine,
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

func (h *HospitalityDomain) OrderEngine() hospitality_order_engine_service.HospitalityOrderEngineServiceInterface {
	return h.orderEngine
}
