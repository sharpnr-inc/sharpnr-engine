package hotelrepo

import (
	"context"
	"errors"
	"fmt"
	hotel_models "sharpnr-engine/business_ops/hospitality/hotel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom(ctx context.Context, room *hotel_models.HotelRoom) error {
	return h.db.WithContext(ctx).Create(room).Error
}

func (h *hotelRepository) GetRoomByStatus(ctx context.Context, roomStatus string) ([]*hotel_models.HotelRoom, error) {
	var rooms []*hotel_models.HotelRoom

	// We filter by the "status" column.
	// Ensure "status" matches the column name in your database.
	err := h.db.WithContext(ctx).
		Where("status = ?", roomStatus).
		Find(&rooms).Error

	if err != nil {
		// Find returns an empty slice and no error if no rooms match the status.
		return nil, err
	}

	return rooms, nil
}

func (h *hotelRepository) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*hotel_models.HotelRoom, error) {
	var room hotel_models.HotelRoom

	err := h.db.WithContext(ctx).
		Where("id = ?", roomID).
		First(&room).Error

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (h *hotelRepository) UpdateRoomStatus(ctx context.Context, roomID uuid.UUID, status string) error {
	// We use .Model() to target the table and .Update() for a single column change.
	// This is safer than .Save() as it avoids accidental overwrites of other fields.
	result := h.db.WithContext(ctx).
		Model(&hotel_models.HotelRoom{}).
		Where("id = ?", roomID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	// Optional: Check if any row was actually updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("room not found or status already set to %s", status)
	}

	return nil
}

func (h *hotelRepository) GetRoomsByStatus(ctx context.Context, status string) ([]*hotel_models.HotelRoom, error) {
	var rooms []*hotel_models.HotelRoom

	// We filter by the 'status' column.
	// .Find() expects a slice and returns nil error (not ErrRecordNotFound) if empty.
	err := h.db.WithContext(ctx).
		Where("status = ?", status).
		Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (h *hotelRepository) CreateRoomDetails(ctx context.Context, roomDetail *hotel_models.HotelRoomDetail) error {
	// We use .Create() to insert the record.
	// GORM will automatically populate fields like CreatedAt, UpdatedAt,
	// and the ID if it's a serial/auto-incrementing field.
	if err := h.db.WithContext(ctx).Create(roomDetail).Error; err != nil {
		return err
	}

	return nil
}

func (h *hotelRepository) GetRoomDetail(ctx context.Context, roomID uuid.UUID) (*hotel_models.HotelRoomDetail, error) {
	var detail hotel_models.HotelRoomDetail

	// We search the 'hotel_room_details' table where the 'room_id' matches.
	// .First() is used because typically a room has only one detail record.
	err := h.db.WithContext(ctx).
		Where("room_id = ?", roomID).
		First(&detail).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Returning nil, nil is common if "no details" isn't a hard error,
			// but returning the error allows the service to choose (e.g., 404).
			return nil, err
		}
		return nil, err
	}

	return &detail, nil
}

func (h *hotelRepository) GetRoomRate(ctx context.Context, roomID uuid.UUID) (*hotel_models.HotelRoomRate, error) {
	var rate hotel_models.HotelRoomRate

	// We filter the rates table by room_id.
	// Using .First() ensures we get the most relevant single record.
	err := h.db.WithContext(ctx).
		Where("room_id = ?", roomID).
		First(&rate).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return the specific GORM error so the service layer
			// knows the rate hasn't been set yet.
			return nil, err
		}
		return nil, err
	}

	return &rate, nil
}

func (h *hotelRepository) CreateRoomRate(ctx context.Context, roomRate *hotel_models.HotelRoomRate) error {
	// .Create handles the insertion of the record.
	// It will return an error if a constraint is violated (e.g., the RoomID doesn't exist).
	if err := h.db.WithContext(ctx).Create(roomRate).Error; err != nil {
		return err
	}

	return nil
}

func (h *hotelRepository) CreateRoomRateHistory(ctx context.Context, roomRateHistory *hotel_models.HotelRoomRateHistory) error {
	if err := h.db.WithContext(ctx).Create(roomRateHistory).Error; err != nil {
		return err
	}

	return nil
}

func (h *hotelRepository) GetHotelRoomsByFloorID(ctx context.Context, floorID uuid.UUID) ([]*hotel_models.HotelRoom, error) {
	var rooms []*hotel_models.HotelRoom

	// 1. Target the HotelRoom model
	// 2. Filter by the floor_id column
	err := h.db.WithContext(ctx).
		Where("hotel_floor_id = ?", floorID).
		Find(&rooms).Error

	if err != nil {
		// We return the error only if it's a database failure.
		// If no rooms are found, GORM returns nil error and an empty slice.
		return nil, err
	}

	return rooms, nil
}

func (h *hotelRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	panic("unimplemented")
}
