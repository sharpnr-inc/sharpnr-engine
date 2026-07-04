package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// ListAllOrgProperty godoc
// @Summary      List all properties of the organization
// @Description  Fetches all properties belonging to the current organization, scoped to the authenticated employee.
// @Tags         Orchestrator.Properties
// @Accept       json
// @Produce      json
// @Param        Cookie  header  string  true  "Session cookie containing org_id and empl_id"
// @Success      200  {object}  utils.Response  "Properties fetch successfully"
// @Failure      401  {object}  utils.Response  "Unauthorized"
// @Failure      500  {object}  utils.Response  "Internal Server Error"
// @Router       /orchestrator/property [get]
func (c *OrchestratorController) ListAllOrgProperty(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetAllPropertiesOfOrg(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Properties fetch successfully", res)
}

/* ----------------------------------- <> ----------------------------------- */

// GetPropertyDetails godoc
// @Summary      Get property details
// @Description  Fetches the full details of a single property belonging to the current organization.
// @Tags         Orchestrator.Properties
// @Accept       json
// @Produce      json
// @Param        Cookie      header  string  true  "Session cookie containing org_id"
// @Param        propertyID  path    string  true  "Property ID (UUID)"
// @Success      200  {object}  utils.Response  "Property details fetch successfully"
// @Failure      400  {object}  utils.Response  "Invalid property ID"
// @Failure      401  {object}  utils.Response  "Unauthorized"
// @Failure      500  {object}  utils.Response  "Internal Server Error"
// @Router       /orchestrator/property/{propertyID} [get]
func (c *OrchestratorController) GetPropertyDetails(ctx *gin.Context) {
	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	// 2. Param Parsing using your helper
	pId, err := c.getPropertyFromParams(ctx)
	if err != nil {
		return // getPropertyFromParams already sent the 400 Failure response
	}

	// 4. Service Call
	res, err := c.orchestService.GetPropertyDetail(ctx.Request.Context(), orgID, emplID, pId)
	if err != nil {
		failFromServiceError(ctx, err)
		return
	}

	// 5. Success Response
	utils.Success(ctx, http.StatusOK, "Property details fetch successfully", res)
}

// ------------------------------------------------------------------------------//

// AddProperty godoc
// @Summary      Create a property
// @Description  Creates a new property under the current organization.
// @Tags         Orchestrator.Properties
// @Accept       json
// @Produce      json
// @Param        Cookie   header  string                                    true  "Session cookie containing org_id"
// @Param        request  body    orchestrator_dtos.CreatePropertyReq_dto   true  "Property details"
// @Success      200  {object}  utils.Response  "Property Created Successfully"
// @Failure      400  {object}  utils.Response  "Validation Error"
// @Failure      401  {object}  utils.Response  "Unauthorized"
// @Failure      500  {object}  utils.Response  "Internal Server Error"
// @Router       /orchestrator/property/create [post]
func (c *OrchestratorController) AddProperty(ctx *gin.Context) {

	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	var req *orchestrator_dtos.CreatePropertyReq_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err := c.orchestService.CreateProperty(ctx.Request.Context(), orgID, emplID, req); err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Property Created Successfully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddBusinessOperationModulesInPropertyHandler(ctx *gin.Context) {
	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	proprtyID, err := c.getPropertyFromParams(ctx)
	if err != nil {
		utils.Failure(ctx, http.StatusBadGateway, "Important Data is missing", err)
	}

	req, ok := utils.BindJSON[orchestrator_dtos.AddBusinessOpModulesRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddBusinessOperationModulesInProperty(ctx.Request.Context(), orgID, emplID, proprtyID, *req); err != nil {
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

	propertyID, err := c.getPropertyFromParams(ctx)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Important data is missing", err)
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

func (c *OrchestratorController) GetInfraModulesOfPropertyHandler(ctx *gin.Context) {

	// 1. User Authentication Check
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return // Response already sent by helper
	}

	propertyID, err := c.getPropertyFromParams(ctx)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Important data is missing", err)
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

	propertyID, err := c.getPropertyFromParams(ctx)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Important data is missing", err)
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
