package orchestrator_controllers

import (
	"net/http"
	org_customer_contracts "sharpnr-engine/core/customer/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) CreateCustomerHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[org_customer_contracts.AddCustomerRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddCustomer(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Some Bad Happed", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Customer Added Successfully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) SearchCustomerHandler(ctx *gin.Context) {
	// 1. Resolve multi-tenant context
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		// resolveRequestContexts should handle its own utils.Failure response
		return
	}

	// 2. Bind and Validate Input
	// Using a struct allows you to use 'binding' tags for automatic validation
	var req struct {
		SearchQuery string `form:"searchQuery" binding:"required,min=3"`
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Search query must be at least 3 characters", err)
		return
	}

	// 3. Service Call
	results, err := c.orchestService.SearchCustomer(ctx.Request.Context(), orgID, emplID, req.SearchQuery)
	if err != nil {
		// Log the actual error internally, return a cleaner message to the user
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to retrieve search results", err)
		return
	}

	// 4. Success Response
	utils.Success(ctx, http.StatusOK, "Customers fetched successfully", results)
}
