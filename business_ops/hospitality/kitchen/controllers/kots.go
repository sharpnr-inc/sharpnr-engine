package kitchencontrollers

import (
	"fmt"
	"io"
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *KitchenController) GetKots(ctx *gin.Context) {
	// 1. Resolve Auth and KitchenID
	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	// 2. Subscribe using the EXACT same string format as the Service
	// Since kitchenID is a UUID, we use .String() and %s
	topic := fmt.Sprintf("kitchen_%s", kitchenID.String())
	clientChan := c.sseEngine.Subscribe(topic)

	// 3. Clean up when the client leaves
	defer c.sseEngine.Unsubscribe(topic, clientChan)

	// 4. Set Headers for SSE
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")

	// 5. Immediate First Load
	// This ensures the user sees their KOTs instantly without waiting for an update
	initialData, err := c.kitchenService.GetKots(ctx.Request.Context(), kitchenID)
	if err == nil {
		ctx.SSEvent("message", initialData)
		ctx.Writer.Flush()
	}

	// 6. Real-time Stream
	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-ctx.Request.Context().Done():
			// Browser closed or connection lost
			return false

		case data := <-clientChan:
			// Triggered by s.sseEngine.MessageChan in the Service
			ctx.SSEvent("message", data)
			ctx.Writer.Flush()
			return true
		}
	})
}

func (c *KitchenController) UpdateKotsStatus(ctx *gin.Context) {

	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	kotID, ok := c.resolveKOTsID(ctx)
	if !ok {
		return
	}

	// 2. Bind the new status from the request body
	var input struct {
		Status string `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "", err)
		return
	}

	if err := c.kitchenService.UpdateKotStatus(ctx.Request.Context(), kitchenID, kotID, input.Status); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "KOT updated successfully")

}
