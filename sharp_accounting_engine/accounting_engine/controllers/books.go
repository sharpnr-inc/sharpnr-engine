package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *AccountingController) GetAccountingBooks(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	legalEntityID, ok := c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	response, err := c.accSvc.GetAccountingBook(ctx.Request.Context(), orgID, emplID, legalEntityID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Books Fetched Successfully", response)

}
