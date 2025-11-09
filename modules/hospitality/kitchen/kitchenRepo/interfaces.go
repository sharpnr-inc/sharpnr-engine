package kitchenrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

type KitchenRepositoryInterface interface {
	CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error
	GetOrders() ([]interface{}, error)
	UpdateOrder(orderID string, status string) error
}
