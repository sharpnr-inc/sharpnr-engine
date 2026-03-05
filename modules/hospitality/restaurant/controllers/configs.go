package restaurantcontrollers

import (
	"errors"
	"net/http"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *RestaurantController) GetGeneralDetails(ctx *gin.Context) {

	orgID, employeeID, ok := c.resolveContext(ctx)
	if !ok {
		return // Context resolution failed, response already sent

	}

	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Restaurant ID resolution failed, response already sent
	}

	response, err := c.resSvc.GetGeneralDetails(ctx, orgID, employeeID, restaurantID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Restaurant general details fetched successfully", response)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *RestaurantController) UpdateGeneralDetails(ctx *gin.Context) {

	orgID, employeeID, ok := c.resolveContext(ctx)
	if !ok {
		return // Context resolution failed, response already sent

	}

	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Restaurant ID resolution failed, response already sent
	}

	var req restaurant_contracts.UpdateBodyRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "SomeThing went Wrong", err)
		return
	}

	err := c.resSvc.UpdateGeneralDetails(ctx, orgID, employeeID, restaurantID, req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Restaurant general details updated successfully", nil)
}

/* ---------------------------------- <***> --------------------------------- */

func (c *RestaurantController) AddLayout(ctx *gin.Context) {

	orgID, employeeID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Restaurant ID resolution failed, response already sent
	}

	var req restaurant_contracts.AddLayoutRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, "SomeThing went Wrong", err)
		return
	}

	err := c.resSvc.SetupRestaurantLayout(ctx, orgID, employeeID, restaurantID, req)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Restaurant layout added successfully", nil)

}

func (c *RestaurantController) GetRestaurantPreferredKitchen(ctx *gin.Context) {

	orgID, employeeID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Restaurant ID resolution failed, response already sent
	}

	response, err := c.resSvc.GetRestaurantPreferredKitchenID(ctx, orgID, employeeID, restaurantID)
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Restaurant preferred kitchen fetched successfully", response)

}

func (c *RestaurantController) AddRestaurantPreferredKitchen(ctx *gin.Context) {
	orgID, employeeID, ok := c.resolveContext(ctx)
	if !ok {
		return
	}

	restaurantID, ok := c.resolveRestaurantID(ctx)
	if !ok {
		return // Restaurant ID resolution failed, response already sent
	}

	pfkStr := ctx.Param("pfKitchenID")
	parsedPfkStr, err := uuid.Parse(pfkStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid preferred Kitchen ID format"))
		return
	}

	if err := c.resSvc.AddPreferredKitchen(ctx.Request.Context(), orgID, employeeID, restaurantID, parsedPfkStr); err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Preferred KitchenID", nil)
}
