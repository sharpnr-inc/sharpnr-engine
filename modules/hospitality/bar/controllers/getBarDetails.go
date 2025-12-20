package barcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *barController) GetBarDetailsConfig(ctx *gin.Context) {

	// ✅ Get OrgUUID from context (set by middleware)
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	res, err := c.barService.GetBarDetailsConfig(ctx, orgID)

	if err != nil {
		response := dto.Response[*hospitalitydtos.BAR_UpdateBarDetails_Response]{
			StatusCode: http.StatusOK,
			Message:    "There is Error Fetching Bar Details for Config",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
		}

		ctx.JSON(http.StatusOK, response)
	}

	response := dto.Response[*hospitalitydtos.BAR_UpdateBarDetails_Response]{
		StatusCode: http.StatusOK,
		Message:    "Bar details Fetched Successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &res,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *barController) Update(ctx *gin.Context) {

	// Get OrgUUID from context
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}

	// Bind JSON
	var req hospitalitydtos.BAR_BarUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// BarID is required
	if req.ID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bar ID is required"})
		return
	}

	// Call service
	err := c.barService.UpdateBar(ctx, orgID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bar updated successfully",
	})
}
