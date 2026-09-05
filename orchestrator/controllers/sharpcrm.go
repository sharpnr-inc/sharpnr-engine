package orchestrator_controllers

import (
	"errors"
	"net/http"

	"sharpnr-engine/shared/utils"
	sharpCrmEngineDtos "sharpnr-engine/sharpCRM_engine/dtos"
	sharpCrmEngineErrs "sharpnr-engine/sharpCRM_engine/errs"

	"github.com/gin-gonic/gin"
)

// failFromCrmError maps the SharpCRM module's sentinel errors onto status codes
// before falling back to the shared mapper.
//
// "Not started" and "disabled" are both 409 rather than 404: the org is real and
// the caller is allowed to ask, but the service is in a state that refuses the
// request — and the fix differs (start it vs. it was switched off), which is why
// the two errors stay apart all the way to the client.
func failFromCrmError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, sharpCrmEngineErrs.NoCustomerFoundOfthisID):
		utils.Failure(ctx, http.StatusNotFound, err)
	case errors.Is(err, sharpCrmEngineErrs.ErrCrmServiceNotInitialized),
		errors.Is(err, sharpCrmEngineErrs.ErrCrmAccountIsDisabled),
		errors.Is(err, sharpCrmEngineErrs.ErrCrmServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, sharpCrmEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetCrmAccountStatusHandler godoc
// @Summary      SharpCRM service status
// @Description  Reports whether the organization has started the SharpCRM service and whether that account is currently switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.SharpCRM
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=sharpCrmEngineDtos.SharpCrmAccountStatusResponse} "SharpCRM status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/sharpcrm [get]
func (c *OrchestratorController) GetCrmAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetCrmAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCrmError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "SharpCRM status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartCrmServiceHandler godoc
// @Summary      Start the SharpCRM service
// @Description  Provisions the organization's one SharpCRM account together with its settings row. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.SharpCRM
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                          true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      sharpCrmEngineDtos.StartCRMServiceRequestDto    true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "SharpCRM service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "SharpCRM service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/sharpcrm/start [post]
func (c *OrchestratorController) StartCrmServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[sharpCrmEngineDtos.StartCRMServiceRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartCrmService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromCrmError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "SharpCRM service started successfully")
}
