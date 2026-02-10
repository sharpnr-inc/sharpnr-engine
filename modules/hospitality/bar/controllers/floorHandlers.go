package barcontrollers

import (
	"github.com/gin-gonic/gin"
)

func (c *barController) AddFloorInBar(ctx *gin.Context) {

	// 1️⃣ Get OrgUUID from middleware
	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusUnauthorized,
	// 		Message:    "Organization UUID missing",
	// 		Error: dto.ErrorResponse{
	// 			Message: "OrgUUID not found in context",
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}

	// 	ctx.JSON(http.StatusUnauthorized, response)
	// 	return
	// }

	// orgID, ok := orgUUID.(uuid.UUID)
	// if !ok {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid organization UUID format",
	// 		Error: dto.ErrorResponse{
	// 			Message: "OrgUUID format mismatch",
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}

	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // 2️⃣ Bind request body
	// var req hospitalitydtos.BAR_Add_Floor_Req
	// if err := ctx.ShouldBindJSON(&req); err != nil {

	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid request payload",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}

	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // 3️⃣ Call service
	// if err := c.barService.CreateFloor(ctx, orgID, req); err != nil {

	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Failed to add floor",
	// 		Error: dto.ErrorResponse{
	// 			Message: err.Error(),
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}

	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // 4️⃣ SUCCESS RESPONSE
	// response := dto.Response[any]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Floor added successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, response)
}
