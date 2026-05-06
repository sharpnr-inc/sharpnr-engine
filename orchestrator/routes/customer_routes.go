package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

// These are Customer related routes for the Orchestrator App to find and manage customers in the system. These routes will be used by the frontend to perform customer-related operations such as
// creating a new customer or searching for existing customers. The handlers for these routes will be defined in the OrchestratorController, which will contain the business logic for handling customer
// data and interactions with the database or other services.

func registerCustomerRoutes(rg *gin.RouterGroup, orc *orchestrator_controllers.OrchestratorController) {

	orgCust := rg.Group("/org/customers")
	{
		orgCust.POST("/add", orc.CreateCustomerHandler)
		orgCust.GET("/search", orc.SearchCustomerHandler)
	}
}
