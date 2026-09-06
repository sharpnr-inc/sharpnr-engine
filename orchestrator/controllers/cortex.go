package orchestrator_controllers

import (
	"errors"
	"net/http"

	cortexEngineDtos "sharpnr-engine/cortex_engine/dtos"
	cortexEngineErrs "sharpnr-engine/cortex_engine/errs"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// failFromCortexError maps the Cortex module's sentinel errors onto status
// codes before falling back to the shared mapper.
//
// "Not found" and "disabled" are both 409 rather than 404: the org is real and
// the caller is allowed to ask, but the service is in a state that refuses the
// request — and the fix differs (start it vs. it was switched off), which is why
// the two errors stay apart all the way to the client.
func failFromCortexError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, cortexEngineErrs.ErrNoCortexAccountFound),
		errors.Is(err, cortexEngineErrs.ErrCortexAccountIsDisabled),
		errors.Is(err, cortexEngineErrs.ErrCortexServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, cortexEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetCortexAccountStatusHandler godoc
// @Summary      Cortex service status
// @Description  Reports whether the organization has started the Cortex (IT controls and access management) service and whether that account is currently switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.Cortex
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=cortexEngineDtos.CortexAccountStatusResponse} "Cortex status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cortex [get]
func (c *OrchestratorController) GetCortexAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetCortexAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCortexError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Cortex status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartCortexServiceHandler godoc
// @Summary      Start the Cortex service
// @Description  Provisions the organization's one Cortex account together with its settings row. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.Cortex
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                        true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      cortexEngineDtos.StartCortexServiceRequestDto true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "Cortex service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Cortex service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cortex/start [post]
func (c *OrchestratorController) StartCortexServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cortexEngineDtos.StartCortexServiceRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartCortexService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromCortexError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Cortex service started successfully")
}
