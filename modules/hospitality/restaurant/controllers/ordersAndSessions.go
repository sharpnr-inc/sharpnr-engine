package restaurantcontrollers

import (
	"fmt"
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) CreateSession(ctx *gin.Context) {

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

	var req hospitalitydtos.Hospitality_Restaurant_CreateSession_REQUEST

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := c.resSvc.CreateTableBookingSession(ctx, orgID, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "Booking session Created"})

}

//---------------------------------------------------------------------------------------------------------------//

func (c *restaurantController) AddOrdersInSession(ctx *gin.Context) {

	/*
		So we will need the SessionId and request body for the Order
		MenuItemId and its VariantId and Addons
		We should take the sessionId in the params
		body might have the
		OrgId based on that we can Fetch the Kitchen
		After that Create the KOT
		Give the MenuItemId and give SessionId as ref
	*/

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
}

func (c *restaurantController) CreateInvoiceForSession(ctx *gin.Context) {}

func (c *restaurantController) GetBookingSessionById(ctx *gin.Context) {

	// Here we have the OrgId and we can take the OrgId and Fetch the Restaurant
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

	// --------------------------------
	// 2️⃣ Read tableId from URL param
	// --------------------------------
	tableIdParam := ctx.Param("id")
	if tableIdParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Table ID not provided",
		})
		return
	}

	tableID, err := uuid.Parse(tableIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid table ID format",
		})
		return
	}

	session, err := c.resSvc.GetResTableBookingSessionByTableId(ctx, orgID, tableID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "something bad happen"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"res": session})

}
