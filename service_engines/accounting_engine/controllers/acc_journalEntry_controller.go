package accounting_controllers

import (
	"net/http"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
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

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */

func (c *AccountingController) CreateJournalEntry(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	legalEntityID, ok := c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[accountingdtos.JournalEntryExternalRequestDto](ctx)
	if !ok {
		return
	}

	journalEntryID, err := c.accSvc.CreateJournalEntry(ctx.Request.Context(), orgID, emplID, legalEntityID, *req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to Create Journal Entry", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Journal Entry Created Successfully", journalEntryID)

}
