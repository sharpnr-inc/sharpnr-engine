package barcontrollers

import (
	"fmt"
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Update godoc
// @Summary Update Bar Details
// @Description
// @Tags Hospitality - Bar
func (c *barController) UpdateBarGenDetails(ctx *gin.Context) {

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
	var req hospitalitydtos.BAR_BarUpdateReq

	fmt.Println(&req)
	fmt.Println(orgID)

}
