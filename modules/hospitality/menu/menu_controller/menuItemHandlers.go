package menucontroller

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *menuController) CreateMenuItem(ctx *gin.Context) {

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

	var req hospitalitydtos.Menu_CreateMenuItem_REQ
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// default served_by if not sent
	if req.ServedBy == "" {
		req.ServedBy = "KITCHEN"
	}

	if err := c.menuService.CreateMenuItem(ctx, &req, orgID); err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create menu item",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Timestamp: time.Now(),
		}
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := dto.Response[any]{
		StatusCode: http.StatusOK,
		Message:    "Menu item created successfully",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      nil,
		Timestamp: time.Now(),
	}
	ctx.JSON(http.StatusOK, res)
}
