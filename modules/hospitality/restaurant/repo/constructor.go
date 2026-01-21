package restaurantrepo

import (
	"context"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type restaurantRepo struct {
	db *gorm.DB
}

// GetAllTablesInFloor implements RestaurantRepositoryInterface.
func (r *restaurantRepo) GetAllTablesInFloor(ctx context.Context, floorID uuid.UUID) ([]*restaurant_models.RestaurantTable, error) {
	panic("unimplemented")
}

// GetDiningSessionByStatus implements RestaurantRepositoryInterface.
func (r *restaurantRepo) GetDiningSessionByStatus(ctx context.Context, restaurantID uuid.UUID, status string) ([]*restaurant_models.RestaurantDiningSession, error) {
	panic("unimplemented")
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
