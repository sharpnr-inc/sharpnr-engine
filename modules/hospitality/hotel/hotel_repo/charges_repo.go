package hotelrepo

import (
	"context"
	"errors"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"

	"github.com/google/uuid"
)

func (d *hotelRepository) GetChargesByBookingID(
	ctx context.Context,
	bookingId uuid.UUID,
) ([]hotel_models.HotelBookingCharge, error) {

	var charges []hotel_models.HotelBookingCharge

	err := d.db.WithContext(ctx).
		Where("hotel_room_booking_id = ?", bookingId).
		Order("created_at ASC").
		Find(&charges).Error

	if err != nil {
		return nil, err
	}

	return charges, nil
}

func (d *hotelRepository) AddChargesInBooking(
	ctx context.Context,
	charge *hotel_models.HotelBookingCharge,
) error {

	if charge == nil {
		return errors.New("charge is nil")
	}

	if err := d.db.WithContext(ctx).Create(charge).Error; err != nil {
		return err
	}

	return nil
}
