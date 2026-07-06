package kitchencontrollers

import (
	"io"
	"net/http"
	"time"

	kitchen_contracts "sharpnr-engine/business_ops/hospitality/kitchen/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// keepAliveInterval must stay below proxy/LB idle timeouts (nginx defaults
// to 60s) or quiet streams get cut mid-connection.
const keepAliveInterval = 20 * time.Second

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

	// 2. Subscribe before the initial fetch so an update landing in between
	// is caught by the subscription instead of lost.
	sub := c.sseEngine.Subscribe(kitchen_contracts.KotsTopic(kitchenID))
	defer sub.Close()

	// 3. Immediate first load so the user sees their KOTs instantly. Failing
	// here (before any stream bytes are written) still allows a proper HTTP
	// error; the client's EventSource can react and reconnect.
	initialData, err := c.kitchenService.GetKots(ctx.Request.Context(), kitchenID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	// 4. Set Headers for SSE
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	// Stops nginx from buffering the stream; without it events sit in the
	// proxy until its buffer fills.
	ctx.Header("X-Accel-Buffering", "no")

	ctx.SSEvent("message", initialData)
	ctx.Writer.Flush()

	keepAlive := time.NewTicker(keepAliveInterval)
	defer keepAlive.Stop()

	// 5. Real-time Stream (gin flushes after every step)
	ctx.Stream(func(w io.Writer) bool {
		select {
		case <-ctx.Request.Context().Done():
			// Browser closed or connection lost
			return false

		case ev, open := <-sub.C():
			if !open {
				// Engine shut down (server stopping)
				return false
			}
			ctx.SSEvent("message", ev.Payload)
			return true

		case <-keepAlive.C:
			// SSE comment: invisible to EventSource, keeps intermediaries
			// from closing an idle connection.
			io.WriteString(w, ": keepalive\n\n")
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
