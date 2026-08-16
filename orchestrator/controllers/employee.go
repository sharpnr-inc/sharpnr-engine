package orchestrator_controllers

import (
	"net/http"
	coreEmployeeContracts "sharpnr-engine/core/employees/contracts"
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

func (c *OrchestratorController) CreateSuperAdmin(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[coreEmployeeContracts.AddNewSuperAdminInOrganizationRequestDto](ctx)

	if err := c.orchestService.AddNewSuperAdmin(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Invitation for super admin sent to email")

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
