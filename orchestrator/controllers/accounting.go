package orchestrator_controllers

import (
	"errors"
	"net/http"

	"sharpnr-engine/shared/utils"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"
	accountingEngineErrs "sharpnr-engine/sharp_accounting_engine/accounting_engine/errs"

	"github.com/gin-gonic/gin"
)

// failFromAccountingError maps the accounting account layer's sentinel errors
// onto status codes before falling back to the shared mapper.
//
// "Not initialized" and "disabled" are both 409 rather than 404: the org is real
// and the caller is allowed to ask, but the service is in a state that refuses
// the request — and the fix differs (start it vs. it was switched off), which is
// why the two errors stay apart all the way to the client.
func failFromAccountingError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, accountingEngineErrs.ErrAccountingServiceNotInitialized),
		errors.Is(err, accountingEngineErrs.ErrAccountingAccountIsDisabled),
		errors.Is(err, accountingEngineErrs.ErrAccountingServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)
	case errors.Is(err, accountingEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)
	default:
		failFromServiceError(ctx, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetAccountingAccountStatusHandler godoc
// @Summary      Accounting service status
// @Description  Reports whether the organization has taken the accounting service and whether that account is currently switched on. This is the org-level switch, not a legal entity's accounting book. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — this is the read the client makes before it knows.
// @Tags         Orchestrator.Accounting
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=accountingdtos.AccountingAccountStatusResponse} "Accounting status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /orchestrator/sharp-accounting [get]
func (c *OrchestratorController) GetAccountingAccountStatusHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	status, err := c.orchestService.GetAccountingAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromAccountingError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Accounting status fetched successfully", status)
}

/* ---------------------------------- <***> --------------------------------- */

// StartAccountingServiceHandler godoc
// @Summary      Start the accounting service
// @Description  Provisions the organization's one accounting-engine account together with its settings row. This is the org-level switch only — it creates no accounting book, which belongs to a legal entity and is created by the legal-entity setup flow. Starting an account that exists but was switched off turns it back on; starting a live one answers 409 rather than minting a second account. Both flags in the payload must be true — declining the terms is a 400, not a silent no-op.
// @Tags         Orchestrator.Accounting
// @Accept       json
// @Produce      json
// @Param        Cookie   header    string                                              true  "Session cookie containing org_id and empl_id"
// @Param        payload  body      accountingdtos.StartAccountingServiceRequestDto     true  "Terms acceptance and start flag"
// @Success      201      {object}  utils.Response "Accounting service started successfully"
// @Failure      400      {object}  utils.Response "Validation error"
// @Failure      401      {object}  utils.Response "Unauthorized"
// @Failure      409      {object}  utils.Response "Accounting service is already started"
// @Failure      500      {object}  utils.Response "Internal server error"
// @Router       /orchestrator/sharp-accounting/start [post]
func (c *OrchestratorController) StartAccountingServiceHandler(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[accountingdtos.StartAccountingServiceRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.StartAccountingService(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		failFromAccountingError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Accounting service started successfully")
}
