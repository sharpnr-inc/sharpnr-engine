package kitchenrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

type KitchenRepositoryInterface interface {
	//This function for Creating the Create Kitchen
	// this function used in the Org factory to create default Kitchen
	CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error

	GetOrders() ([]interface{}, error)

	UpdateOrder(orderID string, status string) error
}
