package restaurantcontrollers

import (
	"errors"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) GetRestaurantLayout(ctx *gin.Context) {

	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("user not authenticated"))
		return
	}

	// 2. Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, utils.CodeUnauthorized, errors.New("organization not found in context"))
		return
	}

	orgID, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, utils.CodeInternalError, errors.New("invalid org uuid type"))
		return
	}

	restaurantID := ctx.Param("restaurantID")
	parsedRestaurantID, err := uuid.Parse(restaurantID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, utils.CodeValidationFailed, errors.New("invalid restaurantID format"))
		return
	}

	layout, err := c.resSvc.GetRestaurantLayout(ctx.Request.Context(), orgID, parsedRestaurantID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadGateway, utils.CodeInternalError, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "Layout Fetched Successfully", layout)

}
