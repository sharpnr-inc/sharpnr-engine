package hotelservices

import hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

type HotelService interface {
	CreateRoomInHotel(hospitalitydtos.CreateAndAddRoomInHotel_REQ)
}
