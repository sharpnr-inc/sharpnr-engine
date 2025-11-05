package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom(ctx context.Context, room *hospitalitymodels.Room) error {
	if err := h.db.WithContext(ctx).Create(room).Error; err != nil {
		return err
	}
	return nil
}
