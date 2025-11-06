package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *HotelController) GetRequestInHotels(ctx *gin.Context) {
	resData := dto.Response[any]{
		StatusCode: 200,
		Message:    "data fetched successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, resData)
}
