package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/google/uuid"
)

type HotelService interface {
	CreateRoomInHotel(ctx context.Context, req hospitalitydtos.CreateAndAddRoomInHotel_REQ) error

	CreateFloor(ctx context.Context, orgID uuid.UUID, req hospitalitydtos.CreateFloorInHotel_REQ) (*hospitalitydtos.CreateFloorInHotel_RESPONSE, error)

	GetHotelDetails(ctx context.Context, orgId uuid.UUID) (*hospitalitydtos.GetHotel_RESPONSE, error)
}
