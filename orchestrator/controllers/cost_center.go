package orchestrator_controllers

import (
	"errors"
	"net/http"

	orgCostCenterDtos "sharpnr-engine/core/cost_center/dtos"
	orgCostCenterErrs "sharpnr-engine/core/cost_center/errs"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// failFromCostCenterError maps the cost center module's sentinel errors onto
// status codes before falling back to the shared mapper.
//
// InvalidCostCenterErr — the row exists but belongs to another organization —
// answers 404 alongside a genuine miss on purpose: a 403 here would confirm
// that another org's id exists.
func failFromCostCenterError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, orgCostCenterErrs.CostCenterNotFoundErr),
		errors.Is(err, orgCostCenterErrs.InvalidCostCenterErr),
		errors.Is(err, orgCostCenterErrs.NoCostCenterFound):
		utils.Failure(ctx, http.StatusNotFound, err)
	case errors.Is(err, orgCostCenterErrs.CostCenterCodeAlreadyExistsErr),
		errors.Is(err, orgCostCenterErrs.DepartmentCodeAlreadyExistsErr):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, orgCostCenterErrs.CostCenterNotActiveErr):
		// The site exists and is readable, but is INACTIVE or ARCHIVED: the
		// request conflicts with its state rather than being malformed.
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, orgCostCenterErrs.CostCenterNameRequiredErr),
		errors.Is(err, orgCostCenterErrs.CostCenterCodeRequiredErr),
		errors.Is(err, orgCostCenterErrs.DepartmentNameRequiredErr),
		errors.Is(err, orgCostCenterErrs.DepartmentCodeRequiredErr),
		errors.Is(err, orgCostCenterErrs.InvalidDepartmentHeadErr),
		errors.Is(err, orgCostCenterErrs.InvalidCostCenterTypeErr),
		errors.Is(err, orgCostCenterErrs.InvalidCostCenterStatusErr),
		errors.Is(err, orgCostCenterErrs.InvalidTimezoneErr):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// CreateCostCenterHandler godoc
// @Summary      Create a cost center
// @Description  Creates a new operating location for the organization. The code is derived from the name when the payload does not carry one; a new site starts ACTIVE and unlinked from any legal entity.
// @Tags         Orchestrator.CostCenters
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                        true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      orgCostCenterDtos.CreateCostCenterRequestDto  true  "Cost center to create"
// @Success      201      {object}  utils.Response "Cost center created successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Cost center code already exists"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cost_center [post]
func (c *OrchestratorController) CreateCostCenterHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[orgCostCenterDtos.CreateCostCenterRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.CreateNewCostCenter(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromCostCenterError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Cost center created successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// GetCostCenterHandler godoc
// @Summary      Get one cost center
// @Description  Returns a single cost center of the current organization.
// @Tags         Orchestrator.CostCenters
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string  true  "Session cookie containing org_id and empl_id"
// @Param        costCenterID  path      string  true  "Cost center ID"
// @Success      200           {object}  utils.Response{data=orgCostCenterDtos.CostCenterResponseDto} "Cost center fetched successfully"
// @Failure      400           {object}  utils.Response "Invalid cost center id"
// @Failure      401           {object}  utils.Response "Unauthorized"
// @Failure      404           {object}  utils.Response "Cost center not found"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cost_center/{costCenterID} [get]
func (c *OrchestratorController) GetCostCenterHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	costCenterID, ok := c.getCostCenterIDFromParams(ctx)
	if !ok {
		return
	}

	costCenter, err := c.orchestService.GetOrgCostCenterDetails(ctx.Request.Context(), orgID, emplID, costCenterID)
	if err != nil {
		failFromCostCenterError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Cost center fetched successfully", costCenter)
}

/* ---------------------------------- <***> --------------------------------- */

// ListAllCostCentersHandler godoc
// @Summary      List the organization's cost centers
// @Description  Returns every cost center of the current organization. An organization on the department-only track has none, which is an empty list rather than an error.
// @Tags         Orchestrator.CostCenters
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=[]orgCostCenterDtos.CostCenterResponseDto} "Cost centers fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cost_center [get]
func (c *OrchestratorController) ListAllCostCentersHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	costCenters, err := c.orchestService.ListAllOrgCostCenters(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCostCenterError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Cost centers fetched successfully", costCenters)
}

/* ---------------------------------- <***> --------------------------------- */

// CreateCostCenterDepartmentHandler godoc
// @Summary      Create a department inside a cost center
// @Description  Adds a department to one of the organization's cost centers — the Bangalore office's Finance team. These are the cost center's own departments, not the org-level ones under /orchestrator/department; the two tracks never mix. The code is derived from the name when the payload does not carry one, and an INACTIVE or ARCHIVED cost center takes no new departments.
// @Tags         Orchestrator.CostCenters
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string                                                  true  "Session cookie containing org_id and empl_id"
// @Param        costCenterID  path      string                                                  true  "Cost center ID"
// @Param        payload       body      orgCostCenterDtos.CreateCostCenterDepartmentRequestDto  true  "Department to create"
// @Success      201           {object}  utils.Response "Department created successfully"
// @Failure      400           {object}  utils.Response "Validation error"
// @Failure      401           {object}  utils.Response "Unauthorized"
// @Failure      404           {object}  utils.Response "Cost center not found"
// @Failure      409           {object}  utils.Response "Department code already exists in this cost center"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cost_center/{costCenterID}/department [post]
func (c *OrchestratorController) CreateCostCenterDepartmentHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	costCenterID, ok := c.getCostCenterIDFromParams(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[orgCostCenterDtos.CreateCostCenterDepartmentRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.CreateNewDepartmentInCostCenter(ctx.Request.Context(), orgID, emplID, costCenterID, *req); err != nil {
		failFromCostCenterError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Department created successfully")
}
