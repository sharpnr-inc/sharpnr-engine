package hrmsEngineControllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *hrmsControllers) ListAllEmployeesOfOrg(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.hrmsService.ListOrgAllEmployees(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Employees Fetched Success fully", res)
}
