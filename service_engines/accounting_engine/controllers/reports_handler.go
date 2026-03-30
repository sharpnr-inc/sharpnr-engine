package accounting_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *AccountingController) GetProfitAndLossReport(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	legelEntityID, ok := c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	from := "2025-04-01"
	to := "2026-03-31"

	res, err := c.accSvc.GenerateProfitAndLossReport(ctx.Request.Context(), orgID, emplID, legelEntityID, from, to)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "something went wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Profit and Loss fetched successfully", res)

}
