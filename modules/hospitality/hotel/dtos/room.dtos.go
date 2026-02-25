package hoteldtos

import (
	"github.com/google/uuid"
)

type HotelLayoutResponseDTO struct {
	Layout []HotelFloorResponseDTO
}

type HotelFloorResponseDTO struct {
	FloorID     uuid.UUID
	FloorName   string
	FloorNumber int
	Rooms       []HotelRoomResponseDTO
}

type HotelRoomResponseDTO struct {
	RoomID   uuid.UUID
	RoomName string
	Status   string
	Price    float64
	BedType  string
	Code     string
}
