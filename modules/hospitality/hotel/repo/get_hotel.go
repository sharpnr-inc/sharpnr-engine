package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetHotelByID implements HotelRepository.
func (h *hotelRepository) GetHotelByID(ctx context.Context, id uuid.UUID) (*hospitalitymodels.Hotel, error) {
	var hotel hospitalitymodels.Hotel

	err := h.db.WithContext(ctx).
		Preload("Floors").
		// If you also want rooms on each floor, uncomment the next line:
		// Preload("Floors.Rooms").
		First(&hotel, "id = ?", id).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &hotel, nil
}

// GetHotelByID implements HotelRepository.
func (h *hotelRepository) GetHotelByOrgID(ctx context.Context, orgid uuid.UUID) (*hospitalitymodels.Hotel, error) {
	var hotel hospitalitymodels.Hotel

	err := h.db.WithContext(ctx).
		Preload("Floors.Rooms"). // ✅ Preload nested rooms
		Where("organization_id = ?", orgid).
		First(&hotel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &hotel, nil
}
