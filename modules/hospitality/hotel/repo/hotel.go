package hotelrepo

import (
	"context"
	"errors"
	"fmt"
	hotel_contracts "sharpnr-engine/modules/hospitality/hotel/contracts"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create Hotel
func (h *hotelRepository) CreateHotel(ctx context.Context, hotel *hotel_models.Hotel) error {

	if err := h.db.WithContext(ctx).Create(hotel).Error; err != nil {
		return err
	}
	return nil
}

/* --------------------------------- <****> --------------------------------- */

// Update Hotels Details
func (r *hotelRepository) UpdateHotelGenDetails(ctx context.Context, updates *hotel_contracts.HotelMainGenDetailsInput) error {
	// We use a map to ensure GORM updates all fields, even if they are 'false' or empty strings.
	updateData := map[string]interface{}{
		"name":          updates.Name,
		"is_configured": updates.IsConfigured,
		"updated_at":    time.Now(),
	}

	result := r.db.WithContext(ctx).
		Model(&hotel_models.Hotel{}).
		Where("id = ?", updates.HotelID).
		Updates(updateData)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no hotel found with the given ID")
	}

	return nil
}

/* --------------------------------- <****> --------------------------------- */

// This return one Hotel by the its own ID hotel ID
func (h *hotelRepository) GetHotelByID(ctx context.Context, hotelID uuid.UUID) (*hotel_models.Hotel, error) {
	var hotel hotel_models.Hotel
	err := h.db.WithContext(ctx).
		First(&hotel, "id = ?", hotelID).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &hotel, nil
}

func (h *hotelRepository) UpdateHotelStatus(ctx context.Context, hotelID uuid.UUID, status string) error {
	// We target the Hotel model and update only the "status" column.
	result := h.db.WithContext(ctx).
		Model(&hotel_models.Hotel{}).
		Where("id = ?", hotelID).
		Update("status", status)

	// Check for database errors (connection issues, constraints, etc.)
	if result.Error != nil {
		return result.Error
	}

	// Optional: Check if the hotel actually existed.
	// GORM's Update returns nil error even if no rows matched the ID.
	if result.RowsAffected == 0 {
		return fmt.Errorf("hotel not found with ID: %s", hotelID)
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetHotelByID implements HotelRepository.
func (h *hotelRepository) GetHotelsByOrgID(ctx context.Context, orgID uuid.UUID) ([]hotel_models.Hotel, error) {
	// 1. Initialize a slice of pointers to match your return type
	var hotels []hotel_models.Hotel

	// 2. Use .Find() instead of .First() to retrieve multiple records
	err := h.db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Find(&hotels).Error

	if err != nil {
		// GORM .Find() typically does not return ErrRecordNotFound;
		// it just returns an empty slice. But we keep the check for safety.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// 3. 'hotels' is already a slice of pointers ([]*Hotel)
	return hotels, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (d *hotelRepository) GetHotelsByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]hotel_models.Hotel, error) {

	var hotels []hotel_models.Hotel

	err := d.db.WithContext(ctx).Where("property_id = ?", propertyID).Find(&hotels).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return hotels, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetHotelByOrgID implements HotelRepositoryInterface.
func (h *hotelRepository) GetHotelByOrgID(
	ctx context.Context,
	orgid uuid.UUID,
) ([]*hotel_models.Hotel, error) {

	hotels := make([]*hotel_models.Hotel, 0)

	err := h.db.
		WithContext(ctx).
		Where("organization_id = ?", orgid).
		Find(&hotels).Error

	if err != nil {
		return nil, err
	}

	return hotels, nil
}

/* ---------------------------------- <***> --------------------------------- */
