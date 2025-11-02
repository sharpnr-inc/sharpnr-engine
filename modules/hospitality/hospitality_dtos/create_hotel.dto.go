package hospitalitydtos

type CreateAndAddRoomInHotel_REQ struct {
	RoomName    string
	FloorID     string
	AmenitiesID []string
	RoomStatus  string
	Status      string
}

type CreateAndAddRoomInHotel_RES struct {
	RoomName    string
	FloorID     string
	AmenitiesID []string
	RoomStatus  string
	Status      string
}
