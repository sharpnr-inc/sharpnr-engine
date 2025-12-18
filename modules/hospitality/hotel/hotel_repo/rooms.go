package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

func (h *hotelRepository) UpdateRoom(ctx context.Context, fields map[string]interface{}, roomId uuid.UUID) error {
	return h.db.WithContext(ctx).
		Model(&hospitalitymodels.Room{}).
		Where("id = ?", roomId).
		Updates(fields).Error
}
