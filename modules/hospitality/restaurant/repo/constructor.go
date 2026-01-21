package restaurantrepo

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type restaurantRepo struct {
	db *gorm.DB
}

// UpdateRestaurantOrderItemKotReference implements RestaurantRepository.
func (r *restaurantRepo) UpdateRestaurantOrderItemKotReference(ctx context.Context, orderItemID uuid.UUID, ticketID uuid.UUID, ticketItemID uuid.UUID) error {
	panic("unimplemented")
}

func NewRestaurantRepository(db *gorm.DB) *restaurantRepo {
	return &restaurantRepo{
		db: db,
	}
}
