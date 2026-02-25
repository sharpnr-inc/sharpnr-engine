package barcontrollers

import (
	"github.com/gin-gonic/gin"
)

// It will take array of tables in a floor
func (c *barController) AddTableInBar(ctx *gin.Context) {

	// 1️⃣ Get OrgUUID from middleware
	// orgUUID, exists := ctx.Get("OrgUUID")
	// if !exists {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusUnauthorized,
	// 		Message:    "Organization UUID missing",
	// 		Error: dto.ErrorResponse{
	// 			Message: "OrgUUID not provided in middleware",
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
	// 			Message: "UUID casting failed",
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // 2️⃣ Bind request
	// var req hospitalitydtos.Bar_CreateTablesInBar_Request
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid request body",
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

	// // 3️⃣ Validate BarId
	// if req.BarId == uuid.Nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "BarId is required",
	// 		Error: dto.ErrorResponse{
	// 			Message: "bar_id cannot be empty",
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}

	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // 4️⃣ Validate Table List
	// if len(req.TableList) == 0 {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Table list cannot be empty",
	// 		Error: dto.ErrorResponse{
	// 			Message: "Provide at least one table",
	// 			IsError: true,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}

	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // 5️⃣ CALL SERVICE
	// err := c.barService.CreateTablesInBar(ctx, orgID, req)
	// if err != nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Failed to create tables",
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

	// // 6️⃣ SUCCESS RESPONSE
	// response := dto.Response[any]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Tables created successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      nil,
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, response)
}
