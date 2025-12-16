package hospitalitydtos

import (
	"time"

	"github.com/google/uuid"
)

type RoomResponse struct {
	RoomID      uuid.UUID `json:"room_id"`
	RoomName    string    `json:"room_name"`
	RoomPrice   int       `json:"room_price"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}
type FloorResponse struct {
	FloorID     uuid.UUID      `json:"floor_id"`
	FloorNumber int            `json:"floor_number"`
	FloorName   string         `json:"floor_name"`
	Rooms       []RoomResponse `json:"rooms"`
}

type GetHotel_RESPONSE struct {
	HotelId          string `json:"hotel_id"`
	HotelName        string `json:"hotel_name"`
	SharpId          string
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

type Hotel_UpdateOrConfigure_REQUEST struct {
	DisplayName      *string `json:"display_name"`
	LegalName        *string `json:"legal_name"`
	Address          *string `json:"address"`
	City             *string `json:"city"`
	State            *string `json:"state"`
	Country          *string `json:"country"`
	Pincode          *string `json:"pincode"`
	HotelGstNumber   *string `json:"hotel_gst_number"`
	HotelFSSAINumber *string `json:"hotel_fssai_number"`
}

// This is the Response Struct for the Hotel Booking Terminal
// which will be used to booking and enquiry on Front desk
type HotelBookingTerminal_Response struct {
	Hotel_Id uuid.UUID
	Org_Id   uuid.UUID
	// Floor response is reused here for
	Floors []FloorResponse
}

type Hotel_Room_booking_Req struct {
	Hotel_Id     uuid.UUID
	RoomId       uuid.UUID
	CheckInnDate time.Time
	CheckoutDate time.Time
	// We need check inn time to calculate the date and time
	BookingDate     time.Time
	NumberOfGuest   int
	GuestEmailAddr  string
	GuestName       string
	GuestMobile     string
	AdvancedPayment float64
}
