package hotelrepo

import (
	"context"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

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
	CreateHotel(ctx context.Context, hotel *hotel_models.Hotel) error
	GetHotelByID(ctx context.Context, id uuid.UUID) (*hotel_models.Hotel, error)
	GetHotelByOrgID(ctx context.Context, orgid uuid.UUID) (*hotel_models.Hotel, error)

	UpdateHotelFields(
		ctx context.Context,
		hotelID uuid.UUID,
		updates map[string]interface{},
	) error

	//-------------------------------------------------------------
	//   Floors
	// ------------------------------------------------------------

	CreateFloorInHotel(ctx context.Context, floor *hotel_models.HotelFloor) error
	GetFloorByID(ctx context.Context, id string) (*hotel_models.HotelFloor, error)
	GetAllRoomInFloor()
	GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hotel_models.HotelFloor, error)
	UpdateFloor(ctx context.Context, floor *hotel_models.HotelFloor) error
	DeleteFloor(ctx context.Context, floorID uuid.UUID) error

	//-------------------------------------------------------------
	//   Rooms
	// ------------------------------------------------------------
	CreateRoom(ctx context.Context, room *hotel_models.Room) error
	GetRoomByID(ctx context.Context, roomID uuid.UUID) (*hotel_models.Room, error)
	UpdateRoom(ctx context.Context, fields map[string]interface{}, roomId uuid.UUID) error
	UpdateRoomStatus(ctx context.Context, hotelId, roomId uuid.UUID, status string) error
	DeleteRoom(ctx context.Context, roomID uuid.UUID) error
	AddCurrentBookingIdToRoom(ctx context.Context, roomId, bookingId uuid.UUID) error
	CreatePriceForRoom(ctx context.Context, roomPrice *hotel_models.HotelRoomPrice) error
	GetRoomPriceByRoomId(ctx context.Context, roomId uuid.UUID) (*hotel_models.HotelRoomPrice, error)

	//-------------------------------------------------------------
	//   Guests
	// ------------------------------------------------------------
	CreatePrimaryGuest(ctx context.Context, guest *hotel_models.HotelGuest) error
	GetHotelGuests(ctx context.Context, hotelId uuid.UUID) ([]hotel_models.HotelGuest, error)
	SearchHotelGuest(ctx context.Context, hotelId uuid.UUID, qry string) ([]hotel_models.HotelGuest, error)
	GetHotelGuestById(ctx context.Context, guestId uuid.UUID) (hotel_models.HotelGuest, error)
	CreateBookingGuest(ctx context.Context, bookingGuest *hotel_models.BookingGuest) error

	//-------------------------------------------------------------
	//   Booking
	// ------------------------------------------------------------
	CreateRoomBooking(ctx context.Context, booking *hotel_models.HotelRoomBooking) error
	CreateOrAddBookingPayment(ctx context.Context, bookingPayment *hotel_models.HotelRoomBookingPayment) error
	CreateCharges(ctx context.Context, charges *hotel_models.HotelBookingCharge) error
	GetBookingByBookingId(ctx context.Context, bookingId uuid.UUID) (*hotel_models.HotelRoomBooking, error)
	UpdateBookingStatus(
		ctx context.Context,
		bookingID uuid.UUID,
		status hospitalityconstants.HotelBookingStatusType,
	) error

	//-------------------------------------------------------------
	//   Charges
	// ------------------------------------------------------------
	GetChargesByBookingID(
		ctx context.Context,
		bookingId uuid.UUID,
	) ([]hotel_models.HotelBookingCharge, error)

	AddChargesInBooking(
		ctx context.Context,
		charges *hotel_models.HotelBookingCharge,
	) error

	//-------------------------------------------------------------
	//   Bills
	// ------------------------------------------------------------
	GetBillByBookingId(ctx context.Context, bookingId uuid.UUID) (*hotel_models.HotelBill, error)

	CreateBookingBill(ctx context.Context, bill *hotel_models.HotelBill) error

	//-------------------------------------------------------------
	//   Payments
	// ------------------------------------------------------------
	GetPaymentsByBookingId(
		ctx context.Context,
		bookingId uuid.UUID,
	) ([]hotel_models.HotelRoomBookingPayment, error)

	AddPaymentsInBooking(
		ctx context.Context,
		payemnt *hotel_models.HotelRoomBookingPayment,
	) error

	//-------------------------------------------------------------
	//    Tax
	// ------------------------------------------------------------
	CreateTaxRuleOfHotel(
		ctx context.Context,
		taxRule *hotel_models.HotelTaxRule,
	) error

	AddTaxComponentInTaxRule(
		ctx context.Context,
		taxComp *hotel_models.HotelTaxComponent,
	) error

	GetActiveTaxRulesByHotelAndType(
		ctx context.Context,
		hotelId uuid.UUID,
		taxtype hospitalityconstants.TaxType,
		time time.Time,
	) ([]hotel_models.HotelTaxRule, error)

	GetTaxComponentsByRuleID(
		ctx context.Context,
		taxRuleId uuid.UUID,
	) ([]hotel_models.HotelTaxComponent, error)

	//-------------------------------------------------------------
	//    Hotel Bill Tax
	// ------------------------------------------------------------

	CreateHotelBillTax(ctx context.Context, hotelBillTax *hotel_models.HotelBillTax) error
	GetBillTaxes(
		ctx context.Context,
		billID uuid.UUID,
	) ([]hotel_models.HotelBillTax, error)
}

/*******************************************
 * this is the Interface for House Keeping *
 *******************************************/
type HotelHouseKeepingRepositoryInterface interface {

	// TODO :- This need To Implemented
}
