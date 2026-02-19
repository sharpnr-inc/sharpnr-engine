package kitchencontrollers

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *kitchenController) GetKots(ctx *gin.Context) {
	// 1. Resolve Auth and KitchenID
	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	// 2. Set SSE Headers
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")

	// 3. Start the stream
	ctx.Stream(func(w io.Writer) bool {
		// Fetch fresh data from service
		res, err := c.kitchenService.GetKots(ctx.Request.Context(), kitchenID)
		if err != nil {
			// In SSE, we send errors as a specific event or log them
			ctx.SSEvent("error", gin.H{"message": "failed to fetch KOTs"})
			return false // Stop the stream on critical error
		}

		// Send the data to the client
		ctx.SSEvent("message", res)

		// Flush the data immediately to the client
		ctx.Writer.Flush()

		// Wait for a few seconds before the next update (polling interval)
		// Adjust this based on how 'live' you want the kitchen to be
		select {
		case <-ctx.Request.Context().Done():
			return false // Client disconnected, stop the loop
		case <-time.After(5 * time.Second):
			return true // Continue to the next loop iteration
		}
	})
}
