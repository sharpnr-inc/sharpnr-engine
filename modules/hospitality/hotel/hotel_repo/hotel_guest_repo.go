package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *hotelRepository) CreatePrimaryGuest(ctx context.Context, guest *hospitalitymodels.HotelGuest) error {
	if err := d.db.WithContext(ctx).Create(guest).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) GetHotelGuests(ctx context.Context, hotelId uuid.UUID) ([]hospitalitymodels.HotelGuest, error) {
	panic("Unimplemented")
}

func (d *hotelRepository) SearchHotelGuest(ctx context.Context, hotelId uuid.UUID, qry string) ([]hospitalitymodels.HotelGuest, error) {
	var guests []hospitalitymodels.HotelGuest
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

func (d *hotelRepository) GetHotelGuestById(ctx context.Context, guestId uuid.UUID) (hospitalitymodels.HotelGuest, error) {
	var guest hospitalitymodels.HotelGuest

	err := d.db.
		WithContext(ctx).
		Where("id = ?", guestId).
		First(&guest).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return hospitalitymodels.HotelGuest{}, err
		}
		return hospitalitymodels.HotelGuest{}, err
	}

	return guest, nil
}

func (d *hotelRepository) CreateBookingGuest(ctx context.Context, bookingGuest *hospitalitymodels.BookingGuest) error {
	if err := d.db.WithContext(ctx).Create(bookingGuest).Error; err != nil {
		return err
	}
	return nil
}
