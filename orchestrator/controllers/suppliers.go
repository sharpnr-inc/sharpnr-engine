package orchestrator_controllers

import (
	"net/http"
	suppliers_dtos "sharpnr-engine/core/org_supplier/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) CreateNewSupplierHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	request, ok := utils.BindJSON[suppliers_dtos.CreateSuppliersAPIRequestdto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddNewSupplier(ctx.Request.Context(), orgID, emplID, request); err != nil {
		// Descriptive error message for internal tracking/debugging
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to create new supplier", err)
		return
	}

	// Success message confirms the action was completed
	utils.Success(ctx, http.StatusCreated, "Supplier created successfully", nil)
}

func (c *OrchestratorController) ListAllSupplierHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	suppliers, err := c.orchestService.ListAllOrganizationSuppliers(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to retrieve organization suppliers", err)
		return
	}

	// Returning the slice with a clear confirmation message
	utils.Success(ctx, http.StatusOK, "Suppliers retrieved successfully", suppliers)
}
