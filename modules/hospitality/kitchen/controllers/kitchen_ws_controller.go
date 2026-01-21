package kitchencontrollers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (c *kitchenController) KitchenWebSocket(ctx *gin.Context) {
	kitchenID := ctx.Param("kitchenID")

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // ⚠️ In production, restrict by domain
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("❌ WebSocket upgrade failed:", err)
		return
	}

	fmt.Printf("✅ Kitchen connected: %s\n", kitchenID)
	c.kitchenService.RegisterKitchenConnection(kitchenID, conn)
}
