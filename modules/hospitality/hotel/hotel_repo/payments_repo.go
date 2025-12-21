package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

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
	payemnt *hospitalitymodels.HotelRoomBookingPayment,
) error {
	panic("unimplemented")
}
