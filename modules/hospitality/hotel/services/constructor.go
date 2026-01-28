package hotelservices

import (
	"context"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/dtos"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/repo"
	"sharpnr-engine/sere"

	"github.com/google/uuid"
)

type hotelService struct {
	sere      *sere.SERE // 🔥 event engine
	hotelRepo hotelrepo.HotelRepositoryInterface
}

// AddPaymentForBooking implements HotelServiceInterface.
func (h *hotelService) AddPaymentForBooking(ctx context.Context, orgID uuid.UUID, req hoteldtos.AddPaymentsForBookingRequest_dto) error {
	panic("unimplemented")
}

// CreateGuestInHotel implements HotelServiceInterface.
func (h *hotelService) CreateGuestInHotel(ctx context.Context, orgId uuid.UUID, req hoteldtos.CreateHotelGuestsRequest) error {
	panic("unimplemented")
}

// SearchGuests implements HotelServiceInterface.
func (h *hotelService) SearchGuests(ctx context.Context, orgId uuid.UUID, qry string) ([]hoteldtos.HotelGuestResponseDTO, error) {
	panic("unimplemented")
}

// SetupDefaultHotel implements HotelServiceInterface.
func (h *hotelService) SetupDefaultHotel(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (uuid.UUID, error) {
	panic("unimplemented")
}

func NewHotelService(
	sere *sere.SERE,
	hotelRepo hotelrepo.HotelRepositoryInterface,

) *hotelService {
	return &hotelService{
		sere:      sere,
		hotelRepo: hotelRepo,
	}
}
