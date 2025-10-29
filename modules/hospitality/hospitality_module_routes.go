package hospitality

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// Upgrader for WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func OrdersWebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to establish WebSocket"})
		return
	}
	defer conn.Close()

	// Example: send dummy order every 5 seconds
	for {
		order := `{"event":"new_order","orderId":"123","table":"5","items":[{"name":"Pasta","quantity":2}]}`
		err := conn.WriteMessage(websocket.TextMessage, []byte(order))
		if err != nil {
			break
		}

		time.Sleep(5 * time.Second)
	}
}

func RegisterHotelModuleRoutes(v1 *gin.RouterGroup, db *gorm.DB) {
	// Dashboard
	v1.GET("/sharp-hospitality/getdashboard", func(c *gin.Context) {})

	// Hotel / Restaurant Configuration
	v1.POST("/sharp-hospitality/configure-hotel", func(c *gin.Context) {})

	// Tables / Rooms
	v1.GET("/sharp-hospitality/get-tables", func(c *gin.Context) {})
	v1.POST("/sharp-hospitality/add-table", func(c *gin.Context) {})
	v1.PUT("/sharp-hospitality/update-table/:id", func(c *gin.Context) {})
	v1.DELETE("/sharp-hospitality/delete-table/:id", func(c *gin.Context) {})

	// Menu
	v1.GET("/sharp-hospitality/get-menu", func(c *gin.Context) {})
	v1.POST("/sharp-hospitality/add-menu", func(c *gin.Context) {})
	v1.PUT("/sharp-hospitality/update-menu/:id", func(c *gin.Context) {})
	v1.DELETE("/sharp-hospitality/delete-menu/:id", func(c *gin.Context) {})

	// Orders
	v1.POST("/sharp-hospitality/create-order", func(c *gin.Context) {})
	v1.GET("/sharp-hospitality/get-orders", func(c *gin.Context) {})
	v1.PUT("/sharp-hospitality/update-order/:id", func(c *gin.Context) {})

	// Bookings / Reservations
	v1.POST("/sharp-hospitality/create-booking", func(c *gin.Context) {})
	v1.GET("/sharp-hospitality/get-booking", func(c *gin.Context) {})
	v1.PUT("/sharp-hospitality/update-booking/:id", func(c *gin.Context) {})
	v1.DELETE("/sharp-hospitality/cancel-booking/:id", func(c *gin.Context) {})

	// Staff Management (optional MVP)
	v1.POST("/sharp-hospitality/add-staff", func(c *gin.Context) {})
	v1.GET("/sharp-hospitality/get-staff", func(c *gin.Context) {})
	v1.PUT("/sharp-hospitality/update-staff/:id", func(c *gin.Context) {})
	v1.GET("/sharp-hospitality/orders-ws", OrdersWebSocketHandler)
}
