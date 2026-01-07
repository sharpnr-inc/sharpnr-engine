package hospitality_setup

import (
	"sharpnr-engine/core/models"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	"sharpnr-engine/modules/hospitality/restaurant/restaurantrepo"
)

type SetupContext struct {
	Org     *models.Organization
	Kitchen *hospitalitymodels.Kitchen

	BarRepo        barrepo.BarRepository
	HotelRepo      hotelrepo.HotelRepositoryInterface
	RestaurantRepo restaurantrepo.RestaurantRepository
	MenuRepo       menurepo.MenuRepositoryInterface
}
