package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
)

func (d *hotelRepository) GetPaymentsByBookingId(
	ctx context.Context,
	bookingId uuid.UUID,
) ([]hospitalitymodels.HotelRoomBookingPayment, error) {

	var payments []hospitalitymodels.HotelRoomBookingPayment

	err := d.db.WithContext(ctx).
		Where("hotel_room_booking_id = ?", bookingId).
		Order("created_at ASC").
		Find(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (d *hotelRepository) AddPaymentsInBooking(
	ctx context.Context,
	payment *hospitalitymodels.HotelRoomBookingPayment,
) error {

	if payment == nil {
		return errors.New("payment cannot be nil")
	}

	now := time.Now()
	payment.CreatedAt = now

	if err := d.db.
		WithContext(ctx).
		Create(payment).
		Error; err != nil {
		return err
	}

	return nil
}
