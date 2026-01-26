package app

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	barservices "sharpnr-engine/modules/hospitality/bar/services"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	hotelservices "sharpnr-engine/modules/hospitality/hotel/services"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/repos"
	kitchenservices "sharpnr-engine/modules/hospitality/kitchen/services"
	menurepo "sharpnr-engine/modules/hospitality/menu/repo"
	menuservice "sharpnr-engine/modules/hospitality/menu/service"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	restaurantservices "sharpnr-engine/modules/hospitality/restaurant/services"
	"sharpnr-engine/sere"

	"gorm.io/gorm"
)

func bootstrapHospitality(
	db *gorm.DB,
	sereEngine *sere.SERE,
) *hospitalitydomain.HospitalityDomain {

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

	return hospitalitydomain.NewHospitalityDomain(
		hotelService,
		barService,
		restaurantService,
		kitchenService,
		menuService,
	)
}
