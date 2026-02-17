package restaurantservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *restaurantService) StartDineSession(ctx context.Context,
	restaurantID uuid.UUID,
	orgID uuid.UUID,
	employeeID uuid.UUID,
	tableID uuid.UUID,
) (uuid.UUID, error) { // Removed the named 'diningSessionID' for cleaner return

	var sessionID uuid.UUID

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// 1. Generate the ID first so we can capture it
		sessionID = uuid.New()

		session := &restaurant_models.RestaurantDiningSession{
			ID:           sessionID,
			RestaurantID: restaurantID,
			TableID:      tableID,
			Status:       "RUNNING",
			StartedAt:    time.Now(),
			CreatedAt:    time.Now(),
		}

		if err := s.restaurantRepo.CreateRestaurantDiningSession(txCtx, session); err != nil {
			return err
		}

		return nil
	})

	// 2. Check if the transaction failed
	if err != nil {
		return uuid.Nil, err
	}

	// 3. Return the captured ID
	return sessionID, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (s *restaurantService) AddOrdersInSession(ctx context.Context) {
	panic("Unimplemented")
}
