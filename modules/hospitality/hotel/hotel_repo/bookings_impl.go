package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (d *hotelRepository) CreateRoomBooking(ctx context.Context, booking *hospitalitymodels.HotelRoomBooking) error {
	if err := d.db.WithContext(ctx).Create(booking).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) CreateOrAddBookingPayment(ctx context.Context, bookingPayment *hospitalitymodels.BookingPayment) error {
	if err := d.db.WithContext(ctx).Create(bookingPayment).Error; err != nil {
		return err
	}
	return nil
}

func (d *hotelRepository) CreateCharges(ctx context.Context, charges *hospitalitymodels.BookingCharge) error {
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
