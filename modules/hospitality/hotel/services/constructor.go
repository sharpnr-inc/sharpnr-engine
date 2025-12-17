package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
)

type hotelService struct {
	sere        *sere.SERE // 🔥 event engine
	hotelRepo   hotelrepo.HotelRepositoryInterface
	houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
}

func NewHotelService(
	sere *sere.SERE,
	hotelRepo hotelrepo.HotelRepositoryInterface,
	houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface,
	kitchenRepo kitchenrepo.KitchenRepositoryInterface) *hotelService {
	return &hotelService{
		sere:        sere,
		hotelRepo:   hotelRepo,
		houseKeRepo: houseKeRepo,
		kitchenRepo: kitchenRepo,
	}
}
