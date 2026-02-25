package restaurantservices

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *restaurantService) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		
		if err := s.restaurantRepo.UpdateOrderStatus(ctx, orderID, status); err != nil {
			return err
		}
		return nil

	})
}
