package hotelservices

import (
	"context"
	hotel_contracts "sharpnr-engine/property_business_ops/hospitality/hotel/contracts"
	hoteldtos "sharpnr-engine/property_business_ops/hospitality/hotel/dtos"

	"github.com/google/uuid"
)

type HotelServiceInterface interface {
	SetupDefaultHotel(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (uuid.UUID, error)
	// GetHotelDetails(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (*hotel_contracts.HotelsReader, error)
	GetHotelLayout(ctx context.Context, hotelID uuid.UUID) (res *hoteldtos.HotelLayoutResponseDTO, err error)

	GetHotelByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]hotel_contracts.HotelReader, error)

	GetHotelsByOrgID(ctx context.Context, orgID, emplID uuid.UUID) ([]hotel_contracts.Hotels, error)
	SetupFloors(ctx context.Context, req *hotel_contracts.FloorInput) (floorID uuid.UUID, err error)
	SetUpLayout(ctx context.Context, layout []hotel_contracts.FloorInput) error
	SetupRoom(ctx context.Context, req *hotel_contracts.RoomInput) (hotelRoomID uuid.UUID, err error)
	AddHotelGenDetails(ctx context.Context, input *hotel_contracts.AddHotelGeneralInfoInput) error
}
