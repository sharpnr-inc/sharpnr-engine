package hotelservices

import (
	"context"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/dtos"

	"github.com/google/uuid"
)

type HotelServiceInterface interface {
	SetupDefaultHotel(ctx context.Context)

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

	AddPaymentForBooking(ctx context.Context, orgID uuid.UUID, req hoteldtos.AddPaymentsForBookingRequest_dto) error
}
