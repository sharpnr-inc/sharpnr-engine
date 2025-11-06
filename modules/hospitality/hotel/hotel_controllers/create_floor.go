package hotelcontrollers

import (
	"fmt"
	"net/http"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *HotelController) CreateFloorInHotel(ctx *gin.Context) {

	var req hospitalitydtos.CreateFloorInHotel_REQ

	// ✅ Parse JSON body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

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

	fmt.Println("Incoming Floor Request:", req)

	// ✅ Call service layer
	res, err := s.hotelService.CreateFloor(ctx, orgID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)

}
