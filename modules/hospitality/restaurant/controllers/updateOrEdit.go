package restaurantcontrollers

import (
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *restaurantController) UpdateRestaurant(ctx *gin.Context) {
	/*
		   TODO Add Req DTO we Already have Response DTO
		//    *1 Add REQ DTO
		   *2 Add Repo Method and Service
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

	// 2️⃣ Bind JSON request body
	var req hospitalitydtos.Restaurant_UpdateOrEdit_General_Details_REQ

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// Optional: Ensure the HotelId in the payload matches OrgID (if required)
	// if req.HotelId != orgID.String() {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Hotel ID mismatch"})
	// 	return
	// }

	// 3️⃣ Call service method
	okRes, err := c.resSvc.UpdateRestaurantGeneralDetails(ctx, req, orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update restaurant details",
			"details": err.Error(),
		})
		return
	}

	// 4️⃣ Return success response
	ctx.JSON(http.StatusOK, gin.H{
		"success": okRes,
		"message": "Restaurant details updated successfully",
	})

}
