package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type HotelRepositoryInterface interface {
	WithTx(
		ctx context.Context,
		fn func(txRepo HotelRepositoryInterface) error,
	) error

	//--------------------------------------------------------------
	// Hotel
	//--------------------------------------------------------------
	CreateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error
	GetHotelByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.Hotel, error)
	GetHotelByOrgID(ctx context.Context, orgid uuid.UUID) (*hospitalitymodels.Hotel, error)

	UpdateHotelFields(
		ctx context.Context,
		hotelID uuid.UUID,
		updates map[string]interface{},
	) error

	//-------------------------------------------------------------
	//   Floors
	// ------------------------------------------------------------

	CreateFloorInHotel(ctx context.Context, floor *hospitalitymodels.HotelFloor) error
	GetFloorByID(ctx context.Context, id string) (*hospitalitymodels.HotelFloor, error)
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
	UpdateRoomStatus(ctx context.Context, hotelId, roomId uuid.UUID, status string) error
	DeleteRoom(ctx context.Context, roomID uuid.UUID) error
	AddCurrentBookingIdToRoom(ctx context.Context, roomId, bookingId uuid.UUID) error
	CreatePriceForRoom(ctx context.Context, roomPrice *hospitalitymodels.HotelRoomPrice) error
	GetRoomPriceByRoomId(ctx context.Context, roomId uuid.UUID) (*hospitalitymodels.HotelRoomPrice, error)

	//-------------------------------------------------------------
	//   Booking
	// ------------------------------------------------------------
	CreateRoomBooking(ctx context.Context, booking *hospitalitymodels.HotelRoomBooking) error
	CreateOrAddBookingPayment(ctx context.Context, bookingPayment *hospitalitymodels.HotelRoomBookingPayment) error
	CreateCharges(ctx context.Context, charges *hospitalitymodels.HotelBookingCharge) error
	GetBookingByBookingId(ctx context.Context, bookingId uuid.UUID) (*hospitalitymodels.HotelRoomBooking, error)

	//-------------------------------------------------------------
	//   Guests
	// ------------------------------------------------------------
	CreatePrimaryGuest(ctx context.Context, guest *hospitalitymodels.HotelGuest) error
	GetHotelGuests(ctx context.Context, hotelId uuid.UUID) ([]hospitalitymodels.HotelGuest, error)
	SearchHotelGuest(ctx context.Context, hotelId uuid.UUID, qry string) ([]hospitalitymodels.HotelGuest, error)
	GetHotelGuestById(ctx context.Context, guestId uuid.UUID) (hospitalitymodels.HotelGuest, error)
	CreateBookingGuest(ctx context.Context, bookingGuest *hospitalitymodels.BookingGuest) error

	//-------------------------------------------------------------
	//   Charges
	// ------------------------------------------------------------
	GetChargesByBookingID(ctx context.Context, bookingId uuid.UUID) ([]hospitalitymodels.HotelBookingCharge, error)
	AddChargesInBooking(ctx context.Context, charges *hospitalitymodels.HotelBookingCharge) error

	//-------------------------------------------------------------
	//   Bills
	// ------------------------------------------------------------
	GetBillByBookingId(ctx context.Context, bookingId uuid.UUID) (*hospitalitymodels.HotelBill, error)

	//-------------------------------------------------------------
	//   Payments
	// ------------------------------------------------------------
	GetPaymentsByBookingId(ctx context.Context, bookingId uuid.UUID) ([]hospitalitymodels.HotelRoomBookingPayment, error)
}

type HotelHouseKeepingRepositoryInterface interface {

	// TODO :- This need To Implemented
}
