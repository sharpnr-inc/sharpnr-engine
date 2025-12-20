package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *hotelRepository) CreateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error {
	// Example logic:
	if err := h.db.WithContext(ctx).Create(hotel).Error; err != nil {
		return err
	}

	return nil
}

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
		Preload("Organization").
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

func (r *hotelRepository) UpdateHotelFields(
	ctx context.Context,
	hotelID uuid.UUID,
	updates map[string]interface{},
) error {

	if len(updates) == 0 {
		return nil
	}

	return r.db.
		WithContext(ctx).
		Model(&hospitalitymodels.Hotel{}).
		Where("id = ?", hotelID).
		Updates(updates).
		Error
}
