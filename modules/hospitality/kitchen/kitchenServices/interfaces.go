package kitchenservices

import "github.com/gorilla/websocket"

type KitchenServiceInterface interface {
	GetOrders() ([]interface{}, error)
	UpdateOrder(orderID string, status string) error
	RegisterKitchenConnection(kitchenID string, conn *websocket.Conn)
	BroadcastOrderUpdate(kitchenID string, orderData interface{})
}
