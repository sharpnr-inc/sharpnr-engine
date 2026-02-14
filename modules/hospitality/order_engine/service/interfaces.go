package hospitality_order_engine_service

import (
	"context"
	hospitality_order_engine_contracts "sharpnr-engine/modules/hospitality/order_engine/contracts"
)

type HospitalityOrderEngineServiceInterface interface {
	CreateOrder(ctx context.Context, input *hospitality_order_engine_contracts.CreateOrderInput) error
	// GetOrders(ctx context.Context, ownersID string) ([]*hospitality_order_engine_contracts.OrderOutput, error)
}
