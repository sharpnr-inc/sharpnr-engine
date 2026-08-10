package orchestrator_controllers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *OrchestratorController) GetRestaurantDetails(ctx *gin.Context) {
	orgID, employeeID, ok := c.resolveRequestContexts(ctx)
	if !ok {
		return
	}

	restaurantID, ok := c.getRestaurntIDFromParams(ctx)
	if !ok {
		return
	}

	res, err := c.orchestService.GetRestaurantDetailsByRestaurantID(ctx.Request.Context(), orgID, employeeID, restaurantID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
	}

	utils.Success(ctx, http.StatusOK, "Restaurant Fetched Successly", res)
}
