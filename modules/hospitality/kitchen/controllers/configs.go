package kitchencontrollers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *kitchenController) GetKitchenDetails(ctx *gin.Context) {
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
