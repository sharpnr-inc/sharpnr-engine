package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *hotelRepository) CreateFloorInHotel(ctx context.Context, floor *hospitalitymodels.HotelFloor) error {
	if err := h.db.WithContext(ctx).Create(floor).Error; err != nil {
		return err
	}
	return nil
}

func (h *hotelRepository) GetFloorByID(ctx context.Context, id string) (*hospitalitymodels.HotelFloor, error) {
	var floor hospitalitymodels.HotelFloor
	err := h.db.WithContext(ctx).
		Preload("Hotel").
		First(&floor, "id = ?", id).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &floor, nil
}

// DeleteFloor implements HotelRepository.
func (h *hotelRepository) DeleteFloor(ctx context.Context, floorID uuid.UUID) error {
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

// UpdateFloor implements HotelRepository.
func (h *hotelRepository) UpdateFloor(ctx context.Context, floor *hospitalitymodels.HotelFloor) error {
	panic("unimplemented")
}
