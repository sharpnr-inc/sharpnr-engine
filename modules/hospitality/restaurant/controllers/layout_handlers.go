package restaurantcontrollers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *restaurantController) GetRestaurantLayout(ctx *gin.Context) {
	// 1. Resolve Org and User context
	orgID, _, ok := c.resolveContext(ctx)
	if !ok {
		return // Response already handled by helper
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Response already handled by helper
	}

	// 3. Call Service
	layout, err := c.resSvc.GetRestaurantLayout(ctx.Request.Context(), orgID, restaurantID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadGateway, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Layout Fetched Successfully", layout)
}
