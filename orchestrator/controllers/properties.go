package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) ListAllOrgProperty(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetAllPropertiesOfOrg(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Properties fetch successfully", res)
}

/* ----------------------------------- <> ----------------------------------- */

func (c *OrchestratorController) GetPropertyDetails(ctx *gin.Context) {
	// 1. User Authentication Check
	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	// 2. Param Parsing using your helper
	pId, err := c.getPropertyFromParams(ctx)
	if err != nil {
		return // getPropertyFromParams already sent the 400 Failure response
	}

	// 4. Service Call
	res, err := c.orchestService.GetPropertyDetail(ctx.Request.Context(), orgID, pId)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	// 5. Success Response
	utils.Success(ctx, http.StatusOK, "Property details fetch successfully", res)
}

// ------------------------------------------------------------------------------//

func (c *OrchestratorController) AddModulesInProperty(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	pId, err := c.getPropertyFromParams(ctx)
	if err != nil {
		return
	}

	var req *orchestrator_dtos.PropertyModuleConfigurationDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	// 1. Capture both the result and the error
	result, err := c.orchestService.AddModulesInProperty(ctx.Request.Context(), orgID, pId, req)

	if err != nil {
		// Handle specific errors (like the "Only one module" constraint)
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	// 2. Pass the result object back in the success response
	// This gives the frontend the 'Created', 'Skipped', and 'Message' fields
	utils.Success(ctx, http.StatusOK, "Result", result)

}

/* --------------------------------- <****> --------------------------------- */

func (c *OrchestratorController) AddInfraInProperty(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	pId, err := c.getPropertyFromParams(ctx)
	if err != nil {
		return
	}

	var req *orchestrator_dtos.PropertyInfraConfiguration
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	if err := c.orchestService.AddInfraInProperty(ctx.Request.Context(), orgID, pId, req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeConflict, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Result", nil)

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddProperty(ctx *gin.Context) {

	// 1. User Authentication Check
	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	var req *orchestrator_dtos.CreatePropertyReq_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err := c.orchestService.CreateProperty(ctx.Request.Context(), orgID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Property Created Successfully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetPropertyInfra(ctx *gin.Context) {

	orgID, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	// 2. Param Parsing using your helper
	pId, err := c.getPropertyFromParams(ctx)
	if err != nil {
		return // getPropertyFromParams already sent the 400 Failure response
	}

	res, err := c.orchestService.GetPropertyInfra(ctx.Request.Context(), orgID, pId, nil)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Property Infra fetched successfully", res)
}
