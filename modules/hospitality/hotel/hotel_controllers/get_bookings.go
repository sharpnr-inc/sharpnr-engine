package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *HotelController) GetHotelsBookings(ctx *gin.Context) {
	resData := dto.Response[any]{
		StatusCode: http.StatusOK,
		Message:    "This will fetch the Data for the Creating",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, resData)
}
