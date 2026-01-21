package hoteldtos

import "time"

type RoomPriceDTO struct {
	Price         float64    `json:"price"`
	EffectiveFrom time.Time  `json:"effectiveFrom"`
	EffectiveTo   *time.Time `json:"effectiveTo,omitempty"`
}

type CreateRoomInFloor_REQ struct {
	FloorId      string       `json:"floorId"`
	RoomName     string       `json:"roomName"`
	BedSize      string       `json:"bedSize"`
	Amenities    string       `json:"amenities"`
	Description  string       `json:"description"`
	InitialPrice RoomPriceDTO `json:"initialPrice"`
}
