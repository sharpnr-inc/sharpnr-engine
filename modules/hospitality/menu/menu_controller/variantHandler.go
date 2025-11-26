package menucontroller

import (
	"fmt"
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *menuController) AddVariantInMenuItem(ctx *gin.Context) {
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

	fmt.Println("ORG:", orgID)

	var req hospitalitydtos.Menu_CreateVariantGroupAndVariant_REQ

	// ❗ MUST bind JSON, otherwise req is EMPTY
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pass DTO to service
	if err := c.menuService.CreateVariantGroup(ctx, req.MenuItemID, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Variant group added successfully"})
}

//-----------------------------------------------------------------------------------------------------------------------------------------//

func (c *menuController) GetVariantOfMenuItem(ctx *gin.Context) {
	// 1. Get OrgUUID from context
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "organization_uuid_missing"})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid_org_uuid"})
		return
	}

	// 2. Extract menu item ID from URL params
	idParam := ctx.Param("id")
	menuItemID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid_menu_item_id"})
		return
	}

	// 3. Fetch VariantGroup + Variants from service
	resp, err := c.menuService.GetVariantGroupByMenuId(ctx, orgID, menuItemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. Return success JSON
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resp,
	})
}
