package hospitalitydtos

type CreateAndAddRoomInHotel struct {
	RoomName    string
	FloorID     string
	AmenitiesID []string
	RoomStatus  string
	Status      string
}
