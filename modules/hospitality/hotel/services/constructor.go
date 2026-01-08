package hotelservices

import (
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	"sharpnr-engine/sere"
)

type hotelService struct {
	sere      *sere.SERE // 🔥 event engine
	hotelRepo hotelrepo.HotelRepositoryInterface
}

func NewHotelService(
	sere *sere.SERE,
	hotelRepo hotelrepo.HotelRepositoryInterface,

) *hotelService {
	return &hotelService{
		sere:      sere,
		hotelRepo: hotelRepo,
	}
}
