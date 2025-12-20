package hotelrepo

import (
	"context"
	"errors"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (d *hotelRepository) GetBillByBookingId(
	ctx context.Context,
	bookingId uuid.UUID,
) (*hospitalitymodels.HotelBill, error) {

	var bill hospitalitymodels.HotelBill

	err := d.db.WithContext(ctx).
		Where("hotel_room_booking_id = ?", bookingId).
		First(&bill).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &bill, nil
}
