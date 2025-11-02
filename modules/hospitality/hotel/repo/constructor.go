package hotelrepo

import (
	"sharpnr-engine/shared/database"

	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

// CreateFloor implements HotelRepository.
func (h *hotelRepository) CreateFloor() {
	panic("unimplemented")
}

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom() {
	panic("unimplemented")
}

// EditRoom implements HotelRepository.
func (h *hotelRepository) EditRoom() {
	panic("unimplemented")
}

// GetRoomDetails implements HotelRepository.
func (h *hotelRepository) GetRoomDetails() {
	panic("unimplemented")
}

func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{
		db: database.DB,
	}
}
