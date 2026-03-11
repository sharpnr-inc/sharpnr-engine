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
