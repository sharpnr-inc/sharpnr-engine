package hospitality_order_engine_repository

import (
	"context"
	hospitality_order_engine_models "sharpnr-engine/modules/hospitality/order_engine/models"

	"github.com/google/uuid"
)

type HospitalityOrderEngineRepositoryInterface interface {
	CreateOrder(ctx context.Context, order *hospitality_order_engine_models.HospitalityOrder) error
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*hospitality_order_engine_models.HospitalityOrder, error)
	GetOrdersByOwnersID(ctx context.Context, ownersID uuid.UUID) ([]*hospitality_order_engine_models.HospitalityOrder, error)
	GetOrdersByPropertyID(ctx context.Context, propertyID uuid.UUID) ([]*hospitality_order_engine_models.HospitalityOrder, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error

	CreateOrderItem(ctx context.Context, orderItem *hospitality_order_engine_models.HospitalityOrderItem) error
	GetOrderItemByOrderID(ctx context.Context, orderID uuid.UUID) ([]*hospitality_order_engine_models.HospitalityOrderItem, error)

	CreateOrderItemVariant(ctx context.Context, itemVariant *hospitality_order_engine_models.HospitalityOrderItemVariant) error
}
