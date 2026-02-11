package hotelservices

import (
	"context"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"

	"github.com/google/uuid"
)

type HotelServiceInterface interface {
	SetupDefaultHotel(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (uuid.UUID, error)
	GetHotelDetails(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*hotel_contracts.HotelsReader, error)

	// SetupFloors(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID)
}
