package orchestratorcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *orchestratorOrganizationController) GetDetails(ctx *gin.Context) {
	orgUUID, exists := ctx.Get("OrgUUID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Organization UUID missing"})
		return
	}
	// Type assert to UUID type (if you stored it as uuid.UUID)
	orgID, ok := orgUUID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization UUID format"})
		return
	}
	fmt.Println(orgID)

	ctx.JSON(http.StatusOK, gin.H{"Error": "thfjsoij"})
}
