package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetInfraModulesOfPropertyHandler(ctx *gin.Context) {

	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	propertyID, ok := pathUUID(ctx, "propertyID")
	if !ok {
		return
	}

	response, err := c.orchestService.ListAllInfraModulesOfProperty(ctx.Request.Context(), orgID, emplID, propertyID)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Infra modules fetched successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddInfraModulesOfPropertyHandler(ctx *gin.Context) {

	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	propertyID, ok := pathUUID(ctx, "propertyID")
	if !ok {
		return
	}

	req, ok := utils.BindJSON[orchestrator_dtos.AddInfraModulesRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddInfraModulesInProperty(ctx.Request.Context(), orgID, emplID, propertyID, *req); err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Infra modules added successfully")
}

/* ---------------------------------- <***> --------------------------------- */
