package restaurantcontrollers

import (
	"net/http"
	restaurant_contracts "sharpnr-engine/business_ops/hospitality/restaurant/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Inventory Account/Warehouse List Fetched Successfully", acc)

}

/* ---------------------------------- <***> --------------------------------- */

func (c *RestaurantController) AddInventoryInInventoryAccount(ctx *gin.Context) {
	orgID, empID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[restaurant_contracts.AddInventoryRequestDto](ctx)
	if !ok {
		return
	}

	if err := c.resSvc.AddNewInventoryInRestaurant(ctx.Request.Context(), orgID, empID, restaurantID, req); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Inventory Added SuccessFully")

}

/* ---------------------------------- <***> --------------------------------- */

func (c *RestaurantController) GetBusinessUnitInventories(ctx *gin.Context) {
	orgID, empID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return
	}

	// Optional filter
	var inventoryAccountID *uuid.UUID

	if val := ctx.Query("inventoryAccountID"); val != "" {
		parsedID, err := uuid.Parse(val)
		if err != nil {
			utils.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		inventoryAccountID = &parsedID
	}

	res, err := c.resSvc.GetRestaurantAllInventories(
		ctx.Request.Context(),
		orgID, empID, restaurantID,
		inventoryAccountID, // optional filter passed
	)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Inventory Added SuccessFully", res)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *RestaurantController) GetInventoryItems(ctx *gin.Context) {
}

/* ---------------------------------- <***> --------------------------------- */
