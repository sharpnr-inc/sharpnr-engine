package hotel_models

import "github.com/google/uuid"

type HotelLocations struct {
	ID          uuid.UUID
	HotelID     uuid.UUID
	OrgID       uuid.UUID
	AddressLine string
	City        string
	State       string
	Country     string
	Pincode     string
}
