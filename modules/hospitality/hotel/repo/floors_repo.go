package hotelrepo

import (
	"context"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *hotelRepository) CreateFloor(ctx context.Context, floor *hotel_models.HotelFloor) error {
	if err := h.db.WithContext(ctx).Create(floor).Error; err != nil {
		return err
	}
	return nil
}

func (h *hotelRepository) GetFloorByID(ctx context.Context, floorID string) (*hotel_models.HotelFloor, error) {
	var floor hotel_models.HotelFloor
	err := h.db.WithContext(ctx).
		First(&floor, "id = ?", floorID).
		Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &floor, nil
}

// GetFloorsByHotel implements HotelRepository.
func (h *hotelRepository) GetFloorsByHotelID(ctx context.Context, hotelID uuid.UUID) ([]*hotel_models.HotelFloor, error) {
	var floors []*hotel_models.HotelFloor

	// Use .Find to retrieve all matching records into the slice
	err := h.db.WithContext(ctx).
		Where("hotel_id = ?", hotelID).
		Order("floor_number ASC"). // Optional: useful for floors
		Find(&floors).Error

	if err != nil {
		// Find returns an empty slice and nil error if no records are found.
		// We only return an error if a real database issue occurs.
		return nil, err
	}

	return floors, nil
}

// UpdateFloor implements HotelRepository.
func (h *hotelRepository) UpdateFloor(ctx context.Context, floor *hotel_models.HotelFloor) error {
	panic("unimplemented")
}

// DeleteFloor implements HotelRepository.
func (h *hotelRepository) DeleteFloor(ctx context.Context, floorID uuid.UUID) error {
	panic("unimplemented")
}

// GetFloorsByHotel implements HotelRepositoryInterface.
func (h *hotelRepository) GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hotel_models.HotelFloor, error) {
	panic("unimplemented")
}
