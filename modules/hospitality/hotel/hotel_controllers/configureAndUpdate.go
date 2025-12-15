package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *HotelController) ConfigureAndUpdate(ctx *gin.Context) {
	// 1️⃣ Get org UUID from context
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
			Error: dto.ErrorResponse{
				Message: "Organization UUID missing",
				IsError: true,
			},
			Timestamp: time.Now(),
		})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid organization ID",
			Error: dto.ErrorResponse{
				Message: "Invalid organization UUID format",
				IsError: true,
			},
			Timestamp: time.Now(),
		})
		return
	}

	// 2️⃣ Bind request body
	var req hospitalitydtos.Hotel_UpdateOrConfigure_REQUEST
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Timestamp: time.Now(),
		})
		return
	}

	// 3️⃣ Call service
	if err := c.hotelService.UpdateHotelDetails(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update hotel",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Timestamp: time.Now(),
		})
		return
	}

	// 4️⃣ Success response
	ctx.JSON(http.StatusOK, dto.Response[any]{
		StatusCode: http.StatusOK,
		Message:    "Hotel updated successfully",
		Error: dto.ErrorResponse{
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	})
}
