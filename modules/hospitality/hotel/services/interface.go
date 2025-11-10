package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type HotelService interface {
	CreateRoomInHotel(ctx context.Context, orgID uuid.UUID, req hospitalitydtos.CreateRoomInFloor_REQ) (*hospitalitydtos.CreateRoomInFloor_RES, error)

	CreateFloor(ctx context.Context, orgID uuid.UUID, req hospitalitydtos.CreateFloorInHotel_REQ) (*hospitalitydtos.CreateFloorInHotel_RESPONSE, error)

	GetHotelDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetHotel_RESPONSE, error)
}

type HotelBookingServiceInterface interface {
	// This is will implement the booking services
}
