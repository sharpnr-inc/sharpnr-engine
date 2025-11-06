package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
)

// CreateRoomInHotel implements HotelService.
func (h *hotelService) CreateRoomInHotel(ctx context.Context, req hospitalitydtos.CreateAndAddRoomInHotel_REQ) error {
	return nil
}
