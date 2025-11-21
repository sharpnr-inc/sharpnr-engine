package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

// GetAllRoomInFloor implements HotelRepository.
func (h *hotelRepository) GetAllRoomInFloor() {
	panic("unimplemented")
}

// GetFloorsByHotel implements HotelRepository.
func (h *hotelRepository) GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hospitalitymodels.HotelFloor, error) {
	panic("unimplemented")
}

// UpdateFloor implements HotelRepository.
func (h *hotelRepository) UpdateFloor(ctx context.Context, floor *hospitalitymodels.HotelFloor) error {
	panic("unimplemented")
}

func (h *hotelRepository) UpdateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error {
	return h.db.WithContext(ctx).Save(hotel).Error
}

func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{
		db: db,
	}
}
