package orchestrator_controllers

import (
	"errors"
	"net/http"

	hrmsEngineDtos "sharpnr-engine/hrms_engine/dtos"
	hrmsEngineErrs "sharpnr-engine/hrms_engine/errs"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// failFromHrmsError maps the HRMS module's sentinel errors onto status codes
// before falling back to the shared mapper.
//
// "Not found" and "disabled" are both 409 rather than 404: the org is real and
// the caller is allowed to ask, but the service is in a state that refuses the
// request — and the fix differs (start it vs. it was switched off), which is why
// the two errors stay apart all the way to the client.
func failFromHrmsError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, hrmsEngineErrs.ErrNoHRMSAccountFound),
		errors.Is(err, hrmsEngineErrs.ErrAccountDisabled),
		errors.Is(err, hrmsEngineErrs.ErrHrmsServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, hrmsEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetHrmsAccountStatusHandler godoc
// @Summary      HRMS service status
// @Description  Reports whether the organization has started the HRMS (SharpHR) service and whether that account is currently switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.SharpHR
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=hrmsEngineDtos.HrmsAccountStatusResponse} "HRMS status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/sharp-hr [get]
func (c *OrchestratorController) GetHrmsAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetHrmsAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromHrmsError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "HRMS status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartHrmsServiceHandler godoc
// @Summary      Start the HRMS service
// @Description  Provisions the organization's one HRMS account together with its settings row. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.SharpHR
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                    true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      hrmsEngineDtos.StartHrmsServiceRequestDto true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "HRMS service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "HRMS service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/sharp-hr/start [post]
func (c *OrchestratorController) StartHrmsServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[hrmsEngineDtos.StartHrmsServiceRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartHrmsService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromHrmsError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "HRMS service started successfully")
}
