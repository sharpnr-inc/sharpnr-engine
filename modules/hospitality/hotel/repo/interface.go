package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error

	CreateRoom(ctx context.Context, room *hospitalitymodels.Room) error
	GetRoomDetails()
	EditRoom()
	CreateFloor()
}
