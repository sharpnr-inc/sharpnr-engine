package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (d *hotelRepository) GetChargesByBookingID(
	ctx context.Context,
	bookingId uuid.UUID,
) ([]hospitalitymodels.HotelBookingCharge, error) {

	var charges []hospitalitymodels.HotelBookingCharge

	err := d.db.WithContext(ctx).
		Where("hotel_room_booking_id = ?", bookingId).
		Order("created_at ASC").
		Find(&charges).Error

	if err != nil {
		return nil, err
	}

	return charges, nil
}
