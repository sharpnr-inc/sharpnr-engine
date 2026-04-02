package restaurantcontrollers

import (
	"net/http"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *RestaurantController) GetAllInventoryAccounts(ctx *gin.Context) {
	orgID, empID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	_, ok = c.resolveRestaurantID(ctx)
	if !ok {
		return
	}

	// In Inventory Context the RestaurantID is BusinessID
	// If Any ApI need the BusinessID in Inventory Handlers use the restaurantID

	acc, err := c.resSvc.GetAllInventoryAccounts(ctx.Request.Context(), orgID, empID)
	if err != nil {

		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Inventory Account/Warehouse List Fetched Successfully", acc)

}
