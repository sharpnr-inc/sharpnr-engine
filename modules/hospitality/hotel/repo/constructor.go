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

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom(ctx context.Context, room *hospitalitymodels.Room) error {
	panic("unimplemented")
}

// CreateFloorInHotel implements HotelRepository.
func (h *hotelRepository) CreateFloorInHotel() {
	panic("unimplemented")
}

// DeleteFloor implements HotelRepository.
func (h *hotelRepository) DeleteFloor(ctx context.Context, floorID uuid.UUID) error {
	panic("unimplemented")
}

// DeleteRoom implements HotelRepository.
func (h *hotelRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	panic("unimplemented")
}

// GetAllRoomInFloor implements HotelRepository.
func (h *hotelRepository) GetAllRoomInFloor() {
	panic("unimplemented")
}

// GetFloorsByHotel implements HotelRepository.
func (h *hotelRepository) GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hospitalitymodels.HotelFloor, error) {
	panic("unimplemented")
}

// GetHotelByID implements HotelRepository.
func (h *hotelRepository) GetHotelByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.Hotel, error) {
	panic("unimplemented")
}

// GetRoomByID implements HotelRepository.
func (h *hotelRepository) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*hospitalitymodels.Room, error) {
	panic("unimplemented")
}

// UpdateFloor implements HotelRepository.
func (h *hotelRepository) UpdateFloor(ctx context.Context, floor *hospitalitymodels.HotelFloor) error {
	panic("unimplemented")
}

// UpdateHotel implements HotelRepository.
func (h *hotelRepository) UpdateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error {
	panic("unimplemented")
}

// UpdateRoom implements HotelRepository.
func (h *hotelRepository) UpdateRoom(ctx context.Context, room *hospitalitymodels.Room) error {
	panic("unimplemented")
}

func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{
		db: db,
	}
}
