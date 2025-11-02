package hotelservices

import hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"

type hotelService struct {
	hotelRepo hotelrepo.HotelRepository
}

// CreateRoomInHotel implements HotelService.
func (h *hotelService) CreateRoomInHotel() {
	panic("unimplemented")
}

func NewHotelService(hotelRepo hotelrepo.HotelRepository) *hotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}
