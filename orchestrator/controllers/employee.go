package orchestrator_controllers

import (
	"net/http"
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
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Employees fetched successfully", employees)

}
