package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddBusinessOperationModulesInPropertyHandler(ctx *gin.Context) {
	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	propertyID, ok := pathUUID(ctx, "propertyID")
	if !ok {
		return
	}

	req, ok := utils.BindJSON[orchestrator_dtos.AddBusinessOpModulesRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddBusinessOperationModulesInProperty(ctx.Request.Context(), orgID, emplID, propertyID, *req); err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Modules Added Successfully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetBusinessOperationModulesOfPropertyHandler(ctx *gin.Context) {

	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	propertyID, ok := pathUUID(ctx, "propertyID")
	if !ok {
		return
	}

	response, err := c.orchestService.ListAllBusinessOpsModules(ctx.Request.Context(), orgID, emplID, propertyID)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Business operation modules fetched successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */
