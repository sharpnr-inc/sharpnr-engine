package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *HotelController) GetHotel(ctx *gin.Context) {
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	// Type assert to UUID type (if you stored it as uuid.UUID)
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	// ✅ Call service
	res, err := s.hotelService.GetHotelDetails(ctx, orgID)
	if err != nil {
		response := dto.Response[any]{
			StatusCode: http.StatusOK,
			Message:    "There was an Error Fetching Hotel Details",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := dto.Response[*hospitalitydtos.GetHotel_RESPONSE]{
		StatusCode: http.StatusOK,
		Message:    "Hotel Fetched Successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &res,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, response)

}
