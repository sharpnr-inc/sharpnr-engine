package payment_engine_controllers

import (
	"fmt"
	"io"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"

	"github.com/gin-gonic/gin"
)

func (c *paymentEngineController) HandlePhonePeCallback(ctx *gin.Context) {
	// 1. Read raw body — must not be modified before signature verification
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "failed to read request body"})
		return
	}

	// 2. Extract headers PhonePe sends for signature verification
	headers := map[string]string{
		"x-verify": ctx.GetHeader("X-VERIFY"),
	}

	// 3. Confirm payment — verifies signature, updates payment status
	payment, err := c.paymentEngineService.ConfirmPayment(ctx.Request.Context(),
		payment_engine_contracts.ConfirmPaymentRequest{
			GatewayOrderRef: ctx.Query("order_id"), // PhonePe sends this as query param
			RawCallbackBody: string(body),
			Headers:         headers,
		},
	)
	if err != nil {
		// Log the error but still return 200
		// PhonePe retries on non-200 — we don't want infinite retries for bad data
		fmt.Printf("ERROR: PhonePe callback failed: %v\n", err)
		ctx.JSON(200, gin.H{"status": "ok"}) // ← always 200 to PhonePe
		return
	}

	c.paymentEngineService.PublishPaymentConfirmed(ctx.Request.Context(), payment)

	ctx.JSON(200, gin.H{"status": "ok"})
}
