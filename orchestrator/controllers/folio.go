package orchestrator_controllers

import (
	"errors"
	"net/http"

	folioEngineDtos "sharpnr-engine/folio_engine/dtos"
	folioEngineErrs "sharpnr-engine/folio_engine/errs"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// failFromFolioError maps the Folio module's sentinel errors onto status codes
// before falling back to the shared mapper.
//
// "Not initialized" and "disabled" are both 409 rather than 404: the org is real
// and the caller is allowed to ask, but the service is in a state that refuses
// the request — and the fix differs (start it vs. it was switched off), which is
// why the two errors stay apart all the way to the client.
func failFromFolioError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, folioEngineErrs.ErrFolioServiceNotInitialized),
		errors.Is(err, folioEngineErrs.ErrFolioAccountIsDisabled),
		errors.Is(err, folioEngineErrs.ErrFolioServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, folioEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetFolioAccountStatusHandler godoc
// @Summary      Folio service status
// @Description  Reports whether the organization has started the Folio (docs, sheets, slides and whiteboards) service and whether that account is currently switched on. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.Folio
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=folioEngineDtos.FolioAccountStatusResponse} "Folio status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/folio [get]
func (c *OrchestratorController) GetFolioAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetFolioAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromFolioError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Folio status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartFolioServiceHandler godoc
// @Summary      Start the Folio service
// @Description  Provisions the organization's one Folio account together with its settings row. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.Folio
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                      true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      folioEngineDtos.StartFolioServiceRequestDto true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "Folio service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Folio service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/folio/start [post]
func (c *OrchestratorController) StartFolioServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[folioEngineDtos.StartFolioServiceRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartFolioService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromFolioError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Folio service started successfully")
}
