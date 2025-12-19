package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom(ctx context.Context, room *hospitalitymodels.Room) error {
	return h.db.WithContext(ctx).Create(room).Error
}

// GetRoomByID implements HotelRepository.
func (h *hotelRepository) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*hospitalitymodels.Room, error) {
	var room hospitalitymodels.Room

	err := h.db.WithContext(ctx).
		Where("id = ?", roomID).
		First(&room).Error

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (h *hotelRepository) UpdateRoom(ctx context.Context, fields map[string]interface{}, roomId uuid.UUID) error {
	return h.db.WithContext(ctx).
		Model(&hospitalitymodels.Room{}).
		Where("id = ?", roomId).
		Updates(fields).Error
}

// DeleteRoom implements HotelRepository.
func (h *hotelRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	panic("unimplemented")
}
