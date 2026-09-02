package orchestrator_controllers

import (
	"errors"
	"net/http"

	orgDepartmentsDtos "sharpnr-engine/core/departments/dtos"
	orgDepartmentErrs "sharpnr-engine/core/departments/errs"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// failFromDepartmentError maps the departments module's sentinel errors onto
// status codes before falling back to the shared mapper, which handles the
// Sentinel deny and the generic not-found cases.
func failFromDepartmentError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, orgDepartmentErrs.ErrDepartmentNotFound):
		utils.Failure(ctx, http.StatusNotFound, err)
	case errors.Is(err, orgDepartmentErrs.ErrDepartmentAlreadyExists):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, orgDepartmentErrs.ErrDepartmentHasEmployees):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, orgDepartmentErrs.ErrDepartmentNameRequired),
		errors.Is(err, orgDepartmentErrs.ErrInvalidDepartmentHead):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// ListOrgAllDepartmentHandler godoc
// @Summary      List the organization's departments
// @Description  Returns every org-level department of the current organization.
// @Tags         Orchestrator.Departments
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=[]orgDepartmentsDtos.DepartmentViewResponseDto} "Departments fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      403     {object}  utils.Response "Access denied"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/department [get]
func (c *OrchestratorController) ListOrgAllDepartmentHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	departments, err := c.orchestService.ListAllOrgDepartments(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromDepartmentError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Departments fetched successfully", departments)
}

/* ---------------------------------- <***> --------------------------------- */

// CreateOrgDepartmentHandler godoc
// @Summary      Create a department
// @Description  Creates a new org-level department. The code is derived from the name when the payload does not carry one.
// @Tags         Orchestrator.Departments
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                        true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      orgDepartmentsDtos.CreateDepartmentRequestDto true  "Department to create"
// @Success      201      {object}  utils.Response "Department created successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      403      {object}  utils.Response "Access denied"
// @Failure      409      {object}  utils.Response "Department already exists"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/department [post]
func (c *OrchestratorController) CreateOrgDepartmentHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[orgDepartmentsDtos.CreateDepartmentRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.CreateNewOrgDeparment(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromDepartmentError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Department created successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// UpdateOrgDepartmentHandler godoc
// @Summary      Update a department
// @Description  Partially updates a department — only the fields present in the payload are written. The id in the path wins over any id in the body.
// @Tags         Orchestrator.Departments
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string                                        true  "Session cookie containing org_id and empl_id"
// @Param        departmentID  path      string                                        true  "Department ID"
// @Param        payload       body      orgDepartmentsDtos.UpdateDepartmentRequestDto true  "Fields to change"
// @Success      200           {object}  utils.Response "Department updated successfully"
// @Failure      400           {object}  utils.Response "Validation error"
// @Failure      403           {object}  utils.Response "Access denied"
// @Failure      404           {object}  utils.Response "Department not found"
// @Failure      409           {object}  utils.Response "Department already exists"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/department/{departmentID} [put]
func (c *OrchestratorController) UpdateOrgDepartmentHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	departmentID, ok := c.getDepartmentIDFromParams(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[orgDepartmentsDtos.UpdateDepartmentRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.UpdateOrgDepartment(ctx.Request.Context(), orgID, emplID, departmentID, *req); err != nil {
		failFromDepartmentError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Department updated successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// DeleteOrgDepartmentHandler godoc
// @Summary      Delete a department
// @Description  Soft-deletes a department. A department that still has employees assigned to it is refused — move them out first.
// @Tags         Orchestrator.Departments
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string  true  "Session cookie containing org_id and empl_id"
// @Param        departmentID  path      string  true  "Department ID"
// @Success      200           {object}  utils.Response "Department deleted successfully"
// @Failure      400           {object}  utils.Response "Invalid department id"
// @Failure      403           {object}  utils.Response "Access denied"
// @Failure      404           {object}  utils.Response "Department not found"
// @Failure      409           {object}  utils.Response "Department still has employees"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/department/{departmentID} [delete]
func (c *OrchestratorController) DeleteOrgDepartmentHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	departmentID, ok := c.getDepartmentIDFromParams(ctx)
	if !ok {
		return
	}

	if err := c.orchestService.DeleteOrgDepartment(ctx.Request.Context(), orgID, emplID, departmentID); err != nil {
		failFromDepartmentError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Department deleted successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// GetOrgDepartmentHandler godoc
// @Summary      Get one department
// @Description  Returns a single org-level department of the current organization.
// @Tags         Orchestrator.Departments
// @Accept       json
// @Produce      json
// @Param        Cookie        header    string  true  "Session cookie containing org_id and empl_id"
// @Param        departmentID  path      string  true  "Department ID"
// @Success      200           {object}  utils.Response{data=orgDepartmentsDtos.DepartmentViewResponseDto} "Department fetched successfully"
// @Failure      400           {object}  utils.Response "Invalid department id"
// @Failure      403           {object}  utils.Response "Access denied"
// @Failure      404           {object}  utils.Response "Department not found"
// @Failure      500           {object}  utils.Response "Internal server error"
// @Router       /orchestrator/department/{departmentID} [get]
func (c *OrchestratorController) GetOrgDepartmentHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	departmentID, ok := c.getDepartmentIDFromParams(ctx)
	if !ok {
		return
	}

	department, err := c.orchestService.GetOrgDepartmentDetails(ctx.Request.Context(), orgID, emplID, departmentID)
	if err != nil {
		failFromDepartmentError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Department fetched successfully", department)
}
