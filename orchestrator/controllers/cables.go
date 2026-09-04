package orchestrator_controllers

import (
	"errors"
	"net/http"

	cablesEngineDtos "sharpnr-engine/cables_engine/dtos"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

// failFromCablesError maps the cables module's sentinel errors onto status
// codes before falling back to the shared mapper.
//
// "Not initialized" and "disabled" are both 409 rather than 404: the org is
// real and the caller is allowed to ask, but the service is in a state that
// refuses the request — and the fix differs (start it vs. it was switched off),
// which is why the two errors stay apart all the way to the client.
func failFromCablesError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, cablesEngineErrs.ErrNoChannelFound):
		utils.Failure(ctx, http.StatusNotFound, err)
	case errors.Is(err, cablesEngineErrs.ErrCablesServiceNotInitialized),
		errors.Is(err, cablesEngineErrs.ErrCablesAccountIsDisabled),
		errors.Is(err, cablesEngineErrs.ErrCablesServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, cablesEngineErrs.ErrChannelAlreadyExists):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, cablesEngineErrs.ErrTermsNotAccepted),
		errors.Is(err, cablesEngineErrs.ErrChannelNameRequired),
		errors.Is(err, cablesEngineErrs.ErrInvalidChannelType),
		errors.Is(err, cablesEngineErrs.ErrInvalidChannelVisibility):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetCablesAccountStatusHandler godoc
// @Summary      Cables service status
// @Description  Reports whether the organization has started the cables (chat, meetings, calls) service and whether that account is currently switched on. An organization that never started it answers isStarted=false rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.Cables
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=cablesEngineDtos.CablesServiceResponseDto} "Cables status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cables [get]
func (c *OrchestratorController) GetCablesAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetCablesAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Cables status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartCablesServiceHandler godoc
// @Summary      Start the cables service
// @Description  Provisions the organization's one cables account. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account.
// @Tags         Orchestrator.Cables
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                             true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      cablesEngineDtos.StartCablesServiceRequestsDto     true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "Cables service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Cables service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cables/start [post]
func (c *OrchestratorController) StartCablesServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cablesEngineDtos.StartCablesServiceRequestsDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartCablesService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Cables service started successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// CreateCablesChannelHandler godoc
// @Summary      Create a channel
// @Description  Creates a channel in the organization's cables account. The slug is derived from the name — it is not accepted from the caller — and must not collide with a live channel in the same account. Type defaults to TEXT and visibility to PUBLIC when omitted.
// @Tags         Orchestrator.Cables
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                    true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      cablesEngineDtos.CreateChannelRequestDto  true  "Channel to create"
// @Success      201      {object}  utils.Response "Channel created successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Channel already exists, or the cables service is not started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cables/channel [post]
func (c *OrchestratorController) CreateCablesChannelHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[cablesEngineDtos.CreateChannelRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.CreateCablesChannel(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Channel created successfully")
}

/* ---------------------------------- <***> --------------------------------- */

// ListCablesChannelsHandler godoc
// @Summary      List the organization's channels
// @Description  Returns every channel in the organization's cables account, in sidebar order. This is the admin-side list: it is not filtered by membership, because participants do not exist yet.
// @Tags         Orchestrator.Cables
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=[]cablesEngineDtos.ChannelResponseDto} "Channels fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      409     {object}  utils.Response "Cables service is not started for this organization"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cables/channel [get]
func (c *OrchestratorController) ListCablesChannelsHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	channels, err := c.orchestService.ListCablesChannels(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Channels fetched successfully", channels)
}

/* ---------------------------------- <***> --------------------------------- */

// GetCablesChannelHandler godoc
// @Summary      Get one channel
// @Description  Returns a single channel of the organization's cables account. A channel belonging to another organization answers 404 rather than 403 — a 403 would confirm that another org's id exists.
// @Tags         Orchestrator.Cables
// @Accept       json
// @Produce      json
// @Param        Cookie     header    string  true  "Session cookie containing org_id and empl_id"
// @Param        channelID  path      string  true  "Channel ID"
// @Success      200        {object}  utils.Response{data=cablesEngineDtos.ChannelResponseDto} "Channel fetched successfully"
// @Failure      400        {object}  utils.Response "Invalid channel id"
// @Failure      401        {object}  utils.Response "Unauthorized"
// @Failure      404        {object}  utils.Response "Channel not found"
// @Failure      409        {object}  utils.Response "Cables service is not started for this organization"
// @Failure      500        {object}  utils.Response "Internal server error"
// @Router       /orchestrator/cables/channel/{channelID} [get]
func (c *OrchestratorController) GetCablesChannelHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	channelID, ok := c.getCablesChannelIDFromParams(ctx)
	if !ok {
		return
	}

	channel, err := c.orchestService.GetCablesChannelDetails(ctx.Request.Context(), orgID, emplID, channelID)
	if err != nil {
		failFromCablesError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Channel fetched successfully", channel)
}
