package accounting_controllers

import (
	"errors"
	"net/http"

	"sharpnr-engine/shared/utils"
	accountingEngineErrs "sharpnr-engine/sharp_accounting_engine/accounting_engine/errs"

	"github.com/gin-gonic/gin"
)

// failFromAccountingAccountError maps the account layer's sentinel errors onto
// status codes.
//
// It is the client-facing twin of the orchestrator's mapper of the same name;
// the two are separate because the packages are, and this one only has to cover
// the errors the product API can actually raise.
//
// "Not initialized" and "disabled" are both 409 rather than 404: the org is real
// and the caller is allowed to ask, but the service is in a state that refuses
// the request — and the fix differs, which is why the two stay apart all the way
// to the client.
func failFromAccountingAccountError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, accountingEngineErrs.ErrAccountingServiceNotInitialized),
		errors.Is(err, accountingEngineErrs.ErrAccountingAccountIsDisabled),
		errors.Is(err, accountingEngineErrs.ErrAccountingServiceAlreadyStarted):
		utils.Failure(ctx, http.StatusConflict, err)

	case errors.Is(err, accountingEngineErrs.ErrTermsNotAccepted):
		utils.Failure(ctx, http.StatusBadRequest, err)

	default:
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// GetAccountingAccountStatusHandler godoc
// @Summary      Accounting service status
// @Description  Reports whether the organization has taken the accounting service and whether that account is switched on. This is the org-level switch, not a legal entity's accounting book. An organization that never started it answers serviceStarted=false with a null id rather than a 404 — the client asks this before it knows. Starting the service is an orchestrator call, not this one.
// @Tags         Accounting
// @Accept       json
// @Produce      json
// @Param        Cookie  header    string  true  "Session cookie containing org_id and empl_id"
// @Success      200     {object}  utils.Response{data=accountingdtos.AccountingAccountStatusResponse} "Accounting status fetched successfully"
// @Failure      401     {object}  utils.Response "Unauthorized"
// @Failure      500     {object}  utils.Response "Internal server error"
// @Router       /accounting/status [get]
func (c *AccountingController) GetAccountingAccountStatusHandler(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	status, err := c.accSvc.GetAccountingAccountStatus(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		failFromAccountingAccountError(ctx, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Accounting status fetched successfully", status)
}
