package barcontrollers

import (
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *barController) AddFloorInBar(ctx *gin.Context) {
	// Get OrgID
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

	// Bind Request body
	var req hospitalitydtos.BAR_Add_Floor_Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call Service
	if err := c.barService.CreateFloor(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Floor added successfully"})
}
