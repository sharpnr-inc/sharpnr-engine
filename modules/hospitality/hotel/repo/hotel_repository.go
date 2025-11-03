package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

// CreateRoom implements HotelRepository.
func (h *hotelRepository) CreateRoom(ctx context.Context, room *hospitalitymodels.Room) (*hospitalitymodels.Room, error) {
	// Example logic:
	if err := h.db.WithContext(ctx).Create(room).Error; err != nil {
		return nil, err
	}
	return room, nil
}
