package hotelrepo

import (
	"context"
	"errors"
	hotel_models "sharpnr-engine/property_business_ops/hospitality/hotel/models"

	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *hotelRepository) CreatePrimaryGuest(ctx context.Context, guest *hotel_models.HotelGuest) error {
	if err := d.db.WithContext(ctx).Create(guest).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) GetHotelGuests(ctx context.Context, hotelId uuid.UUID) ([]hotel_models.HotelGuest, error) {
	panic("Unimplemented")
}

func (d *hotelRepository) SearchHotelGuest(ctx context.Context, hotelId uuid.UUID, qry string) ([]hotel_models.HotelGuest, error) {
	var guests []hotel_models.HotelGuest
	db := d.db.WithContext(ctx).
		Where("hotel_id = ?", hotelId)

	// Normalize query
	q := strings.TrimSpace(qry)

	if q != "" {
		db = db.Where(
			d.db.Where("name ILIKE ?", "%"+q+"%").
				Or("mobile LIKE ?", q+"%"),
		)
	}

	err := db.
		Order("name ASC").
		Limit(10).
		Find(&guests).Error

	return guests, err
}

func (d *hotelRepository) GetHotelGuestById(ctx context.Context, guestId uuid.UUID) (hotel_models.HotelGuest, error) {
	var guest hotel_models.HotelGuest

	err := d.db.
		WithContext(ctx).
		Where("id = ?", guestId).
		First(&guest).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hotel_models.HotelGuest{}, err
		}
		return hotel_models.HotelGuest{}, err
	}

	return guest, nil
}

func (d *hotelRepository) CreateBookingGuest(ctx context.Context, bookingGuest *hotel_models.HotelGuest) error {
	if err := d.db.WithContext(ctx).Create(bookingGuest).Error; err != nil {
		return err
	}
	return nil
}
