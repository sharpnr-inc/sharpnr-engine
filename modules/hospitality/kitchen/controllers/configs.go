package kitchencontrollers

import (
	"net/http"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	"sharpnr-engine/shared/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *KitchenController) GetKitchenDetails(ctx *gin.Context) {
	// 1. Resolve Auth and KitchenID
	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	res, err := c.kitchenService.GetKitchenDetails(ctx, kitchenID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen details fetched successfully", res)

}

func (c *KitchenController) GetKitchenConfigs(ctx *gin.Context) {
	// 1. Resolve Auth and KitchenID
	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	res, err := c.kitchenService.GetKitchenConfigs(ctx, kitchenID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to fetch kitchen configs", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen configs fetched successfully", res)

}

func (c *KitchenController) UpdateKitchenGenralDetails(ctx *gin.Context) {
	// 1. Resolve Auth and KitchenID
	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	var req_dto kitchen_contracts.UpdateKitchenGeneralDetails_req_dto
	if err := ctx.ShouldBindJSON(&req_dto); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := c.kitchenService.UpdateGeneralDetails(ctx, kitchenID, req_dto); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Failed to update kitchen general details", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen configs updated successfully")

}

func (c *KitchenController) ActivateDeactivateKitchen(ctx *gin.Context) {
	orgID, emplID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	switchParam := ctx.Param("kitchen_switch")

	active, err := strconv.ParseBool(switchParam)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "switch must be true or false", err)
		return
	}

	if err := c.kitchenService.ActivateDeactivate(ctx.Request.Context(), orgID, emplID, kitchenID, active); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Kitchen operation done")

}
