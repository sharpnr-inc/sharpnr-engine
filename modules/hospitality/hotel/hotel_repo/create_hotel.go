package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (h *hotelRepository) CreateHotel(ctx context.Context, hotel *hospitalitymodels.Hotel) error {
	// Example logic:
	if err := h.db.WithContext(ctx).Create(hotel).Error; err != nil {
		return err
	}

	return nil
}
