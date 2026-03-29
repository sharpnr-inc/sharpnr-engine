package accounting_controllers

import (
	"net/http"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	"sharpnr-engine/shared/utils"

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
