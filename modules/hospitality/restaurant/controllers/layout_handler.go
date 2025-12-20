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

	var req hospitalitydtos.Restaurant_CreateOrAddTables_Request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := c.resSvc.CreateOrAddTable(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "awesome"})

}

func (c *restaurantController) EditTablesDetailsOfRestaurant(ctx *gin.Context) {}
