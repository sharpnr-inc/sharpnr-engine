package kitchencontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *kitchenController) GetOrders(ctx *gin.Context) {
	orders, err := c.kitchenService.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}

func (c *kitchenController) UpdateOrder(ctx *gin.Context) {
	var req struct {
		OrderID   string `json:"order_id"`
		NewStatus string `json:"new_status"`
		KitchenID string `json:"kitchen_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := c.kitchenService.UpdateOrder(req.OrderID, req.NewStatus); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Broadcast order status update to all WebSocket clients of this kitchen
	c.kitchenService.BroadcastOrderUpdate(req.KitchenID, map[string]string{
		"order_id": req.OrderID,
		"status":   req.NewStatus,
	})

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}
