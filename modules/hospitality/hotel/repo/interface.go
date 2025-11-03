package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, room *hospitalitymodels.Hotel) (*hospitalitymodels.Hotel, error)

	CreateRoom(ctx context.Context, room *hospitalitymodels.Room) (*hospitalitymodels.Room, error)
	GetRoomDetails()
	EditRoom()
	CreateFloor()
}
