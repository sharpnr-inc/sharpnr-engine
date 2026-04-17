package restaurantcontrollers

import (
	"net/http"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (c *RestaurantController) CreateDineSession(ctx *gin.Context) {

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

/* ----------------- <**Add Order in the Dining Session ** > ---------------- */

func (c *RestaurantController) AddOrderInDiningSession(ctx *gin.Context) {

	orgID, _, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Response already handled by helper
	}

	dineInSessionID, ok := c.resolveDineInSessionID(ctx)
	if !ok {
		return
	}

	var req *restaurant_contracts.DiningSessionOrder_req_dto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "", err)
		return
	}

	if err := c.resSvc.AddOrdersInSession(ctx.Request.Context(),
		dineInSessionID,
		req,
		orgID,
		restaurantID,
	); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "", "Add successfully")
}

/* ----------------------------------- <*> ---------------------------------- */

func (c *RestaurantController) GetDineInSessionOrders(ctx *gin.Context) {
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

/* ---------------------------------- <***> --------------------------------- */

func (c *RestaurantController) CloseDineInSession(ctx *gin.Context) {

	orgID, empID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	// 2. Resolve Restaurant ID from params
	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Response already handled by helper
	}

	dineInSessionID, ok := c.resolveDineInSessionID(ctx)
	if !ok {
		return
	}

	req, ok := utils.BindJSON[restaurant_contracts.CloseSessionRequest](ctx)
	if !ok {
		return
	}

	// this is the 1 step we get the request for the Closing the session

	response, err := c.resSvc.CloseDineInSession(
		ctx.Request.Context(),
		orgID, empID, restaurantID, dineInSessionID,
		*req,
	)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, "Something Went Wrong", err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Session closed", response)

}

/* ---------------------------------- <***> --------------------------------- */
