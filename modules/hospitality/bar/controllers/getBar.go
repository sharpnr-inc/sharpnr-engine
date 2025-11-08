package barcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *barController) GetBar(ctx *gin.Context) {

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

	// ✅ 2️⃣ Call service to get bar details
	res, err := c.barService.GetBarDetails(ctx, orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to fetch bar details",
			"message": err.Error(),
		})
		return
	}

	// ✅ 3️⃣ Handle bar not found
	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no bar found for this organization"})
		return
	}

	// ✅ 4️⃣ Send success response
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "bar details retrieved successfully",
		"data":    res,
	})

}
