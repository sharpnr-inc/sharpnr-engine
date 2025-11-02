package hospitality

import (
	"sharpnr-engine/modules/hospitality/bar"
	"sharpnr-engine/modules/hospitality/hotel"
	"sharpnr-engine/modules/hospitality/kitchen"
	"sharpnr-engine/modules/hospitality/restaurant"
)

type HospitalityModuleRepository interface {
	bar.BarRepository
	hotel.HotelRepository
	kitchen.KitchenRepository
	restaurant.RestaurantRepository
}
