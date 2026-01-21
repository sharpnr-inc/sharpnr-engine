package kitchenservices

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type KitchenConnection struct {
	ID   string
	Conn *websocket.Conn
}

var kitchenConnections = make(map[string][]*websocket.Conn)
var mu sync.Mutex

func (s *kitchenService) RegisterKitchenConnection(kitchenID string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	kitchenConnections[kitchenID] = append(kitchenConnections[kitchenID], conn)

	go s.listenToKitchen(kitchenID, conn)
}

func (s *kitchenService) listenToKitchen(kitchenID string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		mu.Lock()
		conns := kitchenConnections[kitchenID]
		for i, c := range conns {
			if c == conn {
				kitchenConnections[kitchenID] = append(conns[:i], conns[i+1:]...)
				break
			}
		}
		mu.Unlock()
		fmt.Printf("❌ Kitchen disconnected: %s\n", kitchenID)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("📩 Message from kitchen [%s]: %s\n", kitchenID, string(msg))
	}
}

// ----------------------------
// 🔥 Broadcast to Kitchens
// ----------------------------
func (s *kitchenService) BroadcastOrderUpdate(kitchenID string, orderData interface{}) {
	mu.Lock()
	defer mu.Unlock()

	conns := kitchenConnections[kitchenID]
	if len(conns) == 0 {
		fmt.Printf("⚠️ No active connections for kitchen %s\n", kitchenID)
		return
	}

	data, _ := json.Marshal(orderData)
	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Printf("❌ Write error: %v\n", err)
		}
	}
}
