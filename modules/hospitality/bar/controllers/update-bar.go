package barcontrollers

import (
	"fmt"
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Update godoc
// @Summary Update Bar Details
// @Description
// @Tags Hospitality - Bar
func (c *barController) UpdateBarGenDetails(ctx *gin.Context) {

	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "organization UUID missing"})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization UUID format"})
		return
	}
	var req hospitalitydtos.BAR_BarUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(*req.FSSAINumber)
	fmt.Println(orgID)

}
