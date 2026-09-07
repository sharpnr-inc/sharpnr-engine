package hrmsEngineControllers

import (
	"net/http"

	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// HrmsAccountStatusHandler godoc
// @Summary      HRMS service status
// @Description  Reports whether the caller's organization has started the HRMS (SharpHR) service and whether that account is switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — the client asks this before it knows. Starting the service is an orchestrator call, not this one.
// @Tags         SharpHR
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=hrmsEngineDtos.HrmsAccountStatusResponse} "HRMS status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /sharp-hr/status [get]
func (c *hrmsControllers) HrmsAccountStatusHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.hrmsService.GetHrmsAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromHrmsError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "HRMS status fetched successfully", status)
}
