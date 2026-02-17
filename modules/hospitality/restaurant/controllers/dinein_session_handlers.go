package restaurantcontrollers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *restaurantController) CreateDineSession(ctx *gin.Context) {

	orgID, empID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Response already handled by helper
	}

	tableID, ok := c.resolveTableID(ctx)
	if !ok {
		return
	}

	_, err := c.resSvc.StartDineSession(ctx.Request.Context(), restaurantID, orgID, empID, tableID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "DineIn Session Started Created Successfully")

}

/* ----------------------------------- <*> ---------------------------------- */

func (c *restaurantController) GetDineInSessionOrders(ctx *gin.Context) {
	_, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	_, ok = c.resolveRestaurantID(ctx)
	if !ok {
		return // Response already handled by helper
	}

	dineInSessionID, ok := c.resolveDineInSessionID(ctx)
	if !ok {
		return
	}

	res, err := c.resSvc.GetDineSessionOrders(ctx, dineInSessionID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Orders fetched successfully", res)
}
