package hotelrepo

import (
	"context"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
)

type HotelRepositoryInterface interface {

	//--------------------------------------------------------------
	// Hotel
	//--------------------------------------------------------------
	CreateHotel(ctx context.Context, hotel *hotel_models.Hotel) error
	GetHotelByID(ctx context.Context, id uuid.UUID) (*hotel_models.Hotel, error)
	GetHotelByOrgID(ctx context.Context, orgid uuid.UUID) ([]*hotel_models.Hotel, error)
	UpdateHotelStatus(ctx context.Context, hotelID uuid.UUID, status string) error
	GetHotelsByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*hotel_models.Hotel, error)

	//-------------------------------------------------------------
	//   Floors
	// ------------------------------------------------------------

	CreateFloor(ctx context.Context, floor *hotel_models.HotelFloor) error
	GetFloorByID(ctx context.Context, id string) (*hotel_models.HotelFloor, error)
	GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hotel_models.HotelFloor, error)
	UpdateFloor(ctx context.Context, floor *hotel_models.HotelFloor) error
	DeleteFloor(ctx context.Context, floorID uuid.UUID) error

	//-------------------------------------------------------------
	//   Rooms
	// ------------------------------------------------------------
	CreateRoom(ctx context.Context, room *hotel_models.HotelRoom) error
	GetRoomByID(ctx context.Context, roomID uuid.UUID) (*hotel_models.HotelRoom, error)
	GetRoomByStatus(ctx context.Context, roomStatus string) ([]*hotel_models.HotelRoom, error)

	GetRoomsByStatus(ctx context.Context, status string) ([]*hotel_models.HotelRoom, error)
	UpdateRoomStatus(ctx context.Context, roomID uuid.UUID, status string) error

	CreateRoomDetails(ctx context.Context, roomDetail *hotel_models.HotelRoomDetail) error
	GetRoomDetail(ctx context.Context, roomID uuid.UUID) (*hotel_models.HotelRoomDetail, error)

	GetRoomRate(ctx context.Context, roomID uuid.UUID) (*hotel_models.HotelRoomRate, error)
	CreateRoomRate(ctx context.Context, roomRate *hotel_models.HotelRoomRate) error

	GetHotelRoomsByFloorID(ctx context.Context, floorID uuid.UUID) ([]*hotel_models.HotelRoom, error)
	DeleteRoom(ctx context.Context, roomID uuid.UUID) error

	//-------------------------------------------------------------
	//   Guests
	// ------------------------------------------------------------
	CreatePrimaryGuest(ctx context.Context, guest *hotel_models.HotelGuest) error
	GetHotelGuests(ctx context.Context, hotelId uuid.UUID) ([]hotel_models.HotelGuest, error)
	SearchHotelGuest(ctx context.Context, hotelId uuid.UUID, qry string) ([]hotel_models.HotelGuest, error)
	GetHotelGuestById(ctx context.Context, guestId uuid.UUID) (hotel_models.HotelGuest, error)
}

/*******************************************
 * this is the Interface for House Keeping *
 *******************************************/
type HotelHouseKeepingRepositoryInterface interface {

	// TODO :- This need To Implemented
}
