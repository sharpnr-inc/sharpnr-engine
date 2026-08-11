package hotel_contracts

import "github.com/google/uuid"

type HotelReader struct {
	ID           uuid.UUID
	Name         string
	Status       bool
	IsConfigured bool
}
