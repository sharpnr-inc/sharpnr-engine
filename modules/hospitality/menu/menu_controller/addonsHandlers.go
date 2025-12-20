package menucontroller

import (
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *menuController) CreateAddons(ctx *gin.Context) {
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

	var req hospitalitydtos.Menu_CreateAddonsInMenu_REQUEST

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// 3. Call service
	if err := s.menuService.CreateAddForMenuItem(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. Success response
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "addon created successfully",
	})
}
