package stockwaveEngineControllers

import (
	"net/http"

	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// StockwaveAccountStatusHandler godoc
// @Summary      Stockwave service status
// @Description  Reports whether the caller's organization has taken the stockwave (stock and warehouse) service and whether that account is switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — the client asks this before it knows. Starting the service is an orchestrator call, not this one.
// @Tags         Stockwave
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=stockwaveEngineDtos.StockwaveEngineStatusResponseDto} "Stockwave status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /stockwave/status [get]
func (c *stockwaveController) StockwaveAccountStatusHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.stockwaveService.GetStockwaveAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromStockwaveError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Stockwave status fetched successfully", status)
}
