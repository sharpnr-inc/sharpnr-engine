package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func paymentOrchestratorRoutes(rg *gin.RouterGroup, c *orchestrator_controllers.OrchestratorController) {

	payRoutes := rg.Group("payments")
	payRoutes.GET("", c.GetPaymentAccount)
	payRoutes.POST("/create", c.CreatePaymentAccount)

}
