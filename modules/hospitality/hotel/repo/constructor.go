package hotelrepo

import (
	"context"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

// GetFloorsByHotel implements HotelRepositoryInterface.
func (h *hotelRepository) GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hotel_models.HotelFloor, error) {
	panic("unimplemented")
}

//-----------------------------------------------------------------------------------------------------------------------------------//

// NewHotelRepository creates a new instance of hotelRepository.
//
// It initializes the repository with a gorm.DB connection and returns
// an implementation of the HotelRepository interface.
//
// This constructor should be used by the service layer or dependency
// injection setup to ensure a single, consistent database connection
// is used across all hotel-related persistence operations.
func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{
		db: db,
	}
}

// -----------------------------------------------------------------------------------------------------------------------------------//

type hotelHouseKeepingRepository struct {
	db *gorm.DB
}

func NewHotelHouseKeepingRepository(db *gorm.DB) *hotelHouseKeepingRepository {
	return &hotelHouseKeepingRepository{
		db: db,
	}
}
