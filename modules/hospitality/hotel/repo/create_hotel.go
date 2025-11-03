package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

func (h *hotelRepository) CreateHotel(ctx context.Context, room *hospitalitymodels.Hotel) (*hospitalitymodels.Hotel, error) {
	// Example logic:
	if err := h.db.WithContext(ctx).Create(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}
