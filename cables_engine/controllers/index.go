package cablesEngineControllers

import (
	"net/http"

	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// GetCablesAccountStatus godoc
// @Summary      Cables service status
// @Description  Reports whether the caller's organization has started the cables service and whether that account is switched on. An organization that never started it answers isStarted=false rather than a 404 — the client asks this before it knows.
// @Tags         Cables
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=cablesEngineDtos.CablesServiceResponseDto} "Cables status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /cables/status [get]
func (c *cablesEngineController) GetCablesAccountStatus(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.cablesService.GetCablesServiceDetails(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Cables status fetched successfully", status)
}
