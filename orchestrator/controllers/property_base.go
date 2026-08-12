package orchestrator_controllers

import (
	"net/http"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	healthcare_utilities "sharpnr-engine/property_business_ops/healthcare/utility"
	hospitality_utilities "sharpnr-engine/property_business_ops/hospitality/utilities"
	general_infra_utils "sharpnr-engine/property_infra/general_infra/utils"
	healthcare_infra_utils "sharpnr-engine/property_infra/healthcare_infra/utils"
	hospitality_infra_utils "sharpnr-engine/property_infra/hospitality_infra/utils"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// The /property route tree's handlers are split across property_*.go files
// mirroring orchestrator/service/property: this file is the property itself,
// property_business_modules.go and property_infra_modules.go are the module
// composition endpoints, and the per-entity handlers (hotel, restaurant, bar,
// kitchen, menu) land in property_hosp_<entity>.go as they are implemented.
// They all stay methods on the one OrchestratorController, so the route tree in
// routes/property_routes.go keeps its single controller argument.

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
	pId, ok := pathUUID(ctx, "propertyID")
	if !ok {
		return // pathUUID already sent the 400 Failure response
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
		utils.Failure(ctx, http.StatusBadRequest, err)
		return
	}

	if err := c.orchestService.CreateProperty(ctx.Request.Context(), orgID, emplID, req); err != nil {
		failFromServiceError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Property Created Successfully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) InfraModulesHandler(ctx *gin.Context) {

	// Ignore the organization ID and employeeID
	// this is for Authentication purpose, so no loose enpoint is exposed
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	data := gin.H{
		"general":     general_infra_utils.GeneralInfraModules,
		"hospitality": hospitality_infra_utils.HospitalityInfraModules,
		"healthcare":  healthcare_infra_utils.HealthcareInfraModules,
	}

	utils.Success(ctx, http.StatusOK, "Infra modules fetched successfully", data)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) BusinessOperationModulesHandler(ctx *gin.Context) {

	// Ignore the organization ID and employeeID
	// this is for Authentication purpose, so no loose enpoint is exposed
	_, _, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	data := gin.H{
		"hospitality": hospitality_utilities.HospitalityModules,
		"healthcare":  healthcare_utilities.HealthcareModules,
	}

	utils.Success(ctx, http.StatusOK, "Business modules fetched successfully", data)
}
