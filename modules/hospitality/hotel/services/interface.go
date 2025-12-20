package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"

	"github.com/google/uuid"
)

type HotelService interface {
	CreateRoomInHotel(
		ctx context.Context,
		orgID uuid.UUID,
		req hoteldtos.CreateRoomInFloor_REQ,
	) error

	UpdateHotelDetails(
		ctx context.Context,
		orgId uuid.UUID,
		req hospitalitydtos.Hotel_UpdateOrConfigure_REQUEST,
	) error

	CreateFloor(
		ctx context.Context,
		orgID uuid.UUID,
		req hospitalitydtos.CreateFloorInHotel_REQ,
	) (*hospitalitydtos.CreateFloorInHotel_RESPONSE, error)

	GetHotelDetails(
		ctx context.Context,
		orgId uuid.UUID,
	) (*hospitalitydtos.GetHotel_RESPONSE, error)

	GetFrontDeskTerminal(
		ctx context.Context,
		orgId uuid.UUID,
	) (*hospitalitydtos.HotelBookingTerminal_Response, error)

	UpdateRoomStatus(ctx context.Context, orgId, hotelId, roomId uuid.UUID, status string) error

	// -----------------
	// Bookings
	// ------------------
	CreateBookingForRoom(ctx context.Context, orgId uuid.UUID, req *hoteldtos.RoomBookingRequest_dto) error
	GetRoomBookingDetails(ctx context.Context, roomId, orgID uuid.UUID) (*hoteldtos.BookingDetailsDTO, error)

	// -----------------
	// Guest
	// ------------------

	SearchGuests(
		ctx context.Context,
		orgId uuid.UUID,
		qry string,
	) (
		[]hoteldtos.HotelGuestResponseDTO,
		error)

	CreateGuestInHotel(
		ctx context.Context,
		orgId uuid.UUID,
		req hoteldtos.CreateHotelGuestsRequest,
	) error
}
