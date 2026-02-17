package restaurantservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	restaurantrepo "sharpnr-engine/modules/hospitality/restaurant/repo"
	"sharpnr-engine/sero"
)

type restaurantService struct {
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface
	seroEngine     *sero.SEROEngine
	txManager      dbtx.TxManager
}

// CreateOrder implements RestaurantServiceInterface.
func (r *restaurantService) CreateOrder(ctx context.Context) {
	panic("unimplemented")
}

// AddRestaurantGeneralDetails implements RestaurantServiceInterface.
func (r *restaurantService) AddRestaurantGeneralDetails(ctx context.Context) {
	panic("unimplemented")
}

func NewRestaurantService(
	restaurantRepo restaurantrepo.RestaurantRepositoryInterface,
	seroEngine *sero.SEROEngine,
	txManager dbtx.TxManager,
) *restaurantService {
	return &restaurantService{
		restaurantRepo: restaurantRepo,
		seroEngine:     seroEngine,
		txManager:      txManager,
	}

}
