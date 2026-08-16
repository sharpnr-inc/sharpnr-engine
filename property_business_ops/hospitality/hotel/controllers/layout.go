package hotelcontrollers

import (
	"errors"
	"net/http"
	webMiddleware "sharpnr-engine/middlewares/web"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *hotelController) GetHotelLayout(ctx *gin.Context) {

	if _, exists := ctx.Get(webMiddleware.UserIDContextKey); !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("user not authenticated"))
		return
	}

	// 2. Org context check
	orgUUIDRaw, exists := ctx.Get(webMiddleware.OrgUUIDContextKey)
	if !exists {
		utils.Failure(ctx, http.StatusUnauthorized, errors.New("organization not found in context"))
		return
	}

	_, ok := orgUUIDRaw.(uuid.UUID)
	if !ok {
		utils.Failure(ctx, http.StatusInternalServerError, errors.New("invalid org uuid type"))
		return
	}

	hotelID := ctx.Param("hotelID")
	parsedHotelID, err := uuid.Parse(hotelID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadRequest, errors.New("invalid restaurantID format"))
		return
	}

	res, err := c.hotelService.GetHotelLayout(ctx.Request.Context(), parsedHotelID)
	if err != nil {
		utils.Failure(ctx, http.StatusBadGateway, err)
		return
	}

	utils.Success(ctx, http.StatusOK, "HotelLayout Fetched", res)

}
