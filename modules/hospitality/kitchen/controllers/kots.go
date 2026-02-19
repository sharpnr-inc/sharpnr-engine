package kitchencontrollers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *kitchenController) GetKots(ctx *gin.Context) {

	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	kitchenID, ok := c.resolveKitchenID(ctx)
	if !ok {
		return
	}

	res, err := c.kitchenService.GetKots(ctx.Request.Context(), kitchenID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}


	utils.Success(ctx, http.StatusOK, "KOT fetchet", res)

}
