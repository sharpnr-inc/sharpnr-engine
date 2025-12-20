package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

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

func (h *hotelRepository) UpdateRoomStatus(
	ctx context.Context,
	hotelId uuid.UUID,
	roomId uuid.UUID,
	status string,
) error {

	result := h.db.WithContext(ctx).
		Model(&hospitalitymodels.Room{}).
		Where("id = ? AND hotel_id = ?", roomId, hotelId).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("room not found or status unchanged")
	}

	return nil
}

// DeleteRoom implements HotelRepository.
func (h *hotelRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	panic("unimplemented")
}

func (d *hotelRepository) AddCurrentBookingIdToRoom(
	ctx context.Context,
	roomId uuid.UUID,
	bookingId uuid.UUID,
) error {

	result := d.db.WithContext(ctx).
		Model(&hospitalitymodels.Room{}).
		Where("id = ?", roomId).
		Updates(map[string]interface{}{
			"current_booking_id": bookingId,
			"updated_at":         time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("room not found or booking already attached")
	}

	return nil
}
