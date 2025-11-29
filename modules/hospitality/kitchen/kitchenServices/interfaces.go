package kitchenservices

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type KitchenServiceInterface interface {
	GetOrdersList(ctx context.Context, orgId uuid.UUID) ([]interface{}, error)
	UpdateOrder(orderID string, status string) error
	RegisterKitchenConnection(kitchenID string, conn *websocket.Conn)
	BroadcastOrderUpdate(kitchenID string, orderData interface{})
}
