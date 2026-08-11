package restaurantcontrollers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// resolveContext extracts the OrgID and EmployeeID (UserID) from the middleware context.
func (c *RestaurantController) resolveContext(ctx *gin.Context) (orgID uuid.UUID, employeeID uuid.UUID, ok bool) {
	// 1. Resolve and Parse Employee/User ID
	userIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, uuid.Nil, false
	}

	// Convert interface to string and parse
	var err error
	employeeID, err = uuid.Parse(fmt.Sprintf("%v", userIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid user uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	// 2. Resolve and Parse Org ID
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, uuid.Nil, false
	}

	// Convert interface to string and parse
	orgID, err = uuid.Parse(fmt.Sprintf("%v", orgUUIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid org uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, employeeID, true
}

// resolveRestaurantID extracts and parses the "restaurantID" from the URL parameters.
func (c *RestaurantController) resolveRestaurantID(ctx *gin.Context) (uuid.UUID, bool) {
	restaurantIDStr := ctx.Param("restaurantID")
	parsedID, err := uuid.Parse(restaurantIDStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid restaurantID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

// resolveRestaurantID extracts and parses the "tableID" from the URL parameters.
func (c *RestaurantController) resolveTableID(ctx *gin.Context) (uuid.UUID, bool) {
	tableIDStr := ctx.Param("tableID")
	parsedID, err := uuid.Parse(tableIDStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid tableID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

func (c *RestaurantController) resolveDineInSessionID(ctx *gin.Context) (uuid.UUID, bool) {
	dineSessionIDStr := ctx.Param("dineSessionID")
	parsedID, err := uuid.Parse(dineSessionIDStr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid Dine In Session ID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

func (c *RestaurantController) resolveInventoryAccountID(ctx *gin.Context) (uuid.UUID, bool) {
	inventoryAccountID := ctx.Param("inventoryAccountID")
	parsedID, err := uuid.Parse(inventoryAccountID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid inventory Account ID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

// resolveBody binds the JSON request body to the provided struct type.
func (c *RestaurantController) resolveBody(ctx *gin.Context, target interface{}) bool {
	if err := ctx.ShouldBindJSON(target); err != nil {
		utils.Failure(ctx, http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))
		return false
	}
	return true
}
