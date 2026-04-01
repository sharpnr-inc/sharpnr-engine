package orchestrator_controllers

import (
	"net/http"
	inventory_dtos "sharpnr-engine/service_engines/inventory_engine/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) NewInventoryAccountHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	request, ok := utils.BindJSON[inventory_dtos.NewInventoryAccountRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddNewInventoryAccount(ctx.Request.Context(), orgID, emplID, request); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "")
}

func (c *OrchestratorController) ListAllInventoryAccountHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	// Call the orchestrator to get all accounts for the organization
	accounts, err := c.orchestService.ListAllInventoryAccount(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to retrieve inventory accounts", err)
		return
	}

	// Return the list with a success message and the data payload
	utils.Success(ctx, http.StatusOK, "Inventory accounts retrieved successfully", accounts)
}
