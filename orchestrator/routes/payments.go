package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func paymentOrchestratorRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	payRoutes := rg.Group("payments")

	{
		payRoutes.GET("", c.GetPaymentAccount)
		payRoutes.GET("/details/:paymentAccountID", c.GetPaymentAccountGatewayDetails)
		payRoutes.POST("/add/config", c.AddGatwayConfigInPayment)
		payRoutes.POST("/create", c.CreatePaymentAccount)
	}

}
