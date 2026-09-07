package orchestrator_controllers

import (
	"errors"
	"net/http"

	"sharpnr-engine/shared/utils"
	stockwaveEngineDtos "sharpnr-engine/stockwave_engine/dtos"
	stockwaveEngineErrs "sharpnr-engine/stockwave_engine/errs"

	"github.com/gin-gonic/gin"
)

// failFromStockwaveError maps the stockwave module's sentinel errors onto status
// codes before falling back to the shared mapper.
//
// "Not found" and "disabled" are both 409 rather than 404: the org is real and
// the caller is allowed to ask, but the service is in a state that refuses the
// request — and the fix differs (start it vs. it was switched off), which is why
// the two errors stay apart all the way to the client.
func failFromStockwaveError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, stockwaveEngineErrs.ErrNoAccountFound),
		errors.Is(err, stockwaveEngineErrs.ErrAccountDisabled),
		errors.Is(err, stockwaveEngineErrs.ErrStockwaveServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, stockwaveEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetStockwaveAccountStatusHandler godoc
// @Summary      Stockwave service status
// @Description  Reports whether the organization has taken the stockwave (stock and warehouse) service and whether that account is currently switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.Stockwave
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=stockwaveEngineDtos.StockwaveEngineStatusResponseDto} "Stockwave status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/stockwave [get]
func (c *OrchestratorController) GetStockwaveAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetStockwaveAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromStockwaveError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Stockwave status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartStockwaveServiceHandler godoc
// @Summary      Start the stockwave service
// @Description  Provisions the organization's one stockwave account together with its settings row. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.Stockwave
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                            true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      stockwaveEngineDtos.StartStockwaveEngineRequestDto true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "Stockwave service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Stockwave service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/stockwave/start [post]
func (c *OrchestratorController) StartStockwaveServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[stockwaveEngineDtos.StartStockwaveEngineRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartStockwaveService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromStockwaveError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Stockwave service started successfully")
}
