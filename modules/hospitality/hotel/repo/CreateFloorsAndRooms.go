package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom(ctx context.Context, room *hospitalitymodels.Room) error {
	panic("unimplemented")
}

func (h *hotelRepository) CreateFloorInHotel(ctx context.Context, floor *hospitalitymodels.HotelFloor) error {
	if err := h.db.WithContext(ctx).Create(floor).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFloor implements HotelRepository.
func (h *hotelRepository) DeleteFloor(ctx context.Context, floorID uuid.UUID) error {
	panic("unimplemented")
}

// DeleteRoom implements HotelRepository.
func (h *hotelRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	panic("unimplemented")
}
