package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"
	accountingdtos "sharpnr-engine/sharp_accounting_engine/accounting_engine/dto"

	"github.com/gin-gonic/gin"
)

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

/* ---------------------------------- <***> --------------------------------- */

func (c *AccountingController) GetJournalEntries(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	legalEntityID, ok := c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	res, err := c.accSvc.GetJournalEntries(ctx.Request.Context(), orgID, emplID, legalEntityID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to Fetch Journal Entries", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Journal Entries Fetched Successfully", res)
}
