package restaurantcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) GetRestaurantDetails(ctx *gin.Context) {

	// ✅ Get OrgUUID from context (set by middleware)
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	res, err := c.resSvc.GetRestaurantDetails(ctx, orgID)
	if err != nil {
		response := dto.Response[any]{
			StatusCode: http.StatusOK,
			Message:    "There is Error Feting",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusOK, response)
	}

	response := dto.Response[*hospitalitydtos.Restaurant_GetRestaurant_ResponseDTO]{
		StatusCode: http.StatusOK,
		Message:    "restaurant Fetched Successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &res,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, response)

}
