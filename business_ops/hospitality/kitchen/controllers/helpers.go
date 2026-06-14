package kitchencontrollers

import (
	"errors"
	"fmt"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *KitchenController) resolveContext(ctx *gin.Context) (orgID uuid.UUID, employeeID uuid.UUID, ok bool) {
	// 1. Resolve and Parse Employee/User ID
	userIDRaw, exists := ctx.Get(webMiddleware.UserIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return uuid.Nil, uuid.Nil, false
	}

	// Convert interface to string and parse
	var err error
	employeeID, err = uuid.Parse(fmt.Sprintf("%v", userIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid user uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	// 2. Resolve and Parse Org ID
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return uuid.Nil, uuid.Nil, false
	}

	// Convert interface to string and parse
	orgID, err = uuid.Parse(fmt.Sprintf("%v", orgUUIDRaw))
	if err != nil {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid format"))
		return uuid.Nil, uuid.Nil, false
	}

	return orgID, employeeID, true
}

func (c *KitchenController) resolveKitchenID(ctx *gin.Context) (uuid.UUID, bool) {
	kitchenIDstr := ctx.Param("kitchenID")
	parsedID, err := uuid.Parse(kitchenIDstr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid Kitchen ID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}

func (c *KitchenController) resolveKOTsID(ctx *gin.Context) (uuid.UUID, bool) {
	kotIDstr := ctx.Param("kotID")
	parsedID, err := uuid.Parse(kotIDstr)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid KOT ID format"))
		return uuid.Nil, false
	}
	return parsedID, true
}
