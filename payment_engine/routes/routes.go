package payment_engine_routes

import (
	"sharpnr-engine/app"
	payment_engine_controllers "sharpnr-engine/payment_engine/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentGatewaysRoutes(v1 *gin.RouterGroup, modules app.ModulesProvider) {

	paymentService := modules.PaymentEngineService()
	pc := payment_engine_controllers.NewPaymentEngineController(paymentService)

	v1.POST("/webhooks/phonepe/callback", pc.HandlePhonePeCallback)
}
