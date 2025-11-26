package restaurantcontrollers

import (
	"fmt"
	"net/http"

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

}

func (c *restaurantController) AddOrdersInSession(ctx *gin.Context) {}

func (c *restaurantController) CreateInvoiceForSession(ctx *gin.Context) {}

func (c *restaurantController) GetBookingSessionById(ctx *gin.Context) {}
