package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *hotelRepository) CreateRoomBooking(ctx context.Context, booking *hospitalitymodels.HotelRoomBooking) error {
	if err := d.db.WithContext(ctx).Create(booking).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) CreateOrAddBookingPayment(ctx context.Context, bookingPayment *hospitalitymodels.HotelRoomBookingPayment) error {
	if err := d.db.WithContext(ctx).Create(bookingPayment).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) CreateCharges(ctx context.Context, charges *hospitalitymodels.HotelBookingCharge) error {
	if err := d.db.WithContext(ctx).Create(charges).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) GetBookingDetailsByRoomId(
	ctx context.Context,
	roomId uuid.UUID,
) (*hospitalitymodels.HotelRoomBooking, error) {

	var booking hospitalitymodels.HotelRoomBooking

	err := d.db.WithContext(ctx).
		Preload("Room").
		Preload("PrimaryGuest").
		// Preload("Guests").
		Preload("Charges").
		Preload("Payments").
		Preload("Orders").
		Preload("Orders.Items").
		Preload("Services").
		Preload("Services.Items").
		Where("room_id = ? AND status IN ('RESERVED', 'CHECKED_IN', 'BOOKED')", roomId). // only active bookings
		First(&booking).Error

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (d *hotelRepository) GetBookingByBookingId(ctx context.Context, bookingId uuid.UUID) (*hospitalitymodels.HotelRoomBooking, error) {
	var booking hospitalitymodels.HotelRoomBooking

	err := d.db.WithContext(ctx).
		Where("id = ?", bookingId).
		First(&booking).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &booking, nil
}
