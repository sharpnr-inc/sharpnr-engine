package orchestrator_controllers

import (
	"errors"
	"net/http"

	"sharpnr-engine/shared/utils"
	siloEngineDtos "sharpnr-engine/silo_engine/dtos"
	siloEngineErrs "sharpnr-engine/silo_engine/errs"

	"github.com/gin-gonic/gin"
)

// failFromSiloError maps the Silo module's sentinel errors onto status codes
// before falling back to the shared mapper.
//
// "Not initialized" and "disabled" are both 409 rather than 404: the org is real
// and the caller is allowed to ask, but the service is in a state that refuses
// the request — and the fix differs (start it vs. it was switched off), which is
// why the two errors stay apart all the way to the client.
func failFromSiloError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, siloEngineErrs.ErrSiloServiceNotInitialized),
		errors.Is(err, siloEngineErrs.ErrSiloAccountIsDisabled),
		errors.Is(err, siloEngineErrs.ErrSiloServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, siloEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetSiloAccountStatusHandler godoc
// @Summary      Silo service status
// @Description  Reports whether the organization has started the Silo (storage and data) service and whether that account is currently switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.Silo
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=siloEngineDtos.SiloAccountStatusResponse} "Silo status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/silo [get]
func (c *OrchestratorController) GetSiloAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetSiloAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromSiloError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Silo status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartSiloServiceHandler godoc
// @Summary      Start the Silo service
// @Description  Provisions the organization's one Silo account together with its settings row. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.Silo
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                    true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      siloEngineDtos.StartSiloServiceRequestDto true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "Silo service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Silo service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/silo/start [post]
func (c *OrchestratorController) StartSiloServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[siloEngineDtos.StartSiloServiceRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartSiloService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromSiloError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Silo service started successfully")
}
