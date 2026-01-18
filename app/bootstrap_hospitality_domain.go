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
	"sharpnr-engine/modules/hospitality/restaurant/restaurantservices"
	"sharpnr-engine/sere"

	"gorm.io/gorm"
)

func bootstrapHospitality(
	db *gorm.DB,
	sereEngine *sere.SERE,
) *HospitalityDomain {

	// Kitchen
	kitchenRepo := kitchenrepo.NewKitchenRepository(db)
	kitchenService := kitchenservices.NewKitchenService(kitchenRepo)

	// Menu
	menuRepo := menurepo.NewHospitalityMenuRepository(db)
	menuService := menuservice.NewHospitalityMenuService(menuRepo)

	// Hotel
	hotelRepo := hotelrepo.NewHotelRepository(db)
	hotelService := hotelservices.NewHotelService(sereEngine, hotelRepo)

	// Bar
	barRepo := barrepo.NewBarRepository(db)
	barService := barservices.NewBarServices(barRepo)

	// Restaurant
	restaurantRepo := restaurantrepo.NewRestaurantRepository(db)
	restaurantService := restaurantservices.NewRestaurantService(restaurantRepo)

	return NewHospitalityDomain(
		hotelService,
		barService,
		restaurantService,
		kitchenService,
		menuService,
	)
}
