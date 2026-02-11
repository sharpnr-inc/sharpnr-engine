package hotelrepo

import (
	"context"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
)

func (r *hotelRepository) CreateHotelOperationalTiming(ctx context.Context, hotel_timing *hotel_models.HotelOperationsTiming) error {
	if err := r.db.WithContext(ctx).Create(&hotel_timing); err != nil {
		return err.Error
	}
	return nil
}
