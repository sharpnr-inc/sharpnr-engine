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

func (h *hotelRepository) CreatePriceForRoom(ctx context.Context, roomPrice *hospitalitymodels.HotelRoomPrice) error {
	// 1️⃣ Basic validation
	if roomPrice.Price <= 0 {
		return errors.New("room price must be greater than zero")
	}

	if roomPrice.EffectiveFrom.IsZero() {
		return errors.New("effective_from is required")
	}

	// 2️⃣ Overlap protection (CRITICAL)
	var count int64

	err := h.db.WithContext(ctx).
		Model(&hospitalitymodels.HotelRoomPrice{}).
		Where("room_id = ?", roomPrice.RoomID).
		Where("effective_from <= ?", roomPrice.EffectiveFrom).
		Where(
			h.db.
				Where("effective_to IS NULL").
				Or("effective_to >= ?", roomPrice.EffectiveFrom),
		).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("room price already exists for this date")
	}

	// 3️⃣ Insert
	return h.db.WithContext(ctx).Create(roomPrice).Error
}

// DeleteRoom implements HotelRepository.
func (h *hotelRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	panic("unimplemented")
}
