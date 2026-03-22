package orchestrator_controllers

import (
	"net/http"
	payment_engine_contracts "sharpnr-engine/payment_engine/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) CreatePaymentAccount(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[payment_engine_contracts.CreatePaymentAccountRequest](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddPaymentAccount(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "PaymentAccount Created")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetPaymentAccount(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetAllPaymentAccount(ctx.Request.Context(), orgID, emplID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "PaymentAccount Fetched Successfully", res)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) AddGatwayConfigInPayment(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[payment_engine_contracts.AddGatewayConfigRequest](ctx)
	if !ok {
		return
	}

	if err := c.orchestService.AddGatewayConfigInPaymentAccount(ctx.Request.Context(), orgID, emplID, *req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "Config Added Successfully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *OrchestratorController) GetPaymentAccountGatewayDetails(ctx *gin.Context) {

	orgID, emplID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	paymentAccountID, err := c.getPaymentAccountIDFromParams(ctx)
	if err != nil {
		return
	}

	res, err := c.orchestService.GetPaymentAccountDetails(ctx.Request.Context(), orgID, emplID, paymentAccountID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusCreated, "PaymentAccount Details Fetched Successfully", res)
}

/* ---------------------------------- <***> --------------------------------- */
