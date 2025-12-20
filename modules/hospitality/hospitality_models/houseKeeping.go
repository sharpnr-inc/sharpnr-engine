package hospitalitymodels

import "github.com/google/uuid"

// this will House keeping Model
type HouseKeeping struct {
	Id      uuid.UUID
	OrgId   uuid.UUID
	HotelId uuid.UUID
}
