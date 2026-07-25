package orchestrator_controllers

import (
	"net/http"
	supplier_contracts "sharpnr-engine/core/org_supplier/contracts"
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

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddLegalEntityLinkToSupplierHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	request, ok := utils.BindJSON[*supplier_contracts.AddLegalEntityLinkInSuppliersRequestDto](ctx)
	if !ok {
		return
	}

	suppliersID, ok := pathUUID(ctx, "suppliersID")
	if !ok {
		return
	}

	if err := c.orchestService.AddLegalEntityLinksInSuppliers(ctx.Request.Context(), orgID, emplID, suppliersID, *request); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to link legal entity to supplier", err)
		return
	}

	// Using StatusCreated (201) and a clear success message
	utils.Success(ctx, http.StatusCreated, "Legal entity linked to supplier successfully", nil)
}

/* ---------------------------------- <***> --------------------------------- */
