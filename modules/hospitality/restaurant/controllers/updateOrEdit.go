package restaurantcontrollers

import (
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) UpdateRestaurant(ctx *gin.Context) {
	// Step 1: Get OrgUUID from context (set by middleware)
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Organization UUID missing",
		})
		return
	}

	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid organization UUID format",
		})
		return
	}

	// Step 2: Bind and validate JSON request body
	var req hospitalitydtos.Restaurant_UpdateOrEdit_General_Details_REQ

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Step 3: Validate that at least one field is being updated
	if req.RestaurantLegalName == nil &&
		req.RestaurantDisplayName == nil &&
		req.FSSAINumber == nil &&
		req.GSTNumber == nil &&
		req.CINNumber == nil &&
		req.AddressLine == nil &&
		req.City == nil &&
		req.State == nil &&
		req.Country == nil &&
		req.Pincode == nil &&
		req.EmailAddress == nil &&
		req.ContactNumber == nil &&
		req.ManagerName == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No fields to update",
		})
		return
	}

	// Step 4: Call service method
	success, err := c.resSvc.UpdateRestaurantGeneralDetails(ctx, req, orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update restaurant details",
			"details": err.Error(),
		})
		return
	}

	// Step 5: Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": "Restaurant details updated successfully",
	})
}
