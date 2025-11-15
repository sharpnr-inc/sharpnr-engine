package hospitalitydtos

import "github.com/google/uuid"

type RoomResponse struct {
	RoomID      string `json:"room_id"`
	RoomName    string `json:"room_name"`
	RoomPrice   int    `json:"room_price"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type FloorResponse struct {
	FloorID     string         `json:"floor_id"`
	FloorNumber int            `json:"floor_number"`
	FloorName   string         `json:"floor_name"`
	Rooms       []RoomResponse `json:"rooms"`
}

type GetHotel_RESPONSE struct {
	HotelId          string `json:"hotel_id"`
	HotelName        string `json:"hotel_name"`
	HotelDisplayName string
	HotelLegalName   string
	KitchenId        uuid.UUID
	IsKitchen        bool
	Address          string
	City             string
	State            string
	Country          string
	Pincode          string
	HotelGstNumber   string
	HotelFSSAINumber string
	Floors           []FloorResponse `json:"floors"`
}

type CreateRoomInFloor_REQ struct {
	FloorId     string `json:"floor_id"`
	RoomName    string `json:"room_name"`
	RoomPrice   int    `json:"room_price"`
	Description string `json:"description"`
}

type CreateRoomInFloor_RES struct {
	RoomName    string `json:"room_name"`
	RoomPrice   int    `json:"room_price"`
	Description string `json:"description"`
}
