package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type HotelRepository interface {
	//--------------------------------------------------------------
	// Hotel
	//--------------------------------------------------------------
	CreateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error
	GetHotelByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.Hotel, error)
	GetHotelByOrgID(ctx context.Context, orgid uuid.UUID) (*hospitalitymodels.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error

	//-------------------------------------------------------------
	//   Floors
	// ------------------------------------------------------------

	CreateFloorInHotel(ctx context.Context, floor *hospitalitymodels.HotelFloor) error
	GetAllRoomInFloor()
	GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hospitalitymodels.HotelFloor, error)
	UpdateFloor(ctx context.Context, floor *hospitalitymodels.HotelFloor) error
	DeleteFloor(ctx context.Context, floorID uuid.UUID) error

	//-------------------------------------------------------------
	//   Rooms
	// ------------------------------------------------------------
	CreateRoom(ctx context.Context, room *hospitalitymodels.Room) error
	GetRoomByID(ctx context.Context, roomID uuid.UUID) (*hospitalitymodels.Room, error)
	UpdateRoom(ctx context.Context, fields map[string]interface{}, roomId uuid.UUID) error
	DeleteRoom(ctx context.Context, roomID uuid.UUID) error

	GetFloorByID(ctx context.Context, id string) (*hospitalitymodels.HotelFloor, error)

	//-------------------------------------------------------------
	//   Booking
	// ------------------------------------------------------------

	CreateRoomBooking(ctx context.Context, booking *hospitalitymodels.HotelRoomBooking) error
	CreateOrAddBookingPayment(ctx context.Context, bookingPayment *hospitalitymodels.BookingPayment) error
	CreatePrimaryGuest(ctx context.Context, guest *hospitalitymodels.HotelGuest) error
	CreateCharges(ctx context.Context, charges *hospitalitymodels.BookingCharge) error
	GetBookingDetailsByRoomId(ctx context.Context, roomId uuid.UUID) (*hospitalitymodels.HotelRoomBooking, error)
}

type HotelHouseKeepingRepositoryInterface interface {

	// TODO :- This need To Implemented
}
