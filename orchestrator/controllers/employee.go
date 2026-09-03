package orchestrator_controllers

import (
	"errors"
	"net/http"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
	coreEmployeeErrs "sharpnr-engine/core/employees/errs"
	orchestrator_dtos "sharpnr-engine/orchestrator/dtos"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) FetchAllEmployeesHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	employees, err := c.orchestService.GetAllEmployeesOfOrg(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Employees fetched successfully", employees)

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetEmployeeDetails(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	requestedDetails, ok := c.getEmployeeIDFromParams(ctx)
	if !ok {
		return
	}

	employeeDetails, err := c.orchestService.GetEmployeeDetails(ctx.Request.Context(), orgID, emplID, requestedDetails)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Employee Details Fetched Successfully", employeeDetails)
}

/* ---------------------------------- <***> --------------------------------- */

// failFromEmployeeError maps the employee module's sentinel errors onto status
// codes before falling back to the shared mapper, which handles the Sentinel
// deny and the generic not-found cases.
func failFromEmployeeError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, coreEmployeeErrs.EmployeeNotFoundErr):
		utils.Failure(ctx, http.StatusNotFound, err)
	case errors.Is(err, coreEmployeeErrs.EmployeeAlreadyExistsErr):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, coreEmployeeErrs.EmployeeFirstNameRequiredErr),
		errors.Is(err, coreEmployeeErrs.EmployeeEmailRequiredErr),
		errors.Is(err, coreEmployeeErrs.EmployeePasswordRequiredErr):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// CreateSuperAdmin godoc
// @Summary      Add an owner / super admin
// @Description  Creates a second owner (super admin) in the caller's organization and puts them in the seeded Owner role. Only an existing owner may call it.
// @Description  Temporary behaviour: the account is created active from the password in the payload; the invitation flow is not built yet.
// @Tags         Orchestrator.Employees
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                                          true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto  true  "Super admin to create"
// @Success      201      {object}  utils.Response "Super admin created successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      403      {object}  utils.Response "Access denied"
// @Failure      409      {object}  utils.Response "An employee with this email already exists"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/empl/add_super_admin [post]
func (c *OrchestratorController) CreateSuperAdmin(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddNewSuperAdmin(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromEmployeeError(ctx, err)
		return
	}

	// TEMP: the account already exists and can sign in with the password from
	// the payload — no mail is sent. Reword this to the invitation copy when the
	// provisioning-link flow replaces it.
	utils.Success(ctx, http.StatusCreated, "Super admin created successfully")

}

/* ---------------------------------- <***> --------------------------------- */

// this controller will add employees in the bulk in the organization.
// this method will can take the xls format sheet
func (c *OrchestratorController) AddEmployeeInBulk(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}
	req, ok := utils.BindJSON[[]orchestrator_dtos.AddEmployeeBulkRequestDto](ctx)

	result, err := c.orchestService.AddEmployeeInBulk(ctx.Request.Context(), orgID, emplID, req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}
	utils.Success(ctx, http.StatusCreated, "Invitation for super admin sent to email", result)

}

/* ---------------------------------- <***> --------------------------------- */
