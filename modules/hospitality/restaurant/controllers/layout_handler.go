package restaurantcontrollers

import (
	"fmt"
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) CreateOrAddFloor(ctx *gin.Context) {}

func (c *restaurantController) CreateOrAddTables(ctx *gin.Context) {
	fmt.Println("---- CONTROLLER HIT ----")
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
	fmt.Println(orgID)

	fmt.Println("---- CONTROLLER HIT ----")

	var req hospitalitydtos.Restaurant_CreateOrAddTables_Request

	fmt.Println(&req)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	fmt.Println(req)

	if err := c.resSvc.CreateOrAddTable(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "awesome"})

}
