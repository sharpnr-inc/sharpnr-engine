package hotelcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *HotelController) GetHotelOverview(ctx *gin.Context) {
	/*
		TODO We have have fetch the Data How many Room are avaialble
	*/
	data := dto.Response[any]{
		StatusCode: 200,
		Message:    "Fetched Data",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, data)
}
