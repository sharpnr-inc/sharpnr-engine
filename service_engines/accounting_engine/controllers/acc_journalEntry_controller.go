package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *AccountingController) GetJournalEntries(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	legalEntityID, ok := c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	res, err := c.accSvc.GetLegalEntityLedgerEntry(ctx.Request.Context(), orgID, emplID, legalEntityID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Data Fetched Successfully", res)

}

func (c *AccountingController) GetJornalDetails(ctx *gin.Context) {
	_, _, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	_, ok = c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	entryID, ok := c.resolveJournalEntryID(ctx)
	if !ok {
		return
	}

	res, err := c.accSvc.GetJournalEntryDetails(ctx.Request.Context(), entryID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Some thing Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Details Fetched Successfully", res)

}
