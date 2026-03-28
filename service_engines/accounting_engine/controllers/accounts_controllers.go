package accounting_controllers

import (
	"net/http"
	accountingdtos "sharpnr-engine/service_engines/accounting_engine/dto"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *AccountingController) CreateAccount(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContext(ctx)
	if !ok {
		return
	}

	legalEntityID, ok := c.resolveLegalEntityID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[accountingdtos.CreateAccountRequest](ctx)
	if !ok {
		return
	}

	_, err := c.accSvc.CreateAccount(
		ctx.Request.Context(),
		orgID,
		emplID,
		legalEntityID,
		*req,
	)

	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Success")
}
