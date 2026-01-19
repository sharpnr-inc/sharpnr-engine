package hotelcontrollers

import (
	"github.com/gin-gonic/gin"
)

func (s *HotelController) GetHotel(ctx *gin.Context) {
	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
	// 	return
	// }
	// // Type assert to UUID type (if you stored it as uuid.UUID)
	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
	// 	return
	// }

	// // ✅ Call service
	// res, err := s.hotelService.GetHotelDetails(ctx, orgID)
	// if err != nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusOK,
	// 		Message:    "There was an Error Fetching Hotel Details",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, response)
	// 	return
	// }

	// response := dto.Response[*hospitalitydtos.GetHotel_RESPONSE]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Hotel Fetched Successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      &res,
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, response)

}

func (c *HotelController) ConfigureAndUpdate(ctx *gin.Context) {
	// 1️⃣ Get org UUID from context
	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
	// 		StatusCode: http.StatusUnauthorized,
	// 		Message:    "Unauthorized",
	// 		Error: dto.ErrorResponse{
	// 			Message: "Organization UUID missing",
	// 			IsError: true,
	// 		},
	// 		Timestamp: time.Now(),
	// 	})
	// 	return
	// }

	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	ctx.JSON(http.StatusBadRequest, dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid organization ID",
	// 		Error: dto.ErrorResponse{
	// 			Message: "Invalid organization UUID format",
	// 			IsError: true,
	// 		},
	// 		Timestamp: time.Now(),
	// 	})
	// 	return
	// }

	// // 2️⃣ Bind request body
	// var req hospitalitydtos.Hotel_UpdateOrConfigure_REQUEST
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid request body",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Timestamp: time.Now(),
	// 	})
	// 	return
	// }

	// // 3️⃣ Call service
	// if err := c.hotelService.UpdateHotelDetails(ctx, orgID, req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Failed to update hotel",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Timestamp: time.Now(),
	// 	})
	// 	return
	// }

	// // 4️⃣ Success response
	// ctx.JSON(http.StatusOK, dto.Response[any]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Hotel updated successfully",
	// 	Error: dto.ErrorResponse{
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// })
}

func (s *HotelController) GetHotelOverview(ctx *gin.Context) {
	/*
		TODO We have have fetch the Data How many Room are available
	*/
	// data := dto.Response[any]{
	// 	StatusCode: 200,
	// 	Message:    "Fetched Data",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// }
	// ctx.JSON(http.StatusOK, data)
}

func (s *HotelController) GetRequestInHotels(ctx *gin.Context) {
	// resData := dto.Response[any]{
	// 	StatusCode: 200,
	// 	Message:    "data fetched successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, resData)
}
