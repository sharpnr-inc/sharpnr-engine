package barcontrollers

import (
	"github.com/gin-gonic/gin"
)

// GetBar godoc
// @Summary Get bar details for the current organization
// @Description Retrieves bar information linked to the authenticated organization UUID.
// @Tags Hospitality - Bar
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "Bar details retrieved successfully"
// @Failure 400 {object} map[string]string "Invalid organization UUID format"
// @Failure 401 {object} map[string]string "Organization UUID missing"
// @Failure 404 {object} map[string]string "No bar found for this organization"
// @Failure 500 {object} map[string]string "Failed to fetch bar details"
// @Router /hospitality/bar [get]
func (c *barController) GetBarLayout(ctx *gin.Context) {

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

	// // 2️⃣ Call service to get bar details
	// res, err := c.barService.GetBarLayout(ctx, orgID)
	// if err != nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Failed to fetch bar details",
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

	// // 3️⃣ Bar not found
	// if res == nil {
	// 	response := dto.Response[any]{
	// 		StatusCode: http.StatusNotFound,
	// 		Message:    "No bar found for this organization",
	// 		Error: dto.ErrorResponse{
	// 			Message: "",
	// 			IsError: false,
	// 		},
	// 		Data:      nil,
	// 		Timestamp: time.Now(),
	// 	}
	// 	ctx.JSON(http.StatusNotFound, response)
	// 	return
	// }

	// // 4️⃣ SUCCESS RESPONSE
	// response := dto.Response[*hospitalitydtos.GetBarDetails_RESPONSE]{
	// 	StatusCode: http.StatusOK,
	// 	Message:    "Bar details retrieved successfully",
	// 	Error: dto.ErrorResponse{
	// 		Message: "",
	// 		IsError: false,
	// 	},
	// 	Data:      &res, // actual bar layout data
	// 	Timestamp: time.Now(),
	// }

	// ctx.JSON(http.StatusOK, response)
}
