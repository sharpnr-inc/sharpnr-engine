package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (d *hotelRepository) CreatePrimaryGuest(ctx context.Context, guest *hospitalitymodels.HotelGuest) error {
	if err := d.db.WithContext(ctx).Create(guest).Error; err != nil {
		return err
	}
	return nil
}
